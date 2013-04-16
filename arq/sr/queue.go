// Package sr implements a selective repeat ARQ protocol with a sliding window.
package sr

import (
	"errors"
	"fmt"
)

type Queue struct {
	contents []SequenceNumber

	windowSize int

	baseIndex               int
	nextSequenceNumberIndex int
}

type SequenceNumber struct {
	SequenceNumber int
	// Whether or not this sequence number has been sent in a packet
	Sent bool
	// Whether or not this sequence number has been sent and acknowledged
	Acknowledged bool
}

// NewQueue returns a Queue with the given window size at position zero
func NewQueue(windowSize int) *Queue {
	queue := new(Queue)

	queue.contents = make([]SequenceNumber, windowSize)
	queue.windowSize = windowSize
	queue.baseIndex = 0
	queue.nextSequenceNumberIndex = 0

	for i := range queue.contents {
		queue.contents[i].SequenceNumber = i
	}

	return queue
}

// OldestUnacknowledgedSequenceNumber returns the oldest unacknowledged sequence number
func (q *Queue) OldestUnacknowledgedSequenceNumber() (*SequenceNumber, error) {
	sequenceNumber := &q.contents[q.baseIndex]

	if !sequenceNumber.Sent {
		return nil, errors.New("Oldest sequence number has not been sent.")
	}

	return sequenceNumber, nil
}

// Send returns the next available sequence number and marks it as sent
func (q *Queue) Send() (int, error) {
	if q.nextSequenceNumberIndex-q.baseIndex == q.windowSize {
		return 0, errors.New("Window is full.")
	}

	var sequenceNumber *SequenceNumber

	sequenceNumber = &q.contents[q.nextSequenceNumberIndex]

	sequenceNumber.Sent = true

	q.nextSequenceNumberIndex++

	if q.nextSequenceNumberIndex >= cap(q.contents) {
		// Next Sequence Number Index is off the slice
		// Resize the slice

		lastSequenceNumber := sequenceNumber.SequenceNumber
		numberToAdd := cap(q.contents)

		for i := 1; i <= numberToAdd; i++ {
			number := (lastSequenceNumber + i)

			newSequenceNumber := SequenceNumber{number, false, false}

			q.contents = append(q.contents, newSequenceNumber)
		}
	}

	return sequenceNumber.SequenceNumber, nil
}

// MarkAckowledged marks the given sequence number as acknowledged if it is in the window
func (q *Queue) MarkAcknowledged(sequenceNumber int) error {
	if q.nextSequenceNumberIndex-q.baseIndex <= 0 {
		return errors.New("Window is empty")
	}

	for i := q.baseIndex; i < q.nextSequenceNumberIndex; i++ {
		if q.contents[i].SequenceNumber == sequenceNumber {
			q.contents[i].Acknowledged = true

			if i == q.baseIndex {
				q.slideWindow()
			}

			return nil
		}
	}

	return fmt.Errorf("Sequence number %d not found in window.", sequenceNumber)
}

// slideWindow slides the window until the first unacknowledged sequence number
func (q *Queue) slideWindow() {
	for q.contents[q.baseIndex].Acknowledged == true {
		q.baseIndex++
	}
}

func (q *Queue) String() string {
	output := ""

	for i, v := range q.contents {
		if q.baseIndex == i {
			output += "["
		}

		if i == q.baseIndex && q.nextSequenceNumberIndex == q.baseIndex {
			output += "]"
		}

		if v.Sent {
			if !v.Acknowledged {
				output += "-"
			} else if v.Acknowledged {
				output += "A"
			}
		} else {
			output += "_"
		}

		if q.baseIndex != i+1 && q.nextSequenceNumberIndex-1 == i {
			output += "]"
		}
	}

	return output
}
