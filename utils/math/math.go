package math

import "AoC/utils/types"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Gcd(first uint, second uint) uint {
	if first == second {
		return first
	}
	if first == 0 {
		return second
	}
	if second == 0 {
		return first
	}
	big := Max(first, second)
	small := Min(first, second)
	return Gcd(small, big%small)
}

func MaxInSlice[T types.Number](numbers []T) T {
	var max T
	max = numbers[0]
	for _, value := range numbers[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

func MinInSlice[T types.Number](numbers []T) T {
	var min T
	min = numbers[0]
	for _, value := range numbers[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

func Max[T types.Number](first, second T) T {
	if first > second {
		return first
	}
	return second
}

func Min[T types.Number](first, second T) T {
	return -(Max(-first, -second))
}

func MaxInt() int {
	return int(^uint(0) >> 1)
}
