package collections

import (
	"AoC/utils/types"
)

func UnpackIndirect[T any](slice []T, vars ...*T) {
	for i, elem := range slice {
		*vars[i] = elem
	}
}

func UnpackTwo[T any](slice []T) (T, T) {
	return slice[0], slice[1]
}

func UnpackThree[T any](slice []T) (T, T, T) {
	return slice[0], slice[1], slice[2]
}

func Filter[T any](slice []T, filterFunc func(T) bool) []T {
	retSlice := make([]T, 0)

	for _, element := range slice {
		if filterFunc(element) {
			retSlice = append(retSlice, element)
		}
	}

	return retSlice
}

func Contains[T comparable](slice []T, word T) bool {
	for _, element := range slice {
		if element == word {
			return true
		}
	}
	return false
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func ReduceWithIndex[T, M any](s []T, f func(int, M, T) M, initValue M) M {
	acc := initValue
	for i, v := range s {
		acc = f(i, acc, v)
	}
	return acc
}

func Map[T, M any](s []T, f func(T) M) []M {
	var acc []M
	for _, v := range s {
		acc = append(acc, f(v))
	}
	return acc
}

func MapWithIndex[T, M any](s []T, f func(int, T) M) []M {
	var acc []M
	for i, v := range s {
		acc = append(acc, f(i, v))
	}
	return acc
}

func Sum[T types.Number](slice []T) T {
	return Reduce(slice, func(acc T, s T) T {
		return acc + s
	}, 0)
}

func Prod[T types.Number](slice []T) T {
	return Reduce(slice, func(acc T, s T) T {
		return acc * s
	}, 1)
}
