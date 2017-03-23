package main

import (
	"fmt"
	"math/rand"
)

func main() {
	h := &genericHeap{}
	for i := 6; i > 0; i-- {
		h.push(rand.Intn(100))
	}

	fmt.Printf("\nminimum: %d\n", (*h)[0])
	balanced := []int{}
	for h.Len() > 0 {
		balanced = append(balanced, h.Pop().(int))
		/*for _, v := range balanced {
			fmt.Printf("%d ", v)
		}
		fmt.Println()*/
	}
	for i := len(balanced)-1; i >= 0; i-- {
		fmt.Printf("%d ", balanced[i])
	}
	fmt.Println()

	fmt.Println(findMedian(balanced))

}

func findMedian(heap []int) int {
	if (len(heap) % 2 == 0) {
		median := (heap[len(heap) / 2] + heap[len(heap) / 2 - 1]) /2
		return median
	} else {
		median := heap[len(heap) / 2]
		return median
	}
}

type genericHeap []interface{}

func (h genericHeap) getLeftChildIndex(parentIndex int) int {
	return 2 * parentIndex + 1
}
func (h genericHeap) getRightChildIndex(parentIndex int) int {
	return 2 * parentIndex + 2
}
func (h genericHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1 ) / 2
}

func (h genericHeap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < h.Len()
}
func (h genericHeap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < h.Len()
}
func (h genericHeap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h *genericHeap) leftChild(index int) interface{} {
	return (*h)[h.getLeftChildIndex(index)].(int)
}

func (h *genericHeap) rightChild(index int) interface{} {
	return (*h)[h.getRightChildIndex(index)].(int)
}

func (h *genericHeap) parent(index int) interface{} {
	return (*h)[h.getParentIndex(index)].(int)
}

func (h genericHeap) Len() int {
	return len(h)
}

func (h genericHeap) Less(i, j int) bool {

	return h[i].(int) < h[j].(int)
}

func (h *genericHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *genericHeap) extraCapacity() {

}

func (h genericHeap) peek() interface{} {
	if (h.Len() != 0) {
		return h[0].(int)
	} else {
		return nil
	}
}

func (h *genericHeap) Pop() interface{} {
	if (h.Len() != 0) {
		old := *h
		n := old.Len()
		x := old[0]
		*h = old[1: n]
		h.heapifyDown()
		return x
	} else {
		return nil
	}
}

func (h *genericHeap) push(x interface{}) {
	*h = append(*h, x.(int))
	h.heapifyUp()
}

func (h *genericHeap) heapifyUp() {
	index := h.Len() - 1
	for (h.hasParent(index) && (h.parent(index).(int) > (*h)[index].(int))) {
		h.Swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h genericHeap) heapifyDown() {
	index := 0
	for (h.hasLeftChild(index)) {
		smallerChildIndex := h.getLeftChildIndex(index)
		if (h.hasRightChild(index) && h.rightChild(index).(int) < h.leftChild(index).(int)) {
			smallerChildIndex = h.getRightChildIndex(index)
		}
		if (h[index].(int) < h[smallerChildIndex].(int)) {
			break
		} else {
			h.Swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}
