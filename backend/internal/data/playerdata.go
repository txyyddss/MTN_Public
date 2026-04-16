package data

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// NBT tag types
const (
	tagEnd       = 0
	tagByte      = 1
	tagShort     = 2
	tagInt       = 3
	tagLong      = 4
	tagFloat     = 5
	tagDouble    = 6
	tagByteArray = 7
	tagString    = 8
	tagList      = 9
	tagCompound  = 10
	tagIntArray  = 11
	tagLongArray = 12
)

// nbtReader reads NBT data from a binary stream.
type nbtReader struct {
	r   io.Reader
	err error
}

func (n *nbtReader) readByte() byte {
	if n.err != nil {
		return 0
	}
	var b [1]byte
	_, n.err = io.ReadFull(n.r, b[:])
	return b[0]
}

func (n *nbtReader) readShort() int16 {
	if n.err != nil {
		return 0
	}
	var v int16
	n.err = binary.Read(n.r, binary.BigEndian, &v)
	return v
}

func (n *nbtReader) readInt() int32 {
	if n.err != nil {
		return 0
	}
	var v int32
	n.err = binary.Read(n.r, binary.BigEndian, &v)
	return v
}

func (n *nbtReader) readLong() int64 {
	if n.err != nil {
		return 0
	}
	var v int64
	n.err = binary.Read(n.r, binary.BigEndian, &v)
	return v
}

func (n *nbtReader) readFloat() float32 {
	if n.err != nil {
		return 0
	}
	var v float32
	n.err = binary.Read(n.r, binary.BigEndian, &v)
	return v
}

func (n *nbtReader) readDouble() float64 {
	if n.err != nil {
		return 0
	}
	var v float64
	n.err = binary.Read(n.r, binary.BigEndian, &v)
	return v
}

func (n *nbtReader) readString() string {
	if n.err != nil {
		return ""
	}
	length := n.readShort()
	if n.err != nil || length <= 0 {
		return ""
	}
	buf := make([]byte, length)
	_, n.err = io.ReadFull(n.r, buf)
	return string(buf)
}

// readTag reads the tag type and name.
func (n *nbtReader) readTag() (tagType byte, name string) {
	tagType = n.readByte()
	if tagType == tagEnd || n.err != nil {
		return tagType, ""
	}
	name = n.readString()
	return
}

// skipPayload skips over a tag's payload without storing it.
func (n *nbtReader) skipPayload(tagType byte) {
	if n.err != nil {
		return
	}
	switch tagType {
	case tagByte:
		n.readByte()
	case tagShort:
		n.readShort()
	case tagInt:
		n.readInt()
	case tagLong:
		n.readLong()
	case tagFloat:
		n.readFloat()
	case tagDouble:
		n.readDouble()
	case tagByteArray:
		length := n.readInt()
		if length > 0 {
			buf := make([]byte, length)
			_, n.err = io.ReadFull(n.r, buf)
		}
	case tagString:
		n.readString()
	case tagList:
		listType := n.readByte()
		count := n.readInt()
		for i := int32(0); i < count && n.err == nil; i++ {
			n.skipPayload(listType)
		}
	case tagCompound:
		for n.err == nil {
			childType, _ := n.readTag()
			if childType == tagEnd {
				break
			}
			n.skipPayload(childType)
		}
	case tagIntArray:
		length := n.readInt()
		for i := int32(0); i < length && n.err == nil; i++ {
			n.readInt()
		}
	case tagLongArray:
		length := n.readInt()
		for i := int32(0); i < length && n.err == nil; i++ {
			n.readLong()
		}
	}
}

// parsePlayerDat reads a gzip-compressed NBT .dat file and extracts key fields.
func parsePlayerDat(filePath string) (*PlayerInfo, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("gzip reader: %w", err)
	}
	defer gz.Close()

	decompressed, err := io.ReadAll(gz)
	if err != nil {
		return nil, fmt.Errorf("decompress: %w", err)
	}

	reader := &nbtReader{r: bytes.NewReader(decompressed)}
	info := &PlayerInfo{}

	// Read root tag
	rootType, _ := reader.readTag()
	if rootType != tagCompound {
		return nil, fmt.Errorf("expected compound root, got %d", rootType)
	}

	// Parse compound children
	parseCompound(reader, info, "")

	if reader.err != nil {
		return nil, fmt.Errorf("nbt parse: %w", reader.err)
	}

	return info, nil
}

// parseCompound reads a compound tag and extracts known fields.
func parseCompound(reader *nbtReader, info *PlayerInfo, prefix string) {
	for reader.err == nil {
		tagType, name := reader.readTag()
		if tagType == tagEnd {
			return
		}

		fullName := name
		if prefix != "" {
			fullName = prefix + "." + name
		}

		switch fullName {
		case "XpLevel":
			if tagType == tagInt {
				info.XpLevel = int(reader.readInt())
				continue
			}
		case "foodLevel":
			if tagType == tagInt {
				info.FoodLevel = int(reader.readInt())
				continue
			}
		case "Health":
			if tagType == tagFloat {
				info.Health = float64(reader.readFloat())
				continue
			}
		case "Spigot.ticksLived":
			if tagType == tagInt {
				info.TicksLived = int64(reader.readInt())
				continue
			}
		case "Paper.LastSeen":
			if tagType == tagLong {
				info.LastSeen = reader.readLong()
				continue
			}
		case "bukkit.firstPlayed":
			if tagType == tagLong {
				info.FirstPlayed = reader.readLong()
				continue
			}
		case "bukkit.lastKnownName":
			if tagType == tagString {
				info.LastKnownName = reader.readString()
				continue
			}
		}

		// Recurse into compounds to find nested fields
		if tagType == tagCompound {
			switch name {
			case "Spigot", "Paper", "bukkit":
				parseCompound(reader, info, name)
				continue
			}
		}

		// Skip everything else
		reader.skipPayload(tagType)
	}
}

// loadPlayerData loads all .dat files from worldDir/playerdata.
func (s *Store) loadPlayerData() error {
	dir := filepath.Join(s.worldDir, "playerdata")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read playerdata dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".dat") {
			continue
		}

		uuid := strings.TrimSuffix(entry.Name(), ".dat")
		info, err := parsePlayerDat(filepath.Join(dir, entry.Name()))
		if err != nil {
			// Skip malformed files
			continue
		}
		info.UUID = uuid

		if strings.HasPrefix(uuid, "00000000-0000-0000") {
			info.Type = "Bedrock"
			clean := info.LastKnownName
			if strings.HasPrefix(clean, ".") {
				clean = clean[1:]
			} else if strings.HasPrefix(clean, "BE_") {
				clean = clean[3:]
			}
			info.LastKnownName = clean
			info.CleanName = clean
		} else {
			info.Type = "Java"
			info.CleanName = info.LastKnownName
		}

		s.Players[uuid] = info
	}
	return nil
}
