package gowol

import (
	"encoding/hex"
	"net"
	"strings"
)

type MagicPacket [102]byte

// NewMagicPacket allocates a new MagicPacket with the specified MAC.
func NewMagicPacket(macAddr string) (packet MagicPacket, err error) {
	mac, err := hex.DecodeString(strings.Join(strings.Split(macAddr, ":"), ""))
	if err != nil {
		return packet, err
	}

	// write magic bytes to packet
	copy(packet[0:], []byte{255, 255, 255, 255, 255, 255})
	offset := 6

	for i := 0; i < 16; i++ {
		copy(packet[offset:], mac)
		offset += 6
	}

	return packet, nil
}

func sendUdpPacket(mp MagicPacket, addr string) (err error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(mp[:])
	return err
}

// Send writes the MagicPacket to the specified address on port 9.
func (mp MagicPacket) Send(addr string) (err error) {
	return sendUdpPacket(mp, addr+":9")
}

// SendPort writes the MagicPacket to the specified address and port.
func (mp MagicPacket) SendPort(addr string, port string) (err error) {
	return sendUdpPacket(mp, addr+":"+port)
}
