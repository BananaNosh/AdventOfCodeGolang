package search

import (
	"AoC/utils/collections"
	"fmt"
)

type Neighbour[T any] struct {
	Value    T
	Distance int
}

type Item[T any] struct {
	value        T
	goneDistance int
	prev         *Item[T]
	depth        int
}

func AStar[T comparable](start T, neighboursWithDistFunc func(T) [](Neighbour[T]), isGoal func(T) bool, heuristic func(T) int) []T {
	return aStarWithMaxDistanceOrGoal(start, neighboursWithDistFunc, isGoal, heuristic, -1)
}

func AStarWithMaxDistance[T comparable](start T, neighboursWithDistFunc func(T) [](Neighbour[T]), heuristic func(T) int, maxDistance int) []T { // TODO refactor
	return aStarWithMaxDistanceOrGoal(start, neighboursWithDistFunc, func(T) bool {
		return false
	}, heuristic, maxDistance)
}

func aStarWithMaxDistanceOrGoal[T comparable](start T, neighboursWithDistFunc func(T) [](Neighbour[T]), isGoal func(T) bool, heuristic func(T) int, maxDistance int) []T {
	queue := collections.NewPriorityQueue[Item[T]]()
	startHeuristic := heuristic(start)
	queue.Add(Item[T]{value: start, goneDistance: 0}, -startHeuristic)
	seen := make(map[T]int) // the already seen items with their goneDistance
	for queue.Len() > 0 {
		currentItem, _ := queue.Retrieve()
		if isGoal(currentItem.value) {
			return createPath(currentItem)
		}
		fmt.Println("dist", currentItem.goneDistance)
		if maxDistance > 0 && currentItem.goneDistance == maxDistance {
			return createPath(currentItem)
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
			if maxDistance < 0 || maxDistance >= neighbourDist {
				queue.Add(Item[T]{neighbourPos, neighbourDist, &currentItem, currentItem.depth + 1}, -neighbourPrio)
			}
			seen[neighbourPos] = neighbourDist
		}
		if maxDistance > 0 && queue.Len() == 0 {
			return createPath(currentItem)
		}
	}
	return make([]T, 0)
}

func createPath[T any](currentItem Item[T]) []T {
	var result []T
	result = append(result, currentItem.value)
	for currentItem.prev != nil {
		currentItem = *currentItem.prev
		result = append(result, currentItem.value)
	}
	return collections.Reverse(result)
}

func BreadthFirst[T comparable](start T, neighboursFunc func(T) []T, isGoal func(T) bool, maxDepth int) []T {
	queue := collections.NewQueue[Item[T]]()
	queue.Enqueue(Item[T]{value: start})
	for !queue.IsEmpty() {
		currentItem := queue.Dequeue()
		if isGoal(currentItem.value) {
			return createPath(currentItem)
		}
		if currentItem.goneDistance == maxDepth {
			return createPath(currentItem)
		}
		for _, neighbour := range neighboursFunc(currentItem.value) {
			if currentItem.prev != nil && neighbour == currentItem.prev.value {
				continue
			}
			queue.Enqueue(Item[T]{neighbour, currentItem.goneDistance + 1, &currentItem, currentItem.depth + 1})
		}
	}
	return make([]T, 0)
}

func FindBestWithIdentify[T any, C comparable](start T, neighboursFunc func(T) []T, valueFunc func(T) int, maxDepth int, identify func(T) C) []T {
	queue := collections.NewQueue[Item[T]]()
	queue.Enqueue(Item[T]{value: start})
	seen := make(map[C]int) // the already seen items with their goneDistance
	var best Item[T]
	var bestValue *int
	for !queue.IsEmpty() {
		currentItem := queue.Dequeue()
		//fmt.Println(identify(currentItem.value), currentItem.value)
		currentValueValue := valueFunc(currentItem.value)
		if bestValue == nil || currentValueValue > *bestValue {
			best = currentItem
			bestValue = &currentValueValue
		}
		fmt.Println(currentItem.depth)
		if currentItem.depth < maxDepth {
			for _, neighbour := range neighboursFunc(currentItem.value) {
				if currentItem.prev != nil && identify(neighbour) == identify(currentItem.prev.value) {
					continue
				}
				neighbourDist := currentItem.goneDistance + 1
				if collections.HasKey(seen, identify(neighbour)) && seen[identify(neighbour)] >= valueFunc(neighbour) {
					continue
				}
				seen[identify(neighbour)] = valueFunc(neighbour)
				queue.Enqueue(Item[T]{neighbour, neighbourDist, &currentItem, currentItem.depth + 1})
			}
		}
	}
	return createPath(best)

}

func FindBest[T comparable](start T, neighboursFunc func(T) []T, valueFunc func(T) int, maxDepth int) []T {
	return FindBestWithIdentify(start, neighboursFunc, valueFunc, maxDepth, func(t T) T {
		return t
	})
}
