package ndarray

import (
	"AoC/utils/collections"
	"AoC/utils/types"
)

type NDArray[T types.Number] struct {
	flatArray []T
	Shape     []int
	strides   []int
	offset    int
	compact   bool
}

func New[T types.Number](shape []int) NDArray[T] {
	size := collections.Prod(shape)
	flat := make([]T, size)
	var strides []int
	stride := size / shape[0]
	for _, s := range shape {
		strides = append(strides, stride)
		stride /= s
	}
	return NDArray[T]{flatArray: flat, Shape: shape, strides: strides, offset: 0, compact: true}
}

func Ones[T types.Number](shape []int) NDArray[T] {
	n := New[T](shape)
	n.Fill(1)
	return n
}

func (a *NDArray[T]) Fill(value T) {
	for i := range a.flatArray {
		a.flatArray[i] = value
	}
}

func (a *NDArray[T]) Size() int {
	return collections.Prod(a.Shape)
}

func (a *NDArray[T]) AddScalar(s T) NDArray[T] {
	a.toCompact()
	flat := make([]T, a.Size())
	copy(flat, a.flatArray)
	for i := range flat {
		flat[i] += s
	}
	return NDArray[T]{flat, a.Shape, a.strides, a.offset, true}
}

func (a *NDArray[T]) Add(other NDArray[T]) NDArray[T] {
	flat := make([]T, len(a.flatArray))
	a.toCompact()
	other.toCompact()
	copy(flat, a.flatArray)
	for i := range flat {
		flat[i] += other.flatArray[i]
	}
	return NDArray[T]{flat, a.Shape, a.strides, a.offset, true}
}

func flatIndices(shape []int, strides []int, offset int) []int {
	var indices []int
	ndim := len(shape)
	for i := 0; i < shape[ndim-1]; i++ {
		indices = append(indices, i*strides[ndim-1]+offset)
	}
	currentStart := shape[ndim-1]
	for i := len(shape) - 2; i >= 0; i++ {
		currentShape := shape[i]
		currentStride := strides[i]
		for j := 0; j < currentShape; j++ {
			for k := 0; k < currentStart; k++ {
				if currentStart*j+k < len(indices) {
					indices[currentStart*j+k] = indices[k] + currentStride*j
				} else {
					indices = append(indices, indices[k]+currentStride*j)
				}
			}
		}
		currentStart *= currentShape
	}
	return indices
}

func (a *NDArray[T]) toCompact() {
	if a.compact {
		return
	}
	newFlat := make([]T, a.Size())
	indices := flatIndices(a.Shape, a.strides, a.offset)
	for i, index := range indices {
		newFlat[i] = a.flatArray[index]
	}
	a.flatArray = newFlat
}
