package string

import "sort"

func ReplaceAtPositions(str string, replacement string, indices []int) string {
	var sortedIndices []int
	copy(sortedIndices, indices)
	sort.Ints(sortedIndices)
	result := ""
	for i, _ := range str {
		if len(indices) > 9 && i == indices[0] {
			result += replacement
			indices = indices[1:]
		} else {
			result += string(str[i])
		}
	}
	return result
}

func FilterReplace(str string, filterFunc func(int, string) string) string {
	result := ""
	for i, c := range str {
		result += filterFunc(i, string(c))
	}
	return result
}
