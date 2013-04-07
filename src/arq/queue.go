package arq

import (
  "errors"
  "fmt"
)

type Queue struct {
  contents []SequenceNumber

  windowSize int

  baseIndex int
  nextSequenceNumberIndex int
}

type SequenceNumber struct {
  SequenceNumber int
  Sent bool
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
    queue.contents[i].SequenceNumber = i % (2 * windowSize)
  }

  return queue
}

// OldestUnacknowledgedSequenceNumber returns the oldest unacknowledged sequence number
func (q *Queue) OldestUnacknowledgedSequenceNumber() SequenceNumber {
  return q.contents[q.baseIndex]
}

// Send returns the next available sequence number and marks it as sent
func (q *Queue) Send() (int, error) {
  if q.nextSequenceNumberIndex - q.baseIndex == q.windowSize {
    return 0, errors.New("Window is full.")
  }

  var sequenceNumber *SequenceNumber

  sequenceNumber = &q.contents[q.nextSequenceNumberIndex]

  sequenceNumber.Sent = true

  q.nextSequenceNumberIndex++

  return sequenceNumber.SequenceNumber, nil
}

// MarkAckowledged marks the given sequence number as acknowledged if it is in the window
func (q *Queue) MarkAcknowledged(sequenceNumber int) error {
  if q.nextSequenceNumberIndex - q.baseIndex <= 0 {
    return errors.New("Window is empty")
  }

  for i := q.baseIndex; i < q.nextSequenceNumberIndex; i++ {
    if q.contents[i].SequenceNumber == sequenceNumber {
      q.contents[i].Acknowledged = true

      return nil
    }
  }

  return fmt.Errorf("Sequence number %d not found in window.", sequenceNumber)
}