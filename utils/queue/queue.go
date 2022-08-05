package queue

import (
	"errors"
	"fmt"
)

const lazyGrowFactor = 4
const lazyShrinkFactor = 2

type Queue[T any] struct {
	values []T
	front  int
	rear   int
	length int
}

func Create[T any](capacity uint) *Queue[T] {
	q := new(Queue[T])

	// always have a capacity of at least 1
	if capacity == 0 {
		capacity = 1
	}

	q.values = make([]T, capacity)
	q.front = 0
	q.rear = 0
	q.length = 0
	return q
}

func New[T any]() *Queue[T] {
	return Create[T](1)
}

// ----
// GETTERS
// ----
func (q *Queue[T]) Length() int {
	return q.length
}

func (q *Queue[T]) Capacity() int {
	return len(q.values)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) IsFull() bool {
	return q.length == q.Capacity()
}

/* func (q *Queue[T]) IsSparse() bool {
	// NOTE: revise this
	return q.length < q.Capacity()
} */

// ----
// INTROSPECTION
// ----

func (q *Queue[T]) Peek() T {
	return q.values[q.front]
}

func (q *Queue[T]) ToString() string {
	return fmt.Sprintf("\n%+v\n", q)
}

// ----
// MUTATIONS
// ----

func (q *Queue[T]) Resize(size int) error {
	if q.length > size {
		return errors.New(fmt.Sprintf("RESIZE OVERFLOW: cannot resize queue with %d items to a capacity of %d", q.length, size))
	}

	newValues := make([]T, size)

	if q.front < q.rear {
		copy(newValues, q.values[q.front:q.rear])
	} else {
		// repair wrapping
		// copy over front to end of slice
		countFrontToEnd := copy(newValues, q.values[q.front:])
		// copy over start of slice to start of rear
		copy(newValues[countFrontToEnd:], q.values[:q.rear])
	}

	q.values = newValues
	q.front = 0
	q.rear = q.length

	return nil
}

// increment pointer, wrapping if necessary
func (q *Queue[T]) incPointer(pointer int) int {
	if pointer < q.Capacity()-1 {
		return pointer + 1
	}

	return 0
}

// ----
// METHODS
// ----

func (q *Queue[T]) Enqueue(value T) error {
	if q.IsFull() {
		return errors.New("OVERFLOW: queue is full")
	}

	q.rear = q.incPointer(q.rear)

	q.values[q.rear] = value
	q.length += 1

	return nil
}

func (q *Queue[T]) Dequeue() (bool, T) {
	if q.IsEmpty() {
		var nothing T
		return false, nothing
	}

	value := q.values[q.front]

	q.front = q.incPointer(q.front)
	q.length -= 1

	return true, value
}

// ----
// LAZY METHODS
//
// Automatically resizes queue as needed
// ----

func (q *Queue[T]) LazyEnqueue(value T) {
	// If full, 4x capacity
	if q.IsFull() {
		q.Resize(q.Capacity() * 4)
	}

	// Should never error due to clause above.
	// Reusing logic does have penalty of checking fullness twice.
	q.Enqueue(value)
}

func (q *Queue[T]) LazyDequeue() (bool, T) {
	exists, item := q.Dequeue()

	// If there's something in the queue, and the queue is under 25% full, half the capacity.
	if q.length > 1 && q.length < q.Capacity()/4 {
		q.Resize(q.Capacity() / 2)
	}

	return exists, item
}
