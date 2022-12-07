package AoC22

import (
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
)

func findPackageStart(buffer string, distinctCount int) int {
	lastSeen := make(map[rune]int, 4)
	for i, c := range buffer {
		if i >= distinctCount {
			removed := rune(buffer[i-distinctCount])
			lastSeen[removed] -= 1
			if lastSeen[removed] == 0 {
				delete(lastSeen, removed)
			}
		}
		lastSeen[c] += 1
		if len(lastSeen) == distinctCount {
			return i + 1
		}
	}
	return -1
}

func AoC6() {
	year := 2022
	day := 6
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInput(6, 2022)
	fmt.Println("Part 1:")
	start := findPackageStart(input, 4)
	fmt.Println(input, start)
	requests.SubmitAnswer(day, year, start, 1)
	fmt.Println("Part 2:")
	start = findPackageStart(input, 14)
	fmt.Println(start)

	requests.SubmitAnswer(day, year, start, 2)
}
