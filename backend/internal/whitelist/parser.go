package whitelist

import (
	"strings"
)

// Command is a parsed OneBot group command.
type Command struct {
	Action     Action
	Edition    Edition
	PlayerName string
}

// ParseCommand parses exact bind and unbind group command forms.
func ParseCommand(rawMessage string) (*Command, bool, error) {
	fields := strings.Fields(normalizeOneBotCommandText(rawMessage))
	if len(fields) == 0 {
		return nil, false, nil
	}

	var action Action
	switch normalizeCommandToken(fields[0]) {
	case "bind":
		action = ActionAdd
	case "unbind":
		action = ActionRemove
	case "whitelist", "fwhitelist":
		return nil, true, ErrInvalidInput
	default:
		return nil, false, nil
	}

	if len(fields) != 3 {
		return nil, true, ErrInvalidInput
	}

	edition, err := NormalizeEdition(fields[1])
	if err != nil || edition == EditionAll {
		return nil, true, ErrInvalidInput
	}

	playerName := strings.TrimSpace(fields[2])
	if err := ValidatePlayerName(playerName); err != nil {
		return nil, true, err
	}

	return &Command{
		Action:     action,
		Edition:    edition,
		PlayerName: playerName,
	}, true, nil
}

func normalizeOneBotCommandText(rawMessage string) string {
	message := strings.TrimSpace(rawMessage)
	for {
		if !strings.HasPrefix(message, "[CQ:at,") {
			return message
		}

		end := strings.IndexByte(message, ']')
		if end < 0 {
			return message
		}
		message = strings.TrimSpace(message[end+1:])
	}
}

func normalizeCommandToken(token string) string {
	return strings.ToLower(strings.TrimLeft(strings.TrimSpace(token), "/!#"))
}

// ValidatePlayerName checks that a player token is safe to place in an RCON command.
func ValidatePlayerName(playerName string) error {
	name := strings.TrimSpace(playerName)
	if name == "" || len(name) > 32 {
		return ErrInvalidInput
	}

	for _, value := range name {
		switch {
		case value >= 'a' && value <= 'z':
		case value >= 'A' && value <= 'Z':
		case value >= '0' && value <= '9':
		case value == '_' || value == '-' || value == '.':
		default:
			return ErrInvalidInput
		}
	}

	return nil
}
