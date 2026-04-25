package whitelist

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"nhooyr.io/websocket"

	"github.com/mtn-server/backend/config"
)

const oneBotCommandUsage = "Invalid command. Use: bind <java|bedrock> <playername> or unbind <java|bedrock> <playername>."

// OneBotService consumes NapCat OneBot 11 group messages and replies with results.
type OneBotService struct {
	cfg              config.OneBotConfig
	whitelistService *Service
	reconnectDelay   time.Duration
}

// NewOneBotService creates a OneBot command consumer.
func NewOneBotService(cfg config.OneBotConfig, whitelistService *Service) *OneBotService {
	delay := time.Duration(cfg.ReconnectSeconds) * time.Second
	if delay <= 0 {
		delay = 5 * time.Second
	}

	return &OneBotService{
		cfg:              cfg,
		whitelistService: whitelistService,
		reconnectDelay:   delay,
	}
}

// Start connects to NapCat and keeps reconnecting until ctx is cancelled.
func (o *OneBotService) Start(ctx context.Context) {
	if o == nil || !o.cfg.Enabled {
		return
	}
	if o.cfg.WebSocketURL == "" || o.cfg.QQGroupID == "" {
		log.Println("OneBot whitelist service disabled: ws_url or qq_group_id missing")
		return
	}

	for {
		if err := o.run(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Printf("OneBot whitelist connection ended: %v", err)
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(o.reconnectDelay):
		}
	}
}

func (o *OneBotService) run(ctx context.Context) error {
	headers := http.Header{}
	if o.cfg.Token != "" {
		headers.Set("Authorization", "Bearer "+o.cfg.Token)
	}

	conn, _, err := websocket.Dial(ctx, o.cfg.WebSocketURL, &websocket.DialOptions{
		HTTPHeader: headers,
	})
	if err != nil {
		return fmt.Errorf("dial onebot websocket: %w", err)
	}
	defer conn.Close(websocket.StatusNormalClosure, "shutdown")

	log.Println("OneBot whitelist service connected")
	sendMu := &sync.Mutex{}

	for {
		_, data, err := conn.Read(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return fmt.Errorf("read onebot websocket: %w", err)
		}

		event, ok := parseOneBotGroupEvent(data)
		if !ok {
			continue
		}
		if event.PostType != "message" || event.MessageType != "group" || string(event.GroupID) != o.cfg.QQGroupID {
			continue
		}

		command, matched, err := ParseCommand(event.RawMessage)
		if !matched {
			continue
		}
		if err != nil {
			o.sendGroupReply(ctx, conn, sendMu, event.GroupID, oneBotCommandUsage)
			continue
		}

		go o.handleCommand(ctx, conn, sendMu, event, command)
	}
}

func (o *OneBotService) handleCommand(parent context.Context, conn *websocket.Conn, sendMu *sync.Mutex, event oneBotGroupEvent, command *Command) {
	ctx, cancel := context.WithTimeout(parent, 20*time.Second)
	defer cancel()

	admin := event.Sender.Role == "owner" || event.Sender.Role == "admin"
	req := OperationRequest{
		Action:       command.Action,
		Edition:      command.Edition,
		PlayerName:   command.PlayerName,
		QQID:         string(event.UserID),
		ActorQQID:    string(event.UserID),
		Source:       SourceOneBot,
		Admin:        admin,
		EnforceQuota: !admin,
	}

	var (
		result *OperationResult
		err    error
	)
	switch command.Action {
	case ActionAdd:
		result, err = o.whitelistService.Add(ctx, req)
	case ActionRemove:
		result, err = o.whitelistService.Remove(ctx, req)
	default:
		err = ErrInvalidInput
	}

	quota, quotaErr := o.whitelistService.Quota(ctx, string(event.UserID), admin)
	if quotaErr != nil {
		log.Printf("load whitelist quota for qq %s: %v", event.UserID, quotaErr)
	}

	message := formatOneBotResult(command, result, err, quota, admin)
	o.sendGroupReply(parent, conn, sendMu, event.GroupID, message)
}

func formatOneBotResult(command *Command, result *OperationResult, err error, quota *QuotaStatus, admin bool) string {
	if err != nil {
		var detail string
		switch {
		case errors.Is(err, ErrUnavailable):
			detail = "whitelist service unavailable"
		case errors.Is(err, ErrQuotaExceeded):
			detail = "quota reached for this QQ account"
		case errors.Is(err, ErrConflict):
			detail = "this player is already owned by another QQ account; ask an admin to remove it first"
		case errors.Is(err, ErrForbidden):
			detail = "you can only remove whitelist entries linked to your QQ account"
		case errors.Is(err, ErrNotFound):
			detail = "no active whitelist entry was found for this player"
		case errors.Is(err, ErrInvalidInput):
			return oneBotCommandUsage
		default:
			detail = err.Error()
		}

		return appendQuotaLine(fmt.Sprintf("%s failed for %s (%s): %s.", commandVerb(command), command.PlayerName, command.Edition, detail), quota, admin)
	}

	if result != nil && !result.Changed && result.Message == "already whitelisted" {
		return appendQuotaLine(fmt.Sprintf("%s is already whitelisted for %s.", command.PlayerName, command.Edition), resultQuota(result, quota), admin)
	}

	var message string
	switch command.Action {
	case ActionAdd:
		message = fmt.Sprintf("Bound %s to the %s whitelist.", command.PlayerName, command.Edition)
	case ActionRemove:
		message = fmt.Sprintf("Unbound %s from the %s whitelist.", command.PlayerName, command.Edition)
	default:
		message = "Whitelist command completed."
	}
	return appendQuotaLine(message, resultQuota(result, quota), admin)
}

func commandVerb(command *Command) string {
	if command == nil {
		return "Command"
	}
	if command.Action == ActionRemove {
		return "Unbind"
	}
	return "Bind"
}

func resultQuota(result *OperationResult, fallback *QuotaStatus) *QuotaStatus {
	if result != nil && result.Quota != nil {
		return result.Quota
	}
	return fallback
}

func appendQuotaLine(message string, quota *QuotaStatus, admin bool) string {
	if admin {
		return message + "\nQuota: admin exempt."
	}
	if quota == nil {
		return message
	}
	return fmt.Sprintf("%s\nQuota: %d/%d used, %d remaining.", message, quota.Used, quota.Limit, quota.Remaining)
}

func (o *OneBotService) sendGroupReply(ctx context.Context, conn *websocket.Conn, sendMu *sync.Mutex, groupID oneBotID, message string) {
	if conn == nil {
		return
	}

	params := map[string]any{
		"group_id": idParam(groupID),
		"message":  message,
	}
	payload := map[string]any{
		"action": "send_group_msg",
		"params": params,
		"echo":   fmt.Sprintf("whitelist-%d", time.Now().UnixNano()),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("marshal onebot reply: %v", err)
		return
	}

	writeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	sendMu.Lock()
	defer sendMu.Unlock()
	if err := conn.Write(writeCtx, websocket.MessageText, data); err != nil {
		log.Printf("send onebot group reply: %v", err)
	}
}

func parseOneBotGroupEvent(data []byte) (oneBotGroupEvent, bool) {
	var event oneBotGroupEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return event, false
	}
	return event, event.PostType != ""
}

func idParam(id oneBotID) any {
	value := string(id)
	if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
		return parsed
	}
	return value
}

type oneBotGroupEvent struct {
	PostType    string       `json:"post_type"`
	MessageType string       `json:"message_type"`
	GroupID     oneBotID     `json:"group_id"`
	UserID      oneBotID     `json:"user_id"`
	RawMessage  string       `json:"raw_message"`
	Sender      oneBotSender `json:"sender"`
}

type oneBotSender struct {
	Role string `json:"role"`
}

type oneBotID string

func (id *oneBotID) UnmarshalJSON(data []byte) error {
	raw := strings.TrimSpace(string(data))
	if raw == "" || raw == "null" {
		*id = ""
		return nil
	}

	var text string
	if strings.HasPrefix(raw, `"`) {
		if err := json.Unmarshal(data, &text); err != nil {
			return err
		}
		*id = oneBotID(text)
		return nil
	}

	*id = oneBotID(raw)
	return nil
}
