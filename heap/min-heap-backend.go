package heap

import "golang.org/x/exp/constraints"

// minHeapBackend is a struct that implements a generic min-heap using a slice.
type minHeapBackend[T constraints.Ordered] struct {
	zeroVal T
	l       []T
}

// NewMinHeapBackend creates and returns a new instance of minHeapBackend.
func NewMinHeapBackend[T constraints.Ordered]() *minHeapBackend[T] {
	var zeroVal T

	return &minHeapBackend[T]{
		zeroVal: zeroVal,
		l:       []T{},
	}
}

// existsElement checks whether an element exists at the specified index.
func (h *minHeapBackend[T]) existsElement(i int) bool {
	return i >= 0 && i < len(h.l)
}

// lastIndex returns the index of the last element in the heap.
// Boundary check is out of scope for this method. Use `existsElement` to verify if the returned value falls within boundary.
func (h *minHeapBackend[T]) lastIndex() int {
	return len(h.l) - 1
}

// parent returns the index of the parent of the element at the specified index.
// Boundary check is out of scope for this method. Use `existsElement` to verify if the returned value falls within boundary.
func (h *minHeapBackend[T]) parent(i int) int {
	return (i - 1) / 2
}

// leftIndex returns the index of the left child of the element at the specified index.
// Boundary check is out of scope for this method. Use `existsElement` to verify if the returned value falls within boundary.
func (h *minHeapBackend[T]) leftIndex(i int) int {
	return 2*i + 1
}

// rightIndex returns the index of the right child of the element at the specified index.
// Boundary check is out of scope for this method. Use `existsElement` to verify if the returned value falls within boundary.
func (h *minHeapBackend[T]) rightIndex(i int) int {
	return 2 * (i + 1)
}

// swapElements swaps the elements at the two specified indices.
// Boundary check is out of scope for this method. Use `existsElement` to verify if the returned value falls within boundary.
func (h *minHeapBackend[T]) swapElements(i, j int) {
	h.l[i], h.l[j] = h.l[j], h.l[i]
}

// heapifyUp restores the heap property by moving the element at the specified index up the heap.
func (h *minHeapBackend[T]) heapifyUp(i int) {
	if i == 0 || !h.existsElement(i) {
		return
	}

	parentIndex := h.parent(i)
	if h.l[i] < h.l[parentIndex] {
		h.swapElements(i, parentIndex)
		h.heapifyUp(parentIndex)
	}
}

// heapifyDown restores the heap property by moving the element at the specified index down the heap.
func (h *minHeapBackend[T]) heapifyDown(i int) {
	if !h.existsElement(i) || i == h.lastIndex() {
		return
	}

	compareWith := -1

	leftIndex := h.leftIndex(i)
	rightIndex := h.rightIndex(i)
	if h.existsElement(leftIndex) && h.l[leftIndex] < h.l[i] {
		compareWith = leftIndex
	} else if h.existsElement(rightIndex) && h.l[rightIndex] < h.l[i] {
		compareWith = rightIndex
	}

	if !h.existsElement(compareWith) {
		return
	}

	h.swapElements(i, compareWith)
	h.heapifyDown(compareWith)
}

// Insert adds a new element to the heap and maintains the heap property.
func (h *minHeapBackend[T]) Insert(key T) error {
	h.l = append(h.l, key)
	h.heapifyUp(h.lastIndex())

	return nil
}

// Extract removes and returns the root element (smallest) from the heap.
// If the heap is empty, it returns an error.
func (h *minHeapBackend[T]) Extract() (T, error) {
	topVal, err := h.Peek()
	if err != nil {
		return h.zeroVal, err
	}

	h.l[0] = h.l[h.lastIndex()]
	h.l = h.l[:h.lastIndex()]

	h.heapifyDown(0)

	return topVal, nil
}

// Peek returns the root element (smallest) from the heap without removing it.
// If the heap is empty, it returns an error.
func (h *minHeapBackend[T]) Peek() (T, error) {
	if h.IsEmpty() {
		return h.zeroVal, ErrEmptyHeap
	}

	return h.l[0], nil
}

// Size returns the number of elements currently in the heap.
func (h *minHeapBackend[T]) Size() int {
	return len(h.l)
}

// IsEmpty checks if the heap is empty and returns true if it is, otherwise false.
func (h *minHeapBackend[T]) IsEmpty() bool {
	return h.Size() == 0
}
