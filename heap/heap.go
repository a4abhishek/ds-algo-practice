package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Backend[T constraints.Ordered] interface {
	// Insert adds an element to the heap
	Insert(key T) error
	// Extract removes and returns the root element from the heap
	Extract() (T, error)
	// Peek returns the root element without removing it
	Peek() (T, error)
	// Size returns the number of elements in the heap
	Size() int
	// IsEmpty returns true if the heap is empty
	IsEmpty() bool
}

type Heap[T constraints.Ordered] struct {
	backend Backend[T]
}

func NewHeap[T constraints.Ordered](backend Backend[T]) *Heap[T] {
	return &Heap[T]{
		backend: backend,
	}
}

func (h *Heap[T]) Insert(key T) error {
	return h.backend.Insert(key)
}

func (h *Heap[T]) Extract() (T, error) {
	return h.backend.Extract()
}

func (h *Heap[T]) Peek() (T, error) {
	return h.backend.Peek()
}

func (h *Heap[T]) Size() int {
	return h.backend.Size()
}

func (h *Heap[T]) IsEmpty() bool {
	return h.backend.IsEmpty()
}

func Driver() {
	// Assuming you have an implementation like NewMinHeapBackend() or NewMaxHeapBackend()
	// You can uncomment the desired one
	heap := NewHeap[int](NewMinHeapBackend[int]())
	// heap := NewHeap[int](NewMaxHeapBackend[int]())

	var choice int
	for {
		fmt.Println(`
1. Insert
2. Extract
3. Peek
4. Size
5. Is Empty
6. Exit
`)

		fmt.Printf("Enter your choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			var data int
			fmt.Printf("Enter data: ")
			fmt.Scanf("%d", &data)
			if err := heap.Insert(data); err != nil {
				fmt.Println("Insert Error:", err)
			} else {
				fmt.Println("Inserted:", data)
			}
		case 2:
			if heap.IsEmpty() {
				fmt.Println("Heap is empty.")
			} else {
				extracted, err := heap.Extract()
				if err != nil {
					fmt.Println("Extract Error:", err)
				} else {
					fmt.Println("Extracted:", extracted)
				}
			}
		case 3:
			if heap.IsEmpty() {
				fmt.Println("Heap is empty.")
			} else {
				peeked, err := heap.Peek()
				if err != nil {
					fmt.Println("Peek Error:", err)
				} else {
					fmt.Println("Peeked:", peeked)
				}
			}
		case 4:
			fmt.Println("Size:", heap.Size())
		case 5:
			if heap.IsEmpty() {
				fmt.Println("Heap is empty.")
			} else {
				fmt.Println("Heap is not empty.")
			}
		case 6:
			return
		default:
			fmt.Println("Invalid choice.")
		}
		fmt.Println("Done.")
	}
}
