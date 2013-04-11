package sr

import (
	"arq"
	"time"
)

const TimeoutTime = 5
const SleepTime = 1

type Computer struct {
	queue *Queue

	inputChan  chan arq.Packet
	outputChan chan arq.Packet
}

// NewComputer returns a initialized Computer struct given the windowSize and input/output channels
func NewComputer(windowSize int, inputChan, outputChan chan arq.Packet) *Computer {
	return &Computer{NewQueue(windowSize), inputChan, outputChan}
}

// Send returns the sequence number and error of the sent packet.
// The sequence number is gotten from the queue
func (c *Computer) Send() (int, error) {
	sequenceNumber, err := c.queue.Send()

	if err != nil {
		return 0, err
	}

	return c.sendSequenceNumber(sequenceNumber)
}

// sendSequenceNumber sends a packet of the desired sequence number with a time out
func (c *Computer) sendSequenceNumber(sequenceNumber int) (int, error) {
	timeoutTimer := time.AfterFunc(TimeoutTime*time.Second, func() {
		c.timeout(sequenceNumber)
	})

	packet := arq.Packet{sequenceNumber, false, 0, c.inputChan, timeoutTimer}

	c.outputChan <- packet

	time.Sleep(SleepTime * time.Millisecond)

	return sequenceNumber, nil
}

// Receive receives from the input channel, ACK's if necessary, and returns the packet
func (c *Computer) Receive() (arq.Packet, error) {
	packet := <-c.inputChan

	time.Sleep(SleepTime * time.Millisecond)

	if packet.ACK {
		if err := c.queue.MarkAcknowledged(packet.ACKSequenceNumber); err != nil {
			return arq.Packet{}, err
		}

		packet.TimeoutTimer.Stop()

		return packet, nil
	} else {
		packet.ResponseChan <- arq.Packet{0, true, packet.SequenceNumber, c.inputChan, packet.TimeoutTimer}

		time.Sleep(SleepTime * time.Millisecond)
	}

	return packet, nil
}

// timeout resends a packet with the given sequence number
func (c *Computer) timeout(sequenceNumber int) (int, error) {
	return c.sendSequenceNumber(sequenceNumber)
}
