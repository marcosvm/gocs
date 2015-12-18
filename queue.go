package gocs

import (
	"fmt"
	"sync"
)

// Queue is a data structure that provide a FIFO behaviour
// operations are enqueue and dequeue both O(1) in this implementation
// items are stored in a slice of interface{}
type Queue struct {
	sync.Mutex
	s    []interface{}
	head int
	tail int
}

// InitQueue inits a Queue of fixed capacity cap
func InitQueue(cap int) *Queue {
	return &Queue{s: make([]interface{}, cap)}
}

// Enqueue stores the item at the tail of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.Lock()
	defer q.Unlock()

	q.s[q.tail] = item

	if q.tail == len(q.s)-1 {
		q.tail = 0
	} else {
		q.tail += 1
	}
}

// Dequeue removes the item from the head of the queue and returns it
func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()

	item := q.s[q.head]
	q.s[q.head] = nil

	if q.head == len(q.s)-1 {
		q.head = 0
	} else {
		q.head += 1
	}
	return item
}

func (q *Queue) Empty() bool {
	q.Lock()
	defer q.Unlock()
	return q.head == q.tail
}

// String is a string represtation of the Queue, for debugging purposes only.
func (q *Queue) String() string {
	return fmt.Sprintf("h: %d t: %d cap: %d %v", q.head, q.tail, len(q.s), q.s)
}
