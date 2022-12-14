package AoC_22_4

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
)

func checkIfContaining(elf1 []int, elf2 []int) bool {
	if elf1[0] <= elf2[0] && elf1[1] >= elf2[1] {
		return true
	}
	return elf2[0] <= elf1[0] && elf2[1] >= elf1[1]
}

func checkIfOverlapping(elf1 []int, elf2 []int) bool {
	if elf1[0] <= elf2[0] && elf1[1] >= elf2[0] {
		return true
	}
	return elf2[0] <= elf1[0] && elf2[1] >= elf1[0]
}

func AoC4() {
	year := 2022
	day := 4
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputFromRegexPerLineInt("(\\d+)-(\\d+),(\\d+)-(\\d+)", 4, 2022)
	fmt.Println(input)
	fmt.Println("Part 1:")
	total := reduceToTotalCount(input, checkIfContaining)
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 1)

	fmt.Println("Part 2:")
	total = reduceToTotalCount(input, checkIfOverlapping)
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 2)
}

func reduceToTotalCount(input [][]int, checkFunc func([]int, []int) bool) int {
	return collections.Reduce(input, func(totalContaining int, elves []int) int {
		if checkFunc(elves[:2], elves[2:]) {
			return totalContaining + 1
		}
		return totalContaining
	}, 0)
}
