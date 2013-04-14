// Package arq provides structs required for ARQ protocols.
package arq

import (
	"fmt"
	"time"
)

type Packet struct {
	SequenceNumber int

	ACK               bool
	ACKSequenceNumber int

	// When this packet is being acknowledged, "lose" the ACK packet in the network
	AcknowledgementLoss bool

	ResponseChan chan Packet
	TimeoutTimer *time.Timer
}

// String returns a human-readable string of the Packet
func (p Packet) String() string {
	string := fmt.Sprintf("Packet #%d", p.SequenceNumber)

	if p.ACK {
		string += fmt.Sprintf(" - ACK for %d", p.ACKSequenceNumber)
	}

	return string
}
