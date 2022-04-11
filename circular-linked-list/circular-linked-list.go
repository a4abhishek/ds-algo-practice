package circularlinkedlist

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LL struct {
	head *Node
}

func (ll *LL) Len() int {
	if ll.head == nil {
		return 0
	}

	l := 1

	head := ll.head
	for head.Next != ll.head {
		l++
		head = head.Next
	}

	return l
}

// GetLastNode returns last node's pointer
// If ll is empty it will return nil
func (ll *LL) GetLastNode() *Node {
	if ll.head == nil {
		return nil
	}

	head := ll.head
	for head.Next != ll.head {
		head = head.Next
	}

	return head
}

// GetNthNode returns the node address of node at index
// It returns nil if linked list does not have as many nodes
func (ll *LL) GetNthNode(index int) *Node {
	if ll.head == nil {
		return nil
	}

	i := 0
	var node *Node

	head := ll.head
	for head.Next != ll.head && i != index {
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
func (ll *LL) Traverse() *Node {
	if ll.head == nil {
		return nil
	}

	fmt.Println(ll.head.Data)

	last := ll.head
	head := last.Next
	for head != ll.head {
		fmt.Println(head.Data)

		last = head
		head = head.Next
	}

	return last
}

func (ll *LL) Append(x int) {
	node := &Node{
		Data: x,
	}

	last := ll.GetLastNode()

	if last == nil {
		ll.head = node
	} else {
		last.Next = node
	}

	node.Next = ll.head // Need to assign it here, as ll.head might be nil in the above
}

func (ll *LL) Contains(x int) bool {
	if ll.head == nil {
		return false
	}

	head := ll.head
	for head.Next != ll.head {
		if head.Data == x {
			return true
		}

		head = head.Next
	}

	return false
}

func (ll *LL) Insert(index, data int) bool {
	if index < 0 {
		return false
	}

	node := &Node{
		Data: data,
	}

	if index == 0 {
		ll.head = node
		node.Next = ll.head // This should be below, as ll.head is still nil until previous line
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

func (ll *LL) DeleteData(x int) bool {
	head := ll.head
	var prev *Node

	if head == nil {
		return false
	}

	// If the data exists in the first node itself
	if head.Data == x {
		if head.Next == ll.head {
			ll.head = nil
		} else {
			ll.head = head.Next
		}

		return true
	}

	head = head.Next

	for head != ll.head && head.Data != x {
		prev = head
		head = head.Next
	}

	if head == ll.head {
		return false
	}

	prev.Next = head.Next
	return true
}

func (ll *LL) DeleteIndex(index int) bool {
	if ll.head == nil {
		return false
	}

	if index < 0 {
		return false
	}

	if index == 0 {
		if ll.head.Next == ll.head {
			ll.head = nil
		} else {
			ll.head = ll.head.Next
		}

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
	ll := LL{}

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
