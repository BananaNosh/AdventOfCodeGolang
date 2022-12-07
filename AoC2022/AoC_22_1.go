package AoC22

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"sort"
	"strings"
)

func AoC1() {
	year := 2022
	day := 1
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))
	elves := io.ReadAndSplitInput("\n\n", 2022, 1)
	fmt.Println("Part 1:")
	weightsPerElve := collections.Map(elves, func(elf string) []int {
		return types.ToIntSlice(strings.Split(elf, "\n"))
	})
	fmt.Println(weightsPerElve)
	totalPerElve := collections.Map(weightsPerElve, func(weights []int) int {
		return collections.Sum(weights)
	})
	max := math.MaxInSlice(totalPerElve)
	fmt.Println(max)
	//requests.SubmitAnswer(day, year, max, 1)
	fmt.Println("Part 2:")
	sort.Ints(totalPerElve)
	fmt.Println(totalPerElve[:3])
	//maxMultiple := totalPerElve[len(totalPerElve)-3:]
	maxMultiple := collections.Sum(totalPerElve[len(totalPerElve)-3:])
	fmt.Println(maxMultiple)
	requests.SubmitAnswer(day, year, maxMultiple, 2)
}
