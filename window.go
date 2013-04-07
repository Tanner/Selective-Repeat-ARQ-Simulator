package arq

type Queue struct {
  contents []QueuePacket

  window Window
}

type Window struct {
  window []QueuePacket

  location int
}

type QueuePacket struct {
  packet Packet
  acknowledged bool
}

// NewQueue returns a Queue with the given window size at position zero
func NewQueue(windowSize int) *Queue {
  queue := new(Queue)
  window := new(Window)

  queue.contents = make([]QueuePacket, 10)

  window.window = queue.contents[:windowSize]
  window.location = 0

  queue.window = *window

  return queue
}

// AddPacket adds the given packet to the queue and marks it as unacknowledged
func (q *Queue) AddPacket(p Packet) {
  q.contents = append(q.contents, QueuePacket{p, false})
}