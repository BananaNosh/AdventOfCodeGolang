package AoC22

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"regexp"
	"strings"
)

func parseCrates(lines []string) []collections.Stack[string] {
	stackCount := len(strings.Split(strings.TrimSpace(lines[len(lines)-1]), "  "))
	stacks := make([]collections.Stack[string], stackCount)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for itemIndex := 0; itemIndex < stackCount; itemIndex++ {
			runeIndex := itemIndex*4 + 1
			if runeIndex >= len(line) {
				continue
			}
			item := string(line[runeIndex])
			if strings.TrimSpace(item) != "" {
				stacks[itemIndex].Push(item)
			}
		}
	}
	return stacks
}

func moveCrates(lines []string, stacks []collections.Stack[string], newMode bool) {
	regex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	for _, line := range lines {
		numbers := types.ToIntSlice(regex.FindStringSubmatch(line)[1:])
		fmt.Println(numbers)
		count := numbers[0]
		from := numbers[1] - 1
		to := numbers[2] - 1
		if newMode {
			moveMultiple9001(count, stacks, from, to)
		} else {
			moveMultiple9000(count, stacks, from, to)
		}
	}
}

func moveMultiple9000(count int, stacks []collections.Stack[string], from int, to int) {
	for i := 0; i < count; i++ {
		value := stacks[from].Pop()
		stacks[to].Push(value)
	}
}

func moveMultiple9001(count int, stacks []collections.Stack[string], from int, to int) {
	var temp collections.Stack[string]
	for i := 0; i < count; i++ {
		value := stacks[from].Pop()
		temp.Push(value)
	}
	for i := 0; i < count; i++ {
		stacks[to].Push(temp.Pop())
	}
}

func readTopCrates(stacks []collections.Stack[string]) []string {
	var top []string
	for _, stack := range stacks {
		if stack.Len() > 0 {
			top = append(top, stack.Pop())
		}
	}
	return top
}

func AoC5() {
	year := 2022
	day := 5
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInput(5, 2022)
	split := strings.Split(input, "\n\n")
	stackLines := strings.Split(strings.TrimRight(split[0], "\n "), "\n")
	moveLines := strings.Split(strings.TrimSpace(split[1]), "\n")
	fmt.Println(input)
	fmt.Println("Part 1:")
	stacks := parseCrates(stackLines)
	moveCrates(moveLines, stacks, false)
	topCrates := readTopCrates(stacks)
	fmt.Println(topCrates)
	requests.SubmitStringAnswer(day, year, strings.Join(topCrates, ""), 1)
	fmt.Println("Part 2:")
	stacks = parseCrates(stackLines)
	moveCrates(moveLines, stacks, true)
	topCrates = readTopCrates(stacks)
	fmt.Println(topCrates)
	requests.SubmitStringAnswer(day, year, strings.Join(topCrates, ""), 2)
}
