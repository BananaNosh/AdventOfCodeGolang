package AoC22

import (
	"AoC/utils/date"
	"AoC/utils/io"
	"fmt"
	"os"
	"strconv"
)

func AoC12() {
	year := 2022
	day := 12
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInput(12, 2022)
	fmt.Println(input)
	fmt.Println("Part 1:")
	// requests.SubmitAnswer(day, year, 0, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
