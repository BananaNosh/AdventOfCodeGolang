package AoC16

import (
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/ndarray"
	"AoC/utils/requests"
	"fmt"
)

func posForLine(pos ndarray.NDArray[int], line string) ndarray.NDArray[int] {
	for _, d := range line {
		var step ndarray.NDArray[int]
		if d == 'U' {
			step = ndarray.New1D([]int{0, -1})
		}
		if d == 'D' {
			step = ndarray.New1D([]int{0, 1})
		}
		if d == 'L' {
			step = ndarray.New1D([]int{-1, 0})
		}
		if d == 'R' {
			step = ndarray.New1D([]int{1, 0})
		}
		pos = *(pos.Add(step).MaxScalar(0).MinScalar(2))
	}
	return pos
}

func AoC2() {
	year := 2016
	day := 2
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")
	lines := io.ReadInputLines(2, 2016)
	fmt.Println(lines)
	pos := ndarray.Ones[int]([]int{2})
	pad := ndarray.New(math.Range[int](1, 10), []int{3, 3})
	number := 0
	for _, line := range lines {
		pos = posForLine(pos, line)
		fmt.Println(pos)
		//pad[pos[0]]
		//pos.Get([][]int{math.Range(0, 2)})
		posSlice := pos.GetSlice(math.Range(0, 2)) // TODO use toSLice
		number *= 10
		number += pad.GetSlice([]int{posSlice[0]}, posSlice[1])[0]
	}
	fmt.Println("Part 1:")
	fmt.Println(number)
	requests.SubmitAnswer(day, year, number, 1)
	fmt.Println("Part 2:")
	// requests.SendAnswer(day, year, 0, 2)
}
