package AoC_22_25

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
	"math"
	"strings"
)

type SnafuDigit int

var SnafuStrings = [5]string{"=", "-", "0", "1", "2"}

func (snafu SnafuDigit) toString() string {
	return SnafuStrings[int(snafu)+2]
}

func snafuDigitFromString(digit string) SnafuDigit {
	for i, snafuString := range SnafuStrings {
		if digit == snafuString {
			return SnafuDigit(i - 2)
		}
	}
	panic("No such snafu")
}

func toSnafu(number int) string {
	snafu := make([]string, 0)
	borrow := false
	for number > 0 || borrow {
		mod := number % 5
		number /= 5
		if borrow {
			mod += 1
			borrow = false
		}
		if mod > 2 {
			mod -= 5
			borrow = true
		}
		digit := SnafuDigit(mod)
		snafu = append(snafu, digit.toString())
	}
	snafu = collections.Reverse(snafu)
	return strings.Join(snafu, "")
}

func fromSnafu(snafu string) int {
	number := 0
	for i := range snafu {
		currentPosition := len(snafu) - 1 - i
		posValue := int(math.Pow(5, float64(i)))
		sign := snafu[currentPosition]
		number += int(snafuDigitFromString(string(sign))) * posValue
	}
	return number
}

func AoC25() {
	year := 2022
	day := 25
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	lines := io.ReadInputLines(25, 2022)
	fmt.Println(lines)
	fmt.Println("Part 1:")
	numbers := collections.Map(lines, func(line string) int {
		return fromSnafu(line)
	})
	result := collections.Sum(numbers)
	resultSnafu := toSnafu(result)
	fmt.Println(numbers, result, resultSnafu)
	requests.SubmitStringAnswer(day, year, resultSnafu, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
