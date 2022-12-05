package collections

type item[T any] struct {
	value T //value as interface type to hold any data type
	next  *item[T]
}

type Stack[T any] struct {
	top  *item[T]
	size int
}

func (stack *Stack[T]) Len() int {
	return stack.size
}

func (stack *Stack[T]) Push(value T) {
	stack.top = &item[T]{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack[T]) Pop() (value T) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}
	panic("Stack empty")
}

func (stack *Stack[T]) PopOrDefault(def T) (value T) {
	defer func() {
		if a := recover(); a != nil {
			value = def
		}
	}()
	return stack.Pop()
}
