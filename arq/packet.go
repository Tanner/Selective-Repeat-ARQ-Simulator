// Package arq provides structs required for ARQ protocols.
package arq

import "fmt"

type Packet struct {
	SequenceNumber int
	ACK            bool

	ResponseChan chan Packet
}

// String returns a human-readable string of the Packet
func (p Packet) String() string {
	string := fmt.Sprintf("#%d", p.SequenceNumber)

	if p.ACK {
		string += " - ACK"
	}

	return string
}
