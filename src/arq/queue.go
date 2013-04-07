package arq

import "errors"

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
  queue.nextSequenceNumberIndex = 1

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

  return sequenceNumber.SequenceNumber, nil
}

func (q *Queue) 