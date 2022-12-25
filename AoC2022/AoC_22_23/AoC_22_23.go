package AoC_22_23

import (
	. "AoC/AoC2022/AoC_22_9"
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
)

func ParseElfPositions(grid [][]string) *collections.Set[Position] {
	elfSet := collections.NewSet[Position]()
	for y, row := range grid {
		for x, val := range row {
			if val == "#" {
				elfSet.Add(Position{x, y})
			}
		}
	}
	return elfSet
}

func moveElves(elves *collections.Set[Position], rounds int) (*collections.Set[Position], int) {
	elfPositions := elves.Copy()
	directions := []Direction{Up, Down, Left, Right}
	currentFirstDir := 0
	dirsToCheck := [][]Direction{{UpLeft, Up, UpRight}, {DownLeft, Down, DownRight}, {UpLeft, Left, DownLeft}, {UpRight, Right, DownRight}}
	//printElves(elfPositions)
	_rounds := rounds
	if _rounds == 0 {
		_rounds = 1000000
	}
	for i := 0; i < _rounds; i++ {
		proposedPositions := make(map[Position]Position)
		proposedCount := make(map[Position]int)
		for _, position := range elfPositions.AsSlice() {
			allFree := true
			for dirToCheck := range dirsToCheck {
				for _, direction := range dirsToCheck[dirToCheck] {
					newPosition := position.Move(direction)
					if elfPositions.Has(newPosition) {
						allFree = false
						break
					}
				}
				if !allFree {
					break
				}
			}
			if allFree {
				continue
			}
			for dirToCheck := currentFirstDir; dirToCheck < currentFirstDir+4; dirToCheck++ {
				dirToCheckMod := dirToCheck % 4
				free := true
				for _, direction := range dirsToCheck[dirToCheckMod] {
					newPosition := position.Move(direction)
					if elfPositions.Has(newPosition) {
						free = false
						break
					}
				}
				if free {
					newPosition := position.Move(directions[dirToCheckMod])
					proposedPositions[position] = newPosition
					proposedCount[newPosition] += 1
					break
				}
			}
		}
		for current, newPos := range proposedPositions {
			if proposedCount[newPos] == 1 {
				//fmt.Println("from ", current)
				//fmt.Println("to ", newPos)
				elfPositions.Remove(current)
				elfPositions.Add(newPos)
			}
		}
		currentFirstDir += 1
		//fmt.Print("\nRound ", i+1)
		//printElves(elfPositions)
		if len(proposedPositions) == 0 {
			return elfPositions, i + 1
		}

	}
	return elfPositions, rounds
}

func printElves(elvePositions *collections.Set[Position]) {
	maxX, minX, maxY, minY := boundsFromPositions(elvePositions)
	fmt.Println()
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if elvePositions.Has(Position{x, y}) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func boundsFromPositions(elvePositions *collections.Set[Position]) (int, int, int, int) {
	slice := elvePositions.AsSlice()
	xPositions := collections.Map(slice, func(pos Position) int {
		return pos.X
	})
	yPositions := collections.Map(slice, func(pos Position) int {
		return pos.Y
	})
	maxX := collections.MaxNumber(xPositions)
	minX := collections.MinNumber(xPositions)
	maxY := collections.MaxNumber(yPositions)
	minY := collections.MinNumber(yPositions)
	return maxX, minX, maxY, minY
}

func AoC23() {
	year := 2022
	day := 23
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputAs2DStrings(23, 2022)
	fmt.Println(input)
	fmt.Println("Part 1:")
	elfSet := ParseElfPositions(input)
	newElfSet, round := moveElves(elfSet, 10)
	count := elfSet.Size()
	maxX, minX, maxY, minY := boundsFromPositions(newElfSet)
	empty := (maxX+1-minX)*(maxY+1-minY) - count
	fmt.Println(empty, round)
	requests.SubmitAnswer(day, year, empty, 1)

	fmt.Println("Part 2:")
	_, round = moveElves(elfSet, 0)
	fmt.Println(round)
	requests.SubmitAnswer(day, year, round, 2)
}
