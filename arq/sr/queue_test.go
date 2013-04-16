package sr

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	windowSize := 8

	maxSequenceNumber := 2*windowSize - 1

	queue := NewQueue(windowSize)

	if queue.windowSize != windowSize {
		t.Errorf("Queue does not have correct window size: %d expected %d", queue.windowSize, windowSize)
	}

	if queue.baseIndex == 0 && queue.nextSequenceNumberIndex != 0 {
		t.Error("Queue does not start with an empty window")
	}

	for _, sequenceNumber := range queue.contents {
		if sequenceNumber.Sent == true {
			t.Error("Sent for a sequence number is true, expected false")
		}

		if sequenceNumber.Acknowledged == true {
			t.Error("Acknowledged for a sequence number is true, expected false")
		}

		if sequenceNumber.SequenceNumber > maxSequenceNumber {
			t.Errorf("Sequence number is %d, expected something less than %d", sequenceNumber.SequenceNumber, maxSequenceNumber)
		}
	}
}

func TestOldestUnacknowledgedSequenceNumber(t *testing.T) {
	queue := NewQueue(8)

	sequenceNumber, error := queue.OldestUnacknowledgedSequenceNumber()

	if error == nil {
		t.Error("Error not thrown when no sequence numbers have been sent.")
	}

	queue.Send()

	sequenceNumber, error = queue.OldestUnacknowledgedSequenceNumber()

	if sequenceNumber.SequenceNumber != 0 {
		t.Errorf("Oldest sequence number is %d, expected %d", sequenceNumber.SequenceNumber, 0)
	}
}

func TestSend(t *testing.T) {
	queue := NewQueue(8)

	for i := 0; i < queue.windowSize; i++ {
		sequenceNumber, error := queue.Send()

		if error != nil {
			t.Errorf("Send returned error on non-full window: %s", error)
		}

		if sequenceNumber != i {
			t.Errorf("Send did not return correct sequence number (%d), expected %d", sequenceNumber, i)
		}
	}

	_, error := queue.Send()

	if error == nil {
		t.Error("Send failed to error on full window")
	}
}

func TestSendResize(t *testing.T) {
	queue := NewQueue(2)

	previousCap := cap(queue.contents)

	queue.Send()
	queue.Send()

	if cap(queue.contents) == previousCap {
		t.Error("Queue cap did not change before exceeding size")
	}

	for i := 2; i < cap(queue.contents); i++ {
		if queue.contents[i].Sent {
			t.Error("Expected new sequence nubmers to not have been sent")
		}

		if queue.contents[i].Acknowledged {
			t.Error("Expected new sequence nubmers to not have been acknowledged")
		}
	}
}

func TestMarkAcknowledged(t *testing.T) {
	queue := NewQueue(8)

	error := queue.MarkAcknowledged(0)

	if error == nil {
		t.Error("Marked a sequence number with no sent sequence numbers")
	}

	queue.Send()

	error = queue.MarkAcknowledged(0)

	if error != nil {
		t.Errorf("Expected marking a sent packet to go without error: %s", error)
	}

	for i := queue.baseIndex; i < queue.nextSequenceNumberIndex; i++ {
		if queue.contents[i].SequenceNumber == 0 {
			if !queue.contents[i].Acknowledged {
				t.Error("Acknowledged field on sequence number is not true")
			}
		}
	}

	error = queue.MarkAcknowledged(50)

	if error == nil {
		t.Error("Marked a seuqnece number outside of window")
	}
}

func TestSliding(t *testing.T) {
	queue := NewQueue(8)

	sequenceNumber, _ := queue.Send()
	queue.MarkAcknowledged(sequenceNumber)

	if queue.baseIndex == 0 {
		t.Error("Base Index did not increment")
	}

	if queue.contents[queue.baseIndex].Acknowledged {
		t.Error("Window did not slide past acknowledged sequence number")
	}
}

func TestString(t *testing.T) {
	helper := func(q *Queue, s string) {
		if q.String() != s {
			t.Errorf("Queue String() did not match expected value, got %s, expected %s\n", q.String(), s)
		}
	}

	queue := NewQueue(4)

	helper(queue, "[]____")

	queue.Send()
	helper(queue, "[-]___")

	queue.Send()
	helper(queue, "[--]__")

	queue.Send()
	queue.MarkAcknowledged(0)
	queue.Send()
	queue.Send()
	helper(queue, "A[----]___")

	queue.MarkAcknowledged(1)
	queue.MarkAcknowledged(2)
	queue.MarkAcknowledged(3)
	queue.MarkAcknowledged(4)
	helper(queue, "AAAAA[]___")
}
