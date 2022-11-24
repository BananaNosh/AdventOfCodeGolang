package collections

import "AoC/utils/types"

type Key interface {
	comparable
}

type SortableKey interface {
	Key
	Sortable
}

type Sortable interface {
	types.Number | string
}

type Value interface {
	comparable
}

type NumberValue interface {
	Value
	types.Number
}
