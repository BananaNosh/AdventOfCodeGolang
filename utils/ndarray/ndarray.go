package ndarray

import (
	"AoC/utils/collections"
	"AoC/utils/math"
	"AoC/utils/types"
	"fmt"
)

type NDArray[T types.Number] struct {
	flatArray []T
	Shape     []int
	strides   []int
	offset    int
	compact   bool
}

func New1D[T types.Number](slice []T) *NDArray[T] {
	return New[T](slice, []int{len(slice)})
}

func New2D[T types.Number](slice [][]T) *NDArray[T] {
	shape := []int{len(slice), len(slice[0])}
	flatSlice := make([]T, shape[0]*shape[1])
	for i, row := range slice {
		for i2, value := range row {
			flatSlice[i*shape[1]+i2] = value
		}
	}
	return New[T](flatSlice, shape)
}

func Empty[T types.Number](shape []int) *NDArray[T] {
	size := collections.Prod(shape)
	flat := make([]T, size)
	var strides []int
	stride := size / shape[0]
	for _, s := range shape {
		strides = append(strides, stride)
		stride /= s
	}
	return &NDArray[T]{flatArray: flat, Shape: shape, strides: strides, offset: 0, compact: true}
}

func New[T types.Number](flatSlice []T, shape []int) *NDArray[T] {
	n := Empty[T](shape)
	if n.Size() > len(flatSlice) {
		panic("Not matchable")
	}
	for i, value := range flatSlice {
		n.flatArray[i] = value
	}
	return n
}

func Ones[T types.Number](shape []int) NDArray[T] {
	n := Empty[T](shape)
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

func (a *NDArray[T]) AddScalar(s T) *NDArray[T] {
	return scalarFunc(func(old T, value T) T {
		return old + value
	}, s, a)
}

func (a *NDArray[T]) Add(other NDArray[T]) *NDArray[T] {
	return binFunc(func(old T, value T) T {
		return old + value
	}, other, a)
}

func (a *NDArray[T]) MulScalar(s T) *NDArray[T] {
	return scalarFunc(func(old T, value T) T {
		return old * value
	}, s, a)
}

func (a *NDArray[T]) Mul(other NDArray[T]) *NDArray[T] {
	return binFunc(func(old T, value T) T {
		return old * value
	}, other, a)
}

func (a *NDArray[T]) MaxScalar(value T) *NDArray[T] {
	return scalarFunc(math.Max[T], value, a)
}

func (a *NDArray[T]) Max(other NDArray[T]) *NDArray[T] {
	return binFunc(math.Max[T], other, a)
}

func (a *NDArray[T]) MinScalar(value T) *NDArray[T] {
	return scalarFunc(math.Min[T], value, a)
}

func (a *NDArray[T]) Min(other NDArray[T]) *NDArray[T] {
	return binFunc(math.Min[T], other, a)
}
func (a *NDArray[T]) GetSlice(lastDim []int, otherDims ...int) []T {
	offset := 0
	if len(otherDims) != len(a.Shape)-1 {
		panic("Wrong number of dimensions")
	}
	for dim, index := range otherDims {
		if index >= a.Shape[dim] {
			panic(fmt.Sprintf("Index out of range for dim %d", dim))
		}
		offset += index * a.strides[dim]
	}
	out := make([]T, len(lastDim))
	for i, index := range lastDim {
		out[i] = a.flatArray[offset+index*a.strides[len(a.strides)-1]]
	}
	return out
}

//func (a *NDArray[T]) GetFlat(indices [][]int) []T {
//	for i, dimIndices := range indices {
//
//	}
//}

//func (a *NDArray[T]) ToSlice() []T {
//	indices := flatIndices(a.Shape, a.strides, a.offset)
//
//}

//func (a *NDArray[T]) Get(indices [][]int) *NDArray[T] {
//	total_size := collections.Reduce(indices, func(acc int, dim []int) int {
//		return acc + len(dim)
//	}, 0)
//	shape = make([]int, len(a.Shape))
//	stri = make([]int, len(a.Shape))
//	for i, dimIndices := range indices {
//		shape[i] = len(dimIndices)
//	}
//	var flatArray []T
//	NDArray[T]{flatArray, shape}
//	return a
//}

func binFunc[T types.Number](operation func(T, T) T, other NDArray[T], a *NDArray[T]) *NDArray[T] {
	flat := make([]T, len(a.flatArray))
	a.toCompact()
	other.toCompact()
	copy(flat, a.flatArray)
	for i, old := range flat {
		flat[i] = operation(old, other.flatArray[i])
	}
	return &NDArray[T]{flat, a.Shape, a.strides, a.offset, true}
}

func scalarFunc[T types.Number](operation func(T, T) T, value T, a *NDArray[T]) *NDArray[T] {
	a.toCompact()
	flat := make([]T, a.Size())
	copy(flat, a.flatArray)
	for i, old := range flat {
		flat[i] = operation(old, value)
	}
	return &NDArray[T]{flat, a.Shape, a.strides, a.offset, true}
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
