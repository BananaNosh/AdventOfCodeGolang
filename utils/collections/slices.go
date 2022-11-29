package collections

import "AoC/utils/types"

func unpackIndirect[T any](slice []T, vars ...*T) {
	for i, elem := range slice {
		*vars[i] = elem
	}
}

func unpackTwo[T any](slice []T) (T, T) {
	return slice[0], slice[1]
}

func unpackThree[T any](slice []T) (T, T, T) {
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

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func Sum[T types.Number](slice []T) T {
	return reduce(slice, func(s T, acc T) T {
		return acc + s
	}, 0)
}

func Prod[T types.Number](slice []T) T {
	return reduce(slice, func(s T, acc T) T {
		return acc * s
	}, 1)
}
