package sr

import (
	"arq"
	"time"
)

// Time in seconds to wait to resend a packet after not receiving an acknowledgement from the receiver
const TimeoutTime = 5

// Round Trip Time (i.e. time for packet to be sent from sender to receiver plus acknowledgement time back) in milliseconds
const RoundTripTime = 200

type Computer struct {
	queue *Queue

	inputChan  chan arq.Packet
	outputChan chan arq.Packet

	waiting chan int
}

// NewComputer returns a initialized Computer struct given the windowSize and input/output channels
func NewComputer(windowSize int, inputChan, outputChan chan arq.Packet) *Computer {
	return &Computer{NewQueue(windowSize), inputChan, outputChan, make(chan int, windowSize)}
}

// Send returns the sequence number and error of the sent packet.
// Lose specifies whether or not that packet should be "lost" upon sending
// The sequence number is gotten from the queue
func (c *Computer) Send(lose bool) (int, error) {
	sequenceNumber, err := c.queue.Send()

	if err != nil {
		<-c.waiting

		return c.Send(lose)
	}

	return c.sendSequenceNumber(sequenceNumber, lose)
}

// sendSequenceNumber sends a packet of the desired sequence number with a time out
// Lose specifies whether or not that packet should be "lost" upon sending
func (c *Computer) sendSequenceNumber(sequenceNumber int, lose bool) (int, error) {
	timeoutTimer := time.AfterFunc(TimeoutTime*time.Second, func() {
		c.timeout(sequenceNumber)
	})

	packet := arq.Packet{sequenceNumber, false, 0, c.inputChan, timeoutTimer}

	// Don't actually send the packet if we're supposed to "lose" it
	if !lose {
		go func() {
			time.Sleep(RoundTripTime / 2 * time.Millisecond)

			c.outputChan <- packet
		}()
	}

	return sequenceNumber, nil
}

// Receive receives from the input channel, ACK's if necessary, and returns the packet
func (c *Computer) Receive() (arq.Packet, error) {
	packet := <-c.inputChan

	if packet.ACK {
		if err := c.queue.MarkAcknowledged(packet.ACKSequenceNumber); err != nil {
			return arq.Packet{}, err
		}

		c.waiting <- 1

		packet.TimeoutTimer.Stop()

		return packet, nil
	} else {
		go func() {
			time.Sleep(RoundTripTime / 2 * time.Millisecond)

			packet.ResponseChan <- arq.Packet{0, true, packet.SequenceNumber, c.inputChan, packet.TimeoutTimer}
		}()
	}

	return packet, nil
}

// timeout resends a packet with the given sequence number
func (c *Computer) timeout(sequenceNumber int) (int, error) {
	return c.sendSequenceNumber(sequenceNumber, false)
}
