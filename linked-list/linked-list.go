package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T comparable] interface {
	Len() int
	GetLastNode() *Node[T]
	GetNthNode(index int) *Node[T]
	Traverse() *Node[T]
	Append(x T)
	Contains(x T) bool
	Insert(index, data T) bool
	DeleteData(x T) bool
	DeleteIndex(index int) bool
}

type LL[T comparable] struct {
	head *Node[T]
}

func (ll *LL[T]) Len() int {
	l := 0

	head := ll.head
	for head != nil {
		l++
		head = head.Next
	}

	return l
}

// GetLastNode returns last node's pointer
// If ll is empty it will return nil
func (ll *LL[T]) GetLastNode() *Node[T] {
	var last *Node[T]

	head := ll.head
	for head != nil {
		last = head
		head = head.Next
	}

	return last
}

// GetNthNode returns the node address of node at index
// It returns nil if linked list does not have as many nodes
func (ll *LL[T]) GetNthNode(index int) *Node[T] {
	i := -1
	var node *Node[T]

	head := ll.head
	for head != nil && i != index {
		i++
		node = head

		head = head.Next
	}

	if i == index {
		return node
	}
	return nil
}

// Traverse prints all the values in linked list and returns last node's pointer
// If ll is empty it will return nil
func (ll *LL[T]) Traverse() *Node[T] {
	var last *Node[T]

	head := ll.head
	for head != nil {
		fmt.Println(head.Data)

		last = head
		head = head.Next
	}

	return last
}

func (ll *LL[T]) Append(x T) {
	node := &Node[T]{
		Data: x,
	}

	last := ll.GetLastNode()

	if last == nil {
		ll.head = node
	} else {
		last.Next = node
	}
}

func (ll *LL[T]) Contains(x T) bool {
	head := ll.head
	for head != nil {
		if head.Data == x {
			return true
		}

		head = head.Next
	}

	return false
}

func (ll *LL[T]) Insert(index int, data T) bool {
	if index < 0 {
		return false
	}

	node := &Node[T]{
		Data: data,
	}

	if index == 0 {
		node.Next = ll.head
		ll.head = node
		return true
	}

	prevNode := ll.GetNthNode(index - 1)

	if prevNode == nil {
		return false
	}

	node.Next = prevNode.Next
	prevNode.Next = node

	return true
}

func (ll *LL[T]) DeleteData(x T) bool {
	head := ll.head
	var prev *Node[T]

	for head != nil && head.Data != x {
		prev = head
		head = head.Next
	}

	if head == nil {
		return false
	}

	if prev == nil {
		ll.head = head.Next
	} else {
		prev.Next = head.Next
	}

	return true
}

func (ll *LL[T]) DeleteIndex(index int) bool {
	if ll.head == nil {
		return false
	}

	if index < 0 {
		return false
	}

	if index == 0 {
		ll.head = ll.head.Next
		return true
	}

	prevNode := ll.GetNthNode(index - 1)
	if prevNode == nil {
		return false
	}

	prevNode.Next = prevNode.Next.Next
	return true
}

func Driver() {
	var ll LinkedList[int] = new(LL[int])

	i := 0
	for {
		fmt.Println(`
1. Append
2. Insert at given index
3. Print data
4. Search data
5. Delete given data
6. Delete at the given index
7. Print length
8. Exit
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
				ll.Append(data)
			case 2:
				var index, data int
				fmt.Printf("Enter index: ")
				fmt.Scanf("%d", &index)
				fmt.Printf("Enter data: ")
				fmt.Scanf("%d", &data)
				fmt.Println("Success: ", ll.Insert(index, data))
			case 3:
				ll.Traverse()
			case 4:
				var data int
				fmt.Printf("Enter data: ")
				fmt.Scanf("%d", &data)
				fmt.Println("Finding status: ", ll.Contains(data))
			case 5:
				var data int
				fmt.Printf("Enter data: ")
				fmt.Scanf("%d", &data)
				fmt.Println("Deletion status: ", ll.DeleteData(data))
			case 6:
				var index int
				fmt.Printf("Enter index: ")
				fmt.Scanf("%d", &index)
				fmt.Println("Deletion status: ", ll.DeleteIndex(index))
			case 7:
				fmt.Println("Length: ", ll.Len())
			case 8:
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
