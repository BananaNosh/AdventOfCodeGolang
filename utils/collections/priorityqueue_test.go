package collections

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue_Add(t *testing.T) {
	// Some items and their priorities.

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue[string], 0)
	//heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item[string]{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	heap.Push(&pq, &Item[string]{value: "banana", priority: 3})
	heap.Push(&pq, &Item[string]{value: "apple", priority: 2})
	//pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	fmt.Println()

	queue := NewPriorityQueue[string]()
	queue.Add("banana", 3)
	queue.Add("apple", 5)
	queue.Add("orange", 1)

	for queue.Len() > 0 {
		item := heap.Pop(queue).(*Item[string])
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	fmt.Println()
}
