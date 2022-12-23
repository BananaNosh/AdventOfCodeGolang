package AoC_22_9

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"fmt"
	"strconv"
)

type Position struct {
	X int
	Y int
}

type Rope struct {
	headPos               Position
	notHeadKnotsPositions []Position
	seenTailPositions     *collections.Set[Position]
}

type Instruction struct {
	direction Direction
	count     int
}

type Direction int

const (
	Up Direction = iota
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

func DirectionFromString(str string) Direction {
	switch str {
	case "U":
		return Up
	case "R":
		return Right
	case "D":
		return Down
	case "L":
		return Left
	}
	panic("No such Direction")
}

func (pos Position) Move(direction Direction) Position {
	xPos := pos.X
	yPos := pos.Y
	switch direction {
	case Up:
		return Position{xPos, yPos - 1}
	case UpRight:
		return Position{xPos + 1, yPos - 1}
	case Right:
		return Position{xPos + 1, yPos}
	case DownRight:
		return Position{xPos + 1, yPos + 1}
	case Down:
		return Position{xPos, yPos + 1}
	case DownLeft:
		return Position{xPos - 1, yPos + 1}
	case Left:
		return Position{xPos - 1, yPos}
	case UpLeft:
		return Position{xPos - 1, yPos - 1}
	}
	panic("Wrong direction")
}

func (dir Direction) Turn(clockWise bool, countInEighth int) Direction {
	newDirInt := int(dir)
	if clockWise {
		newDirInt += countInEighth
	} else {
		newDirInt += 8 - countInEighth
	}
	newDirInt %= 8
	return Direction(newDirInt)
}

func posRelativeToOther(pos, otherPos Position) Position {
	return Position{pos.X - otherPos.X, pos.Y - otherPos.Y}
}

func (r *Rope) move(instruction Instruction) {
	dir := instruction.direction
	count := instruction.count
	for i := 0; i < count; i++ {
		r.headPos = r.headPos.Move(dir)
		currentPrev := r.headPos
		for i, knotPos := range r.notHeadKnotsPositions {
			relativePos := posRelativeToOther(knotPos, currentPrev)
			knotDirection, moved := knotDirectionFromRelativePos(relativePos)
			if moved {
				knotPos = knotPos.Move(knotDirection)
				r.notHeadKnotsPositions[i] = knotPos
			}
			currentPrev = knotPos
		}
		r.seenTailPositions.Add(collections.Last(r.notHeadKnotsPositions))
		if len(r.notHeadKnotsPositions) < 5 {
			fmt.Println("head", r.headPos, "tail", collections.Last(r.notHeadKnotsPositions))
		} else {
			fmt.Println("head", r.headPos, "tail", r.notHeadKnotsPositions[4])
		}
	}
}

func knotDirectionFromRelativePos(relativePos Position) (Direction, bool) {
	xDist := math.Abs(relativePos.X)
	yDist := math.Abs(relativePos.Y)
	var tailDir Direction
	if yDist == 2 && xDist == 0 {
		if relativePos.Y > 0 {
			tailDir = Up
		} else {
			tailDir = Down
		}
	} else if xDist == 2 && yDist == 0 {
		if relativePos.X > 0 {
			tailDir = Left
		} else {
			tailDir = Right
		}
	} else if xDist > 1 || yDist > 1 { // Needs to go dioganally up or down
		if relativePos.Y > 0 {
			if relativePos.X > 0 {
				tailDir = UpLeft
			} else {
				tailDir = UpRight
			}
		} else {
			if relativePos.X > 0 {
				tailDir = DownLeft
			} else {
				tailDir = DownRight
			}

		}
	} else {
		return 0, false
	}
	return tailDir, true
}

func simulate(rope Rope, instructions []Instruction) *collections.Set[Position] {
	rope.seenTailPositions.Add(collections.Last(rope.notHeadKnotsPositions))
	for _, instruction := range instructions {
		rope.move(instruction)
	}
	return rope.seenTailPositions
}

func AoC9() {
	year := 2022
	day := 9
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	instructionTuples := io.ReadInputFromRegexPerLine("(\\w) (\\d+)", 9, 2022)
	instructions := collections.Map(instructionTuples, func(tuple []string) Instruction {
		count, _ := strconv.Atoi(tuple[1])
		return Instruction{direction: DirectionFromString(tuple[0]), count: count}
	})
	fmt.Println("Part 1:")
	fmt.Println(instructions)
	seenPositions := simulate(newRope(2), instructions)
	fmt.Println(seenPositions, seenPositions.Size())
	requests.SubmitAnswer(day, year, seenPositions.Size(), 1)

	fmt.Println("Part 2:")
	seenPositions = simulate(newRope(10), instructions)
	fmt.Println(seenPositions, seenPositions.Size())
	requests.SubmitAnswer(day, year, seenPositions.Size(), 2)
}

func newRope(knotsCount int) Rope {
	return Rope{notHeadKnotsPositions: make([]Position, knotsCount-1), seenTailPositions: collections.NewSet[Position]()}
}
