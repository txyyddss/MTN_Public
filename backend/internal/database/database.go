// Package database provides MySQL integration for McMMO and Floodgate.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mtn-server/backend/config"
)

// McmmoSkills holds a player's McMMO skill levels.
type McmmoSkills struct {
	UserID      int    `json:"user_id"`
	User        string `json:"user"`
	UUID        string `json:"uuid"`
	Taming      int    `json:"taming"`
	Mining      int    `json:"mining"`
	Woodcutting int    `json:"woodcutting"`
	Repair      int    `json:"repair"`
	Unarmed     int    `json:"unarmed"`
	Herbalism   int    `json:"herbalism"`
	Excavation  int    `json:"excavation"`
	Archery     int    `json:"archery"`
	Swords      int    `json:"swords"`
	Axes        int    `json:"axes"`
	Acrobatics  int    `json:"acrobatics"`
	Fishing     int    `json:"fishing"`
	Alchemy     int    `json:"alchemy"`
	Crossbows   int    `json:"crossbows"`
	Tridents    int    `json:"tridents"`
	Maces       int    `json:"maces"`
	Spears      int    `json:"spears"`
	Total       int    `json:"total"`
}

// LinkedPlayer holds a Floodgate account link.
type LinkedPlayer struct {
	BedrockUUID     string `json:"bedrock_uuid"`
	BedrockUsername string `json:"bedrock_username,omitempty"`
	JavaUUID        string `json:"java_uuid"`
	JavaUsername    string `json:"java_username"`
}

// McmmoDB provides access to the McMMO MySQL database.
type McmmoDB struct {
	db *sql.DB
}

// FloodgateDB provides access to the Floodgate MySQL database.
type FloodgateDB struct {
	db *sql.DB
}

// NewMcmmoDB opens a connection to the McMMO database.
func NewMcmmoDB(cfg config.MySQLConfig) (*McmmoDB, error) {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open mcmmo db: %w", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mcmmo db: %w", err)
	}
	return &McmmoDB{db: db}, nil
}

// GetAllSkills returns McMMO skills for all players.
func (m *McmmoDB) GetAllSkills(ctx context.Context) ([]McmmoSkills, error) {
	query := `
		SELECT u.id, u.user, u.uuid,
		       s.taming, s.mining, s.woodcutting, s.repair, s.unarmed,
		       s.herbalism, s.excavation, s.archery, s.swords, s.axes,
		       s.acrobatics, s.fishing, s.alchemy, s.crossbows, s.tridents,
		       s.maces, s.spears, s.total
		FROM mcmmo_users u
		JOIN mcmmo_skills s ON u.id = s.user_id
	`
	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query mcmmo skills: %w", err)
	}
	defer rows.Close()

	var result []McmmoSkills
	for rows.Next() {
		var s McmmoSkills
		if err := rows.Scan(
			&s.UserID, &s.User, &s.UUID,
			&s.Taming, &s.Mining, &s.Woodcutting, &s.Repair, &s.Unarmed,
			&s.Herbalism, &s.Excavation, &s.Archery, &s.Swords, &s.Axes,
			&s.Acrobatics, &s.Fishing, &s.Alchemy, &s.Crossbows, &s.Tridents,
			&s.Maces, &s.Spears, &s.Total,
		); err != nil {
			return nil, fmt.Errorf("scan mcmmo skills: %w", err)
		}
		result = append(result, s)
	}
	return result, rows.Err()
}

// GetSkillsByUUID returns skills for a specific player.
func (m *McmmoDB) GetSkillsByUUID(ctx context.Context, uuid string) (*McmmoSkills, error) {
	query := `
		SELECT u.id, u.user, u.uuid,
		       s.taming, s.mining, s.woodcutting, s.repair, s.unarmed,
		       s.herbalism, s.excavation, s.archery, s.swords, s.axes,
		       s.acrobatics, s.fishing, s.alchemy, s.crossbows, s.tridents,
		       s.maces, s.spears, s.total
		FROM mcmmo_users u
		JOIN mcmmo_skills s ON u.id = s.user_id
		WHERE u.uuid = ?
	`
	s := &McmmoSkills{}
	err := m.db.QueryRowContext(ctx, query, uuid).Scan(
		&s.UserID, &s.User, &s.UUID,
		&s.Taming, &s.Mining, &s.Woodcutting, &s.Repair, &s.Unarmed,
		&s.Herbalism, &s.Excavation, &s.Archery, &s.Swords, &s.Axes,
		&s.Acrobatics, &s.Fishing, &s.Alchemy, &s.Crossbows, &s.Tridents,
		&s.Maces, &s.Spears, &s.Total,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query mcmmo skills by uuid: %w", err)
	}
	return s, nil
}

// Close shuts down the database connection.
func (m *McmmoDB) Close() error {
	return m.db.Close()
}

// NewFloodgateDB opens a connection to the Floodgate database.
func NewFloodgateDB(cfg config.MySQLConfig) (*FloodgateDB, error) {
	db, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("open floodgate db: %w", err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping floodgate db: %w", err)
	}
	return &FloodgateDB{db: db}, nil
}

// GetLinkedByJavaUUID finds a linked Bedrock account by Java UUID.
func (f *FloodgateDB) GetLinkedByJavaUUID(ctx context.Context, javaUUID string) (*LinkedPlayer, error) {
	query := `SELECT HEX(bedrockId), HEX(javaUniqueId), javaUsername FROM LinkedPlayers WHERE HEX(javaUniqueId) = ?`
	// Convert UUID string (with dashes) to hex (no dashes)
	hexUUID := uuidToHex(javaUUID)

	lp := &LinkedPlayer{}
	err := f.db.QueryRowContext(ctx, query, hexUUID).Scan(&lp.BedrockUUID, &lp.JavaUUID, &lp.JavaUsername)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query linked player: %w", err)
	}
	lp.BedrockUUID = hexToUUID(lp.BedrockUUID)
	lp.JavaUUID = hexToUUID(lp.JavaUUID)
	return lp, nil
}

// GetLinkedByBedrockUUID finds a linked Java account by Bedrock UUID.
func (f *FloodgateDB) GetLinkedByBedrockUUID(ctx context.Context, bedrockUUID string) (*LinkedPlayer, error) {
	query := `SELECT HEX(bedrockId), HEX(javaUniqueId), javaUsername FROM LinkedPlayers WHERE HEX(bedrockId) = ?`
	hexUUID := uuidToHex(bedrockUUID)

	lp := &LinkedPlayer{}
	err := f.db.QueryRowContext(ctx, query, hexUUID).Scan(&lp.BedrockUUID, &lp.JavaUUID, &lp.JavaUsername)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query linked player: %w", err)
	}
	lp.BedrockUUID = hexToUUID(lp.BedrockUUID)
	lp.JavaUUID = hexToUUID(lp.JavaUUID)
	return lp, nil
}

// Close shuts down the database connection.
func (f *FloodgateDB) Close() error {
	return f.db.Close()
}

func uuidToHex(uuid string) string {
	return fmt.Sprintf("%s", removeHyphens(uuid))
}

func removeHyphens(s string) string {
	result := make([]byte, 0, 32)
	for i := 0; i < len(s); i++ {
		if s[i] != '-' {
			result = append(result, s[i])
		}
	}
	return string(result)
}

func hexToUUID(hex string) string {
	hex = removeHyphens(hex)
	if len(hex) < 32 {
		// Pad with leading zeros
		for len(hex) < 32 {
			hex = "0" + hex
		}
	}
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex[0:8], hex[8:12], hex[12:16], hex[16:20], hex[20:32])
}
