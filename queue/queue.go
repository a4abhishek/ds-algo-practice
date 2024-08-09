package queue

import "fmt"

type Backend interface {
	Enqueue(data int) bool
	Dequeue() (int, error)
	Show()
	Size() int
}

type Queue struct {
	backend Backend
}

func NewQueue(backend Backend) *Queue {
	return &Queue{
		backend: backend,
	}
}

func (s *Queue) Enqueue(x int) bool {
	return s.backend.Enqueue(x)
}

func (s *Queue) Dequeue() (int, error) {
	return s.backend.Dequeue()
}

func (s *Queue) Show() {
	s.backend.Show()
}

func (s *Queue) Size() int {
	return s.backend.Size()
}

func Driver() {
	// q := NewQueue(NewLinkedListQueueBackend())
	q := NewQueue(NewSliceBackend())

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
