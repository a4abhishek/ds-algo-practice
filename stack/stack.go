package stack

import "fmt"

type Backend interface {
	Push(data int) bool
	Pop() (int, error)
	Show()
	Size() int
}

type Stack struct {
	backend Backend
}

func NewStack(backend Backend) *Stack {
	return &Stack{
		backend: backend,
	}
}

func (s *Stack) Push(x int) bool {
	return s.backend.Push(x)
}

func (s *Stack) Pop() (int, error) {
	return s.backend.Pop()
}

func (s *Stack) Show() {
	s.backend.Show()
}

func (s *Stack) Size() int {
	return s.backend.Size()
}

func Driver() {
	s := NewStack(NewLinkedListStackBackend())

	i := 0
	for {
		fmt.Println(`
1. Push
2. Pop
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
				fmt.Println("Push Status: ", s.Push(data))
			case 2:
				data, err := s.Pop()
				if err != nil {
					fmt.Println("Pop Error: ", err)
				} else {
					fmt.Println("Poped Data: ", data)
				}
			case 3:
				s.Show()
			case 4:
				fmt.Println("Size: ", s.Size())
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
