package arq

import (
	"testing"
)

func TestString(t *testing.T) {
	packet := Packet{1261, false, 0, false, nil, nil}

	compareString := func(p Packet, s string) {
		if p.String() != s {
			t.Errorf("Packet String() does not match expected value \"%s\", got \"%s\"", s, p.String())
		}
	}

	compareString(packet, "Packet #1261")

	packet.ACK = true

	compareString(packet, "Packet #1261 - ACK for 0")
}
