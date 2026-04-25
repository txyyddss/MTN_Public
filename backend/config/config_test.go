package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadWhitelistDefaults(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(`{}`), 0o600); err != nil {
		t.Fatalf("write temp config: %v", err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if cfg.Whitelist.MaxPerQQ != 3 {
		t.Fatalf("MaxPerQQ = %d, want 3", cfg.Whitelist.MaxPerQQ)
	}
	if cfg.RCON.TimeoutSeconds != 5 {
		t.Fatalf("RCON timeout = %d, want 5", cfg.RCON.TimeoutSeconds)
	}
	if cfg.OneBot.ReconnectSeconds != 5 {
		t.Fatalf("OneBot reconnect = %d, want 5", cfg.OneBot.ReconnectSeconds)
	}
}
