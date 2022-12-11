package collections

type Queue[T any] struct {
	elements []T
}

// NewQueue creates and returns a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0),
	}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) EnqueueMultiple(elements ...T) {
	for _, element := range elements {
		q.elements = append(q.elements, element)
	}
}

func (q *Queue[T]) EnqueueQueue(other *Queue[T]) {
	currentItemsInOther := other.Size()
	for i := 0; i < currentItemsInOther; i++ {
		otherItem := other.Dequeue()
		q.Enqueue(otherItem)
		other.Enqueue(otherItem)
	}
}

func (q *Queue[T]) DequeAllIndex(f func(int, T)) {
	currentItemsInQueue := q.Size()
	for i := 0; i < currentItemsInQueue; i++ {
		f(i, q.Dequeue())
	}
}

func (q *Queue[T]) DequeAll(f func(T)) {
	currentItemsInQueue := q.Size()
	for i := 0; i < currentItemsInQueue; i++ {
		f(q.Dequeue())
	}
}

// Dequeue removes and returns the element at the front of the queue.
// If the queue is empty, it throws an error
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("Queue Empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}
func (q *Queue[T]) DequeueOrDefault(def T) T {
	if q.IsEmpty() {
		return def
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

// Peek returns the element at the front of the queue without removing it.
// If the queue is empty, it returns nil.
func (q *Queue[T]) Peek() *T {
	if q.IsEmpty() {
		return nil
	}
	return &q.elements[0]
}

// IsEmpty returns true if the queue is empty, and false otherwise.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.elements)
}
