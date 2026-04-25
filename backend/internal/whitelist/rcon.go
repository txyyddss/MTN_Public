package whitelist

import (
	"context"
	"fmt"
	"time"

	"github.com/mcstatus-io/mcutil/v4/options"
	"github.com/mcstatus-io/mcutil/v4/rcon"

	"github.com/mtn-server/backend/config"
)

// RCONExecutor executes Minecraft whitelist commands through mcutil RCON.
type RCONExecutor struct {
	host     string
	port     uint16
	password string
	timeout  time.Duration
}

// NewRCONExecutor creates an RCON executor from config.
func NewRCONExecutor(cfg config.RCONConfig) *RCONExecutor {
	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second
	if timeout <= 0 {
		timeout = 5 * time.Second
	}

	return &RCONExecutor{
		host:     cfg.Host,
		port:     cfg.Port,
		password: cfg.Password,
		timeout:  timeout,
	}
}

// Execute sends an edition-specific whitelist command and waits for the response.
func (r *RCONExecutor) Execute(ctx context.Context, edition Edition, action Action, playerName string) (string, error) {
	if r == nil || r.host == "" || r.port == 0 || r.password == "" {
		return "", ErrUnavailable
	}

	command, err := buildRCONCommand(edition, action, playerName)
	if err != nil {
		return "", err
	}

	client, err := rcon.Dial(r.host, r.port, options.RCON{Timeout: r.timeout})
	if err != nil {
		return "", fmt.Errorf("dial rcon: %w", err)
	}
	defer client.Close()

	if err := client.Login(r.password); err != nil {
		return "", fmt.Errorf("login rcon: %w", err)
	}

	execCtx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	output, err := client.Execute(execCtx, command)
	if err != nil {
		return output, fmt.Errorf("execute rcon command: %w", err)
	}

	return output, nil
}

func buildRCONCommand(edition Edition, action Action, playerName string) (string, error) {
	if err := ValidatePlayerName(playerName); err != nil {
		return "", err
	}
	if action != ActionAdd && action != ActionRemove {
		return "", ErrInvalidInput
	}

	switch edition {
	case EditionJava:
		return fmt.Sprintf("whitelist %s %s", action, playerName), nil
	case EditionBedrock:
		return fmt.Sprintf("fwhitelist %s %s", action, playerName), nil
	default:
		return "", ErrInvalidInput
	}
}
