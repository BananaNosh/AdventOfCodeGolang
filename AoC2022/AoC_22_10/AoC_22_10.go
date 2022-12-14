package AoC_22_10

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"fmt"
	"strconv"
)

type Command struct {
	com string
	arg int
}

const (
	Noop = "noop"
	Addx = "addx"
)

func findSignalStrengths(commands []Command) []int {
	durationSoFar := computeCumulativeDuration(commands)
	totalDuration := collections.Last(durationSoFar)
	lastTimeStepIndex := -1
	currentValue := 1
	var strengthValues []int
	for currentTimeStep := 20; currentTimeStep < totalDuration+1; currentTimeStep += 40 {
		timeStepIndex := findTimeStepIndex(currentTimeStep-1, durationSoFar)
		relevantAddArgs := collections.Map(commands[lastTimeStepIndex+1:timeStepIndex+1], func(com Command) int {
			return com.arg
		})
		fmt.Println(relevantAddArgs)
		currentValue = currentValue + collections.Sum(relevantAddArgs)
		fmt.Println("Index", timeStepIndex, currentValue)
		lastTimeStepIndex = timeStepIndex
		strengthValues = append(strengthValues, currentValue*currentTimeStep)
	}
	return strengthValues
}

func computeCumulativeDuration(commands []Command) []int {
	durationSoFar := make([]int, len(commands))
	for i, command := range commands {
		if command.com == Noop {
			durationSoFar[i] = 1
		} else if command.com == Addx {
			durationSoFar[i] = 2
		}
		if i > 0 {
			durationSoFar[i] += durationSoFar[i-1]
		}
	}
	return durationSoFar
}

func findTimeStepIndex(timeStep int, timeSteps []int) int {
	if len(timeSteps) == 0 {
		return -1
	}
	middleIndex := len(timeSteps) / 2
	middle := timeSteps[middleIndex]
	if middle == timeStep {
		return middleIndex
	}
	if middle == timeStep+1 || middle == timeStep+2 {
		if middleIndex == 0 {
			return -1
		}
		prev := timeSteps[middleIndex-1]
		if prev <= timeStep {
			return middleIndex - 1
		}
		return math.Max(-1, middleIndex-2)
	}
	if middle > timeStep {
		return findTimeStepIndex(timeStep, timeSteps[:middleIndex])
	} else {
		return findTimeStepIndex(timeStep, timeSteps[middleIndex+1:]) + middleIndex + 1
	}
}

func renderScreen(commands []Command) string {
	screen := ""
	durationSoFar := computeCumulativeDuration(commands)
	currentCommandIndex := 0
	currentSpritePos := 1
	for timeStep := 0; timeStep < 240; timeStep++ {
		if durationSoFar[currentCommandIndex] <= timeStep {
			currentSpritePos += commands[currentCommandIndex].arg
			currentCommandIndex += 1
		}
		if timeStep%40 >= currentSpritePos-1 && timeStep%40 <= currentSpritePos+1 {
			screen += "▮"
		} else {
			screen += " "
		}
		if (timeStep+1)%40 == 0 {
			screen += "\n"
		}
	}
	return screen
}

func AoC10() {
	year := 2022
	day := 10
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	lineTuples := io.ReadInputFromRegexPerLine("(\\w+) ?(-?\\d+)?", 10, 2022)
	commands := collections.Map(lineTuples, func(t []string) Command {
		com := t[0]
		if com == "nooo" {
			return Command{com: com}
		}
		arg, _ := strconv.Atoi(t[1])
		return Command{com, arg}
	})
	fmt.Println(commands)
	fmt.Println("Part 1:")
	strengths := findSignalStrengths(commands)
	fmt.Println(strengths)
	totalStrength := collections.Sum(strengths)
	fmt.Println(totalStrength)
	requests.SubmitAnswer(day, year, totalStrength, 1)
	fmt.Println("Part 2:")
	screen := renderScreen(commands)
	fmt.Println(screen)
	requests.SubmitStringAnswer(day, year, "PZULBAUA", 2)

	fmt.Println("#  #\n## #\n## #\n# ##\n# ##\n#  #")
	fmt.Println()
	fmt.Println("####\n#   \n### \n#   \n#   \n####")

}
