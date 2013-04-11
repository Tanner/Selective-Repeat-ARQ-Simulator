package sr

import "arq"

type Computer struct {
	queue *Queue

	inputChan  chan arq.Packet
	outputChan chan arq.Packet
}

// NewComputer returns a initialized Computer struct given the windowSize and input/output channels
func NewComputer(windowSize int, inputChan, outputChan chan arq.Packet) *Computer {
	return &Computer{NewQueue(windowSize), inputChan, outputChan}
}

// Send returns the sequence number of the packet that was sent on the output channel using Selective Repeat protocol
func (c *Computer) Send() (int, error) {
	sequenceNumber, err := c.queue.Send()

	if err != nil {
		return 0, err
	}

	c.outputChan <- arq.Packet{sequenceNumber, false, c.inputChan}

	return sequenceNumber, nil
}

// Receive receives from the input channel, ACK's if necessary, and returns the packet
func (c *Computer) Receive() (arq.Packet, error) {
	packet := <-c.inputChan

	if packet.ACK {
		if err := c.queue.MarkAcknowledged(packet.SequenceNumber); err != nil {
			return arq.Packet{}, err
		}

		return packet, nil
	} else {
		packet.ResponseChan <- arq.Packet{0, true, c.inputChan}
	}

	return packet, nil
}
