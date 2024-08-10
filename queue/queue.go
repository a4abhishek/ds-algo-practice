package queue

import (
	"fmt"
)

type Backend[T any] interface {
	Enqueue(data T) bool
	Dequeue() (T, error)
	Show()
	Size() int
}

type Queue[T any] struct {
	backend Backend[T]
}

func NewQueue[T any](backend Backend[T]) *Queue[T] {
	return &Queue[T]{
		backend: backend,
	}
}

func (s *Queue[T]) Enqueue(x T) bool {
	return s.backend.Enqueue(x)
}

func (s *Queue[T]) Dequeue() (T, error) {
	return s.backend.Dequeue()
}

func (s *Queue[T]) Show() {
	s.backend.Show()
}

func (s *Queue[T]) Size() int {
	return s.backend.Size()
}

func Driver() {
	// q := NewQueue(NewLinkedListQueueBackend())
	q := NewQueue[int](NewSliceBackend[int]())

	i := 0
	for {
		fmt.Println(`
1. Enqueue
2. Dequeue
3. Print
4. Size
5. Exit
`)

		inputReceived := false
		for !inputReceived {
			fmt.Printf("Enter your choice: ")
			fmt.Scanf("%d", &i)

			switch i {
			case 1:
				var data int
				fmt.Printf("Enter data: ")
				fmt.Scanf("%d", &data)
				fmt.Println("Push Status: ", q.Enqueue(data))
			case 2:
				data, err := q.Dequeue()
				if err != nil {
					fmt.Println("Pop Error: ", err)
				} else {
					fmt.Println("Poped Data: ", data)
				}
			case 3:
				q.Show()
			case 4:
				fmt.Println("Size: ", q.Size())
			case 5:
				return
			default:
				fmt.Println("Invalid choice.")
				continue
			}

			inputReceived = true
		}
		fmt.Println("Done.")
	}
}
