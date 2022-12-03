package collections

import (
	"fmt"
	"go/types"
)

type Set[T Key] struct {
	elements map[T]types.Nil
}

func NewSet[T Key]() Set[T] {
	set := new(Set[T])
	set.elements = make(map[T]types.Nil)
	return *set
}

func (set *Set[T]) Add(element T) {
	set.elements[element] = types.Nil{}
}

func (set *Set[T]) Remove(element T) {
	if _, isInMap := set.elements[element]; isInMap {
		delete(set.elements, element)
	}
}

func (set *Set[T]) Wipe() {
	set.elements = make(map[T]types.Nil)
}

func (set *Set[T]) Copy() Set[T] {
	resultSet := NewSet[T]()

	for element := range set.elements {
		resultSet.Add(element)
	}

	return resultSet
}

func (set *Set[T]) Has(element T) bool {
	_, has := set.elements[element]
	return has
}

func (set *Set[T]) Size() int {
	return len(set.elements)
}

func (set *Set[T]) Difference(setToIntersectWith Set[T]) Set[T] {
	resultSet := NewSet[T]()
	for element := range set.elements {
		if !setToIntersectWith.Has(element) {
			resultSet.Add(element)
		}
	}

	return resultSet
}

func (set *Set[T]) Intersect(setToIntersectWith Set[T]) Set[T] {
	resultSet := NewSet[T]()
	for element := range set.elements {
		if setToIntersectWith.Has(element) {
			resultSet.Add(element)
		}
	}
	return resultSet
}

func (set *Set[T]) GetRandom() T {
	for element := range set.elements {
		return element
	}
	panic("Set is Empty")
}

func (set *Set[T]) GetRandomOrDefault(def T) T {
	for element := range set.elements {
		return element
	}
	return def
}

func (set Set[T]) String() string {
	str := "{ "
	for element := range set.elements {
		str += fmt.Sprintf("'%v', ", element)
	}

	return str[:len(str)-2] + " }"
}

type CheckFunc[T Key] func(key T) bool

func CheckAll[T Key](set Set[T], check CheckFunc[T]) bool {
	for element := range set.elements {
		if check(element) == false {
			return false
		}
	}
	return true
}

func compare[T Sortable](l, r T) int8 {
	if l == r {
		return 0
	}
	if l > r {
		return 1
	}
	return -1
}

type OrderingSet[T SortableKey] struct {
	Set[T]
}

func NewComparingSet[T SortableKey]() OrderingSet[T] {
	var set OrderingSet[T]
	s := NewSet[T]()
	set.Set = s
	return set
}

func (set *OrderingSet[T]) Max() T {
	var max T
	for element := range set.elements {
		max = element
		break
	}
	for element := range set.elements {
		if compare(element, max) > 0 {
			max = element
		}
	}

	return max
}

func (set *OrderingSet[T]) Min() T {
	var min T
	for element := range set.elements {
		min = element
		break
	}
	for element := range set.elements {
		if compare(element, min) < 0 {
			min = element
		}
	}

	return min
}

//
//func (set *OrderingSet[T]) sorted() []T {
//	var result []T
//	for element := range set.elements {
//		result = append(result, element)
//	}
//
//}
