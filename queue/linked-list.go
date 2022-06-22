package queue

import (
	linkedlist "dsalgo/linked-list"
	"errors"
)

type LLBackend struct {
	ll *linkedlist.LL[int]
}

func NewLinkedListQueueBackend() *LLBackend {
	return &LLBackend{
		ll: &linkedlist.LL[int]{},
	}
}

func (sb *LLBackend) Enqueue(x int) bool {
	sb.ll.Append(x)
	return true
}

func (sb *LLBackend) Dequeue() (int, error) {
	firstNode := sb.ll.GetNthNode(0)
	if firstNode == nil {
		return 0, errors.New("Queue underflow")
	}

	if !sb.ll.DeleteIndex(0) {
		return 0, errors.New("could not pop")
	}

	return firstNode.Data, nil
}

func (sb *LLBackend) Show() {
	sb.ll.Traverse()
}

func (sb *LLBackend) Size() int {
	return sb.ll.Len()
}
