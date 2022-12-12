package search

import (
	"AoC/utils/collections"
)

type Neighbour[T any] struct {
	Value    T
	Distance int
}

type Item[T any] struct {
	value        T
	goneDistance int
	prev         *Item[T]
}

func AlphaStar[T comparable](start T, neighboursWithDistFunc func(T) [](Neighbour[T]), isGoal func(T) bool, heuristic func(T) int) []T {
	queue := collections.NewPriorityQueue[Item[T]]()
	startHeuristic := heuristic(start)
	queue.Add(Item[T]{value: start, goneDistance: 0}, -startHeuristic)
	seen := make(map[T]int) // the already seen items with their goneDistance
	var result []T
	for queue.Len() > 0 {
		currentItem, _ := queue.Retrieve()
		if isGoal(currentItem.value) {
			result = append(result, currentItem.value)
			for currentItem.prev != nil {
				currentItem = *currentItem.prev
				result = append(result, currentItem.value)
			}
			return collections.Reverse(result)
		}
		for _, neighbour := range neighboursWithDistFunc(currentItem.value) {
			neighbourPos := neighbour.Value
			if currentItem.prev != nil && neighbourPos == currentItem.prev.value {
				continue
			}
			neighbourDist := currentItem.goneDistance + neighbour.Distance
			neighbourPrio := neighbourDist + heuristic(neighbourPos)
			if collections.HasKey(seen, neighbourPos) && seen[neighbourPos] <= neighbourDist {
				continue
			}
			queue.Add(Item[T]{neighbourPos, neighbourDist, &currentItem}, -neighbourPrio)
			seen[neighbourPos] = neighbourDist
		}
	}
	return result
}
