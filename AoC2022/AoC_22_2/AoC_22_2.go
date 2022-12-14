package AoC_22_2

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
)

type GameDecision int
type Outcome int

const (
	Rock     GameDecision = iota // 0
	Paper                 = iota // 1
	Scissors              = iota // 2
)

const (
	Lost Outcome = iota
	Draw
	Won
)

func (d GameDecision) play(other GameDecision) Outcome {
	return Outcome((int(d) + 3 - int(other) + 1) % 3)
}

func (o Outcome) points() int {
	return int(o) * 3
}

func (o GameDecision) points() int {
	return int(o) + 1
}

func StringToGameDecision(s string) GameDecision {
	switch s {
	case "A":
		fallthrough
	case "X":
		return Rock
	case "B":
		fallthrough
	case "Y":
		return Paper
	case "C":
		fallthrough
	case "Z":
		return Scissors
	default:
		panic(fmt.Sprintf("No such GameDecision %s", s))
	}
}

func StringToOutcome(s string) Outcome {
	switch s {
	case "X":
		return Lost
	case "Y":
		return Draw
	case "Z":
		return Won
	default:
		panic(fmt.Sprintf("No such Outcome %s", s))
	}
}

func AoC2() {
	year := 2022
	day := 2
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))
	//input := io.ReadInputFromRegexPerLine("(A|B|C) (X|Y|Z)", 2, 2022)
	input := io.ReadInputFromRegexPerLine("\\w", 2, 2022)
	decisions := collections.Map(input, func(str []string) []GameDecision {
		return collections.Map(str, StringToGameDecision)
	})
	outcomes := collections.Map(decisions, func(decision []GameDecision) Outcome {
		return decision[1].play(decision[0])
	})
	total := collections.ReduceWithIndex(outcomes, func(i int, points int, outcome Outcome) int {
		result := outcome.points() + decisions[i][1].points()
		//fmt.Println(outcome, decisions[i], result)
		return result + points
	}, 0)
	fmt.Println("Part 1:")
	//fmt.Println(input)
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 1)
	fmt.Println("Part 2:")
	outcomes = collections.Map(input, func(str []string) Outcome {
		return StringToOutcome(str[1])
	})
	ownDecisions := collections.MapWithIndex(decisions, func(i int, playerDecs []GameDecision) GameDecision {
		opponentDec := playerDecs[0]
		wantedOut := outcomes[i]
		return GameDecision(((int(wantedOut)+3-1)%3 + int(opponentDec)) % 3)
	})
	//fmt.Println(ownDecisions)
	total = collections.ReduceWithIndex(outcomes, func(i int, points int, outcome Outcome) int {
		result := outcome.points() + ownDecisions[i].points()
		//fmt.Println(outcome, decisions[i][0], ownDecisions[i], result)
		return result + points
	}, 0)
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 2)
}
