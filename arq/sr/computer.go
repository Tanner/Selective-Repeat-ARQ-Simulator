package sr

import (
	"arq"
	"time"
)

type Computer struct {
	queue *Queue

	inputChan  chan arq.Packet
	outputChan chan arq.Packet

	waiting chan int

	roundTripDuration time.Duration
	timeoutDuration   time.Duration

	timeoutTriggered func(int)
}

// NewComputer returns a initialized Computer struct given the windowSize and input/output channels
func NewComputer(windowSize int, inputChan, outputChan chan arq.Packet, roundTripDuration, timeoutDuration time.Duration, timeout func(int)) *Computer {
	return &Computer{NewQueue(windowSize), inputChan, outputChan, make(chan int, windowSize), roundTripDuration, timeoutDuration, timeout}
}

// Send returns the sequence number and error of the sent packet.
// Lose specifies whether or not that packet should be "lost" upon sending
// The sequence number is gotten from the queue
func (c *Computer) Send(senderLose, acknowledgementLose bool) (int, error) {
	sequenceNumber, err := c.queue.Send()

	if err != nil {
		<-c.waiting

		return c.Send(senderLose, acknowledgementLose)
	}

	return c.sendSequenceNumber(sequenceNumber, senderLose, acknowledgementLose)
}

// sendSequenceNumber sends a packet of the desired sequence number with a time out
// Lose specifies whether or not that packet should be "lost" upon sending/ACK
func (c *Computer) sendSequenceNumber(sequenceNumber int, senderLose, acknowledgementLose bool) (int, error) {
	// Set up the timeout function
	timeoutTimer := time.AfterFunc(c.timeoutDuration, func() {
		c.timeoutTriggered(sequenceNumber)

		if !senderLose && acknowledgementLose {
			c.timeout(sequenceNumber, false)
		} else {
			c.timeout(sequenceNumber, acknowledgementLose)
		}
	})

	packet := arq.Packet{sequenceNumber, false, 0, acknowledgementLose, c.inputChan, timeoutTimer}

	// Don't actually send the packet if we're supposed to "lose" it
	if !senderLose {
		go func() {
			time.Sleep(c.roundTripDuration / 2)

			c.outputChan <- packet
		}()
	}

	return sequenceNumber, nil
}

// Receive receives from the input channel, ACK's if necessary, and returns the packet
func (c *Computer) Receive() (arq.Packet, error) {
	packet := <-c.inputChan

	if packet.ACK {
		// If this packet is an acknowledement, mark it as acknowledged in the queue
		if err := c.queue.MarkAcknowledged(packet.ACKSequenceNumber); err != nil {
			return arq.Packet{}, err
		}

		// Inform any waiting packets that there is a spot open
		for i := 0; i < c.queue.FreeSpace(); i++ {
			select {
			case c.waiting <- 1:
			default:
			}
		}

		packet.TimeoutTimer.Stop()

		return packet, nil
	} else if !packet.AcknowledgementLoss {
		// If we're not supposed to "lose" the packet, send an acknowledgement that we received it
		go func() {
			time.Sleep(c.roundTripDuration / 2)

			packet.ResponseChan <- arq.Packet{0, true, packet.SequenceNumber, false, c.inputChan, packet.TimeoutTimer}
		}()
	}

	return packet, nil
}

// timeout resends a packet with the given sequence number
func (c *Computer) timeout(sequenceNumber int, acknowledgementLose bool) (int, error) {
	return c.sendSequenceNumber(sequenceNumber, false, acknowledgementLose)
}

func (c *Computer) QueueString() string {
	return c.queue.String()
}
