package types

import "strconv"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToInt(str string) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return number
}

func ToIntOrElse(str string, elseVal int) int {
	number, err := strconv.Atoi(str)

	if err != nil {
		return elseVal
	}

	return number
}
