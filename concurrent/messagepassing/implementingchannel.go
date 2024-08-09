package messagepassing

import (
	"cocurrent/concurrent/sharedmemory"
	"container/list"
	"sync"
)

// A shared queue data structure that acts like a buffer to store the messages
// betwenn sender and receiver
// Concurrent access protection for the shared data structure so that multiple
// senders and receivers do not interfere with each other
// Access control that blocks the execution of a receiver when the buffer is empty
// Access control that blocks the execution of a sender when the buffer is full
type Channel[M any] struct {
	// capacity semaphore to block the sender if the buffer is full
	capactiySema *sharedmemory.Semaphore

	// buffer size semaphore to block the receiver if the buffer is empty
	sizeSema *sharedmemory.Semaphore

	// linked list to be used as a queue data structure
	buffer *list.List

	// mutex protecting our shared list data structure
	mutex sync.Mutex
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capactiySema: sharedmemory.NewSemaphore(capacity),

		sizeSema: sharedmemory.NewSemaphore(0),

		buffer: list.New(),
	}
}

func (c *Channel[M]) Send(message M) {
	// block the goroutine if the buffer is full
	c.capactiySema.Acquire()

	// safeley add the message to the buffer
	c.mutex.Lock()
	c.buffer.PushBack(message)
	c.mutex.Unlock()

	// if any receiver goroutine are blocked, waiting for message,
	// resume one of them.
	c.sizeSema.Release()
}

func (c *Channel[M]) Receive() M {
	// unlock a sender waiting for capacity space.
	c.capactiySema.Release()

	// if the buffer is empty, block the receiver
	c.sizeSema.Acquire()

	// safely consume the next message from the buffer
	c.mutex.Lock()
	v := c.buffer.Remove(c.buffer.Front()).(M)
	c.mutex.Unlock()

	return v
}
