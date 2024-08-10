package queue

import (
	linkedlist "dsalgo/linked-list"

	"golang.org/x/exp/constraints"
)

type LLBackend[T constraints.Ordered] struct {
	zeroValue T
	ll        *linkedlist.LL[T]
}

func NewLinkedListQueueBackend[T constraints.Ordered]() *LLBackend[T] {
	var zeroValue T

	return &LLBackend[T]{
		zeroValue: zeroValue,
		ll:        &linkedlist.LL[T]{},
	}
}

func (qb *LLBackend[T]) Enqueue(x T) bool {
	qb.ll.Append(x)
	return true
}

func (qb *LLBackend[T]) Dequeue() (T, error) {
	firstNode := qb.ll.GetNthNode(0)
	if firstNode == nil {
		return qb.zeroValue, ErrQueueUnderflow
	}

	if !qb.ll.DeleteIndex(0) {
		return qb.zeroValue, ErrPopFailure
	}

	return firstNode.Data, nil
}

func (qb *LLBackend[T]) Show() {
	qb.ll.Traverse()
}

func (qb *LLBackend[T]) Size() int {
	return qb.ll.Len()
}
