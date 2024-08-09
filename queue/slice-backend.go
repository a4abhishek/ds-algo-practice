package queue

import (
	"fmt"
	"os"
)

type SliceBackend struct {
	l []int
}

func NewSliceBackend() *SliceBackend {
	return &SliceBackend{
		l: []int{},
	}
}

func (q *SliceBackend) Enqueue(data int) bool {
	if q == nil {
		fmt.Fprintln(os.Stderr, "Slice backend for queue is not instantiated")
		return false
	}

	if q.l == nil {
		q.l = []int{}
	}

	q.l = append(q.l, data)

	return true
}

func (q *SliceBackend) Dequeue() (int, error) {
	if q == nil {
		fmt.Fprintln(os.Stderr, "Slice backend for queue is not instantiated")
		return 0, ErrPopFailure
	}

	if len(q.l) == 0 {
		return 0, ErrQueueUnderflow
	}

	n := q.l[0]
	q.l = q.l[1:]

	return n, nil
}

func (q *SliceBackend) Show() {
	if q.Size() == 0 {
		fmt.Println("Queue is empty")
		return
	}

	for _, x := range q.l {
		fmt.Printf("%d\t", x)
	}
	fmt.Println()
}

func (q *SliceBackend) Size() int {
	return len(q.l)
}
