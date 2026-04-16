package data

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestNBTReader(t *testing.T) {
	buf := new(bytes.Buffer)

	// Write a String tag (Type 8) with name "Test" and value "Hello"
	buf.WriteByte(8)                              // Tag Type: String
	binary.Write(buf, binary.BigEndian, int16(4)) // Name length
	buf.WriteString("Test")                       // Name
	binary.Write(buf, binary.BigEndian, int16(5)) // Value length
	buf.WriteString("Hello")                      // Value

	// Write an Int tag (Type 3) with name "Level" and value 10
	buf.WriteByte(3)                               // Tag Type: Int
	binary.Write(buf, binary.BigEndian, int16(5))  // Name length
	buf.WriteString("Level")                       // Name
	binary.Write(buf, binary.BigEndian, int32(10)) // Value

	// Write End tag (Type 0)
	buf.WriteByte(0)

	reader := &nbtReader{r: buf}

	// Read first tag
	tagType, name := reader.readTag()
	if tagType != 8 || name != "Test" {
		t.Errorf("expected tag type 8 and name Test, got %d and %s", tagType, name)
	}
	val := reader.readString()
	if val != "Hello" {
		t.Errorf("expected value Hello, got %s", val)
	}

	// Read second tag
	tagType, name = reader.readTag()
	if tagType != 3 || name != "Level" {
		t.Errorf("expected tag type 3 and name Level, got %d and %s", tagType, name)
	}
	level := reader.readInt()
	if level != 10 {
		t.Errorf("expected value 10, got %d", level)
	}

	// Read end tag
	tagType, _ = reader.readTag()
	if tagType != 0 {
		t.Errorf("expected tag type 0, got %d", tagType)
	}
}

func TestParseCompound(t *testing.T) {
	buf := new(bytes.Buffer)

	// Create a compound with XpLevel and bukkit.lastKnownName
	// Root compound doesn't have a name in our parseCompound's loop (it expects to be inside the compound)

	// XpLevel (Int)
	buf.WriteByte(3)
	binary.Write(buf, binary.BigEndian, int16(7))
	buf.WriteString("XpLevel")
	binary.Write(buf, binary.BigEndian, int32(42))

	// bukkit (Compound)
	buf.WriteByte(10)
	binary.Write(buf, binary.BigEndian, int16(6))
	buf.WriteString("bukkit")

	// Inside bukkit: lastKnownName (String)
	buf.WriteByte(8)
	binary.Write(buf, binary.BigEndian, int16(13))
	buf.WriteString("lastKnownName")
	binary.Write(buf, binary.BigEndian, int16(7))
	buf.WriteString("Antigrv")

	// End of bukkit
	buf.WriteByte(0)

	// End of root
	buf.WriteByte(0)

	reader := &nbtReader{r: buf}
	info := &PlayerInfo{}

	parseCompound(reader, info, "")

	if info.XpLevel != 42 {
		t.Errorf("expected XpLevel 42, got %d", info.XpLevel)
	}
	if info.LastKnownName != "Antigrv" {
		t.Errorf("expected LastKnownName Antigrv, got %s", info.LastKnownName)
	}
}
