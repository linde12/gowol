package gowol

import (
	"bytes"
	"testing"
)

const (
	MagicPacketSize = 102
)

func TestNewMagicPacket(t *testing.T) {
	packet, err := NewMagicPacket("AA:AA:AA:AA:AA:AA")
	// check valid mac
	if err != nil {
		t.Errorf("unable to create magic packet: %v", err)
		return
	}

	// check for magic padding
	if !bytes.Contains(packet[:6], []byte{255, 255, 255, 255, 255, 255}) {
		t.Error("packet doesn't contain 6 bytes of 0xFF padding")
		return
	}

	// eqv. of AA:AA:AA:AA:AA:AA
	rawMac := []byte{170, 170, 170, 170, 170, 170}
	pos := 6
	for i := 0; i < 16; i++ {
		if !bytes.Equal(packet[pos:pos+6], rawMac) {
			t.Error("magic packet contains wrong mac")
			return
		}
	}
}

func TestEui48(t *testing.T) {
	// only support EUI-48 addresses, EUI-64 should fail
	_, err := NewMagicPacket("AA:BB:CC:AA:BB:CC:AA:BB")
	if err == nil {
		t.Error("able to construct magic packet with invalid MAC")
	}

	_, err = NewMagicPacket("AA:BB:CC:AA:BB:CC")
	if err != nil {
		t.Error(err)
	}
}
