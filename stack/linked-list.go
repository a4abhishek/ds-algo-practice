package stack

import (
	linkedlist "dsalgo/linked-list"
	"errors"
)

type LLBackend struct {
	ll *linkedlist.LL
}

func NewLinkedListStackBackend() *LLBackend {
	return &LLBackend{
		ll: &linkedlist.LL{},
	}
}

func (sb *LLBackend) Push(x int) bool {
	return sb.ll.Insert(0, x)
}

func (sb *LLBackend) Pop() (int, error) {
	firstNode := sb.ll.GetNthNode(0)
	if firstNode == nil {
		return 0, errors.New("stack underflow")
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
