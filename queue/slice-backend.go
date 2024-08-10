package queue

import (
	"fmt"
	"os"
)

type SliceBackend[T any] struct {
	zeroValue T
	l         []T
}

func NewSliceBackend[T any]() *SliceBackend[T] {
	var zeroValue T

	return &SliceBackend[T]{
		zeroValue: zeroValue,
		l:         []T{},
	}
}

func (q *SliceBackend[T]) Enqueue(data T) bool {
	if q == nil {
		fmt.Fprintln(os.Stderr, "Slice backend for queue is not instantiated")
		return false
	}

	if q.l == nil {
		q.l = []T{}
	}

	q.l = append(q.l, data)

	return true
}

func (q *SliceBackend[T]) Dequeue() (T, error) {
	if q == nil {
		fmt.Fprintln(os.Stderr, "Slice backend for queue is not instantiated")

		var zeroValue T
		return zeroValue, ErrPopFailure
	}

	if len(q.l) == 0 {
		return q.zeroValue, ErrQueueUnderflow
	}

	n := q.l[0]
	q.l = q.l[1:]

	return n, nil
}

func (q *SliceBackend[T]) Show() {
	if q.Size() == 0 {
		fmt.Println("Queue is empty")
		return
	}

	for _, x := range q.l {
		fmt.Printf("%v\t", x)
	}
	fmt.Println()
}

func (q *SliceBackend[T]) Size() int {
	return len(q.l)
}
