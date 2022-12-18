package AoC_22_17

import (
	. "AoC/AoC2022/AoC_22_9"
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
	"os"
	"strconv"
)

type RockStructure struct {
	downOffsets []int
	upOffsets   []int
	height      int
	width       int
}

type Rock struct {
	rockType int
	position Position
}

var (
	rockStructures = []RockStructure{
		{
			downOffsets: []int{0, 0, 0, 0},
			upOffsets:   []int{0, 0, 0, 0},
			height:      1,
			width:       4,
		},
		{
			downOffsets: []int{1, 0, 1},
			upOffsets:   []int{-1, 0, -1},
			height:      3,
			width:       3,
		},
		{
			downOffsets: []int{0, 0, 0},
			upOffsets:   []int{-2, -2, 0},
			height:      3,
			width:       3,
		},
		{
			downOffsets: []int{0},
			upOffsets:   []int{0},
			height:      4,
			width:       1,
		},
		{
			downOffsets: []int{0, 0},
			upOffsets:   []int{0, 0},
			height:      2,
			width:       2,
		},
	}
)

func (rock *Rock) getStructure() RockStructure {
	return rockStructures[rock.rockType]
}

func (rock *Rock) blockedPositions() []Position {
	positions := make([]Position, 0)
	structure := rock.getStructure()
	for y := 0; y < structure.height; y++ {
		for x := 0; x < structure.width; x++ {
			if structure.downOffsets[x] <= y && structure.height+structure.upOffsets[x] > y {
				pos := Position{
					X: x + rock.position.X,
					Y: y + rock.position.Y,
				}
				positions = append(positions, pos)
			}
		}
	}
	return positions
}

// TODO remove dists
func (rock *Rock) isLyingWithDists(currentHeights []int) (bool, []int, []int) {
	/**
	dists only correct when not lying
	*/
	structure := rock.getStructure()
	dists := make([]int, structure.width)
	for x := 0; x < structure.width; x++ {
		height := currentHeights[x+rock.position.X]
		dists[x] = rock.position.Y + structure.downOffsets[x] - height
		if 1 == dists[x] {
			return true, nil, nil
		}
	}
	leftAndRightDist := make([]int, 2)
	if rock.position.X > 0 {
		leftAndRightDist[0] = rock.position.Y + structure.downOffsets[0] - currentHeights[rock.position.X-1]
	}
	if rock.position.X+structure.width < len(currentHeights) {
		leftAndRightDist[1] = rock.position.Y + structure.downOffsets[structure.width-1] - currentHeights[rock.position.X+structure.width]
	}
	return false, dists, leftAndRightDist
}

func parseDirections(input string) []Direction { // TODO move DIrection and Position to utils
	directions := make([]Direction, len(input))
	for i, str := range input {
		if str == '<' {
			directions[i] = Left
		}
		if str == '>' {
			directions[i] = Right
		}
	}
	return directions
}

func getWinds(directions []Direction, index int, count int) []Direction {
	from := index % len(directions)
	to := from + count
	if to > len(directions) {
		return append(directions[from:], directions[:to-len(directions)]...)
	}
	return directions[from:to]
}

func printCurrentConstellation(currentHeights []int, fallingRock *Rock) {
	topBorder := collections.MaxNumber(currentHeights)
	//bottomBorder := collections.MinNumber(currentHeights)
	if fallingRock != nil {
		topBorder = fallingRock.position.Y + fallingRock.getStructure().height
	}
	topBorder += 2
	for y := topBorder - 1; y >= 0; y-- {
		fmt.Println()
		for x := 0; x < 7; x++ {
			if y == 0 {
				fmt.Print("-")
			} else if currentHeights[x] >= y {
				fmt.Print("#")
			} else if collections.Contains(fallingRock.blockedPositions(), Position{X: x, Y: y}) {
				fmt.Print("@")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
}

func simulate(directions []Direction, count int) ([]int, int) {
	currentHeights := make([]int, 7)
	currentTop := 0
	windIndex := 0
	for i := 0; i < count; i++ {
		currentRockType := i % 5
		//rockStructure := rockStructures[currentRockType]
		windsUntilTop := getWinds(directions, windIndex, 4)
		//fmt.Println("Winds", windIndex, windsUntilTop)
		windIndex += 4
		rock := Rock{
			rockType: currentRockType,
			position: Position{X: 2, Y: currentTop + 4},
		}
		//printCurrentConstellation(currentHeights, &rock)
		for j, direction := range windsUntilTop {
			newPos := rock.position.Move(direction)
			if newPos.X+rock.getStructure().width <= 7 && newPos.X >= 0 {
				rock.position = newPos
			}
			if j < len(windsUntilTop)-1 {
				rock.position = rock.position.Move(Up) // We move Down, but system is inverted
			}
		}
		for true {
			lying, _, leftAndRightDist := rock.isLyingWithDists(currentHeights)
			if lying {
				break
			}
			rock.position = rock.position.Move(Up) // We move Down, but system is inverted
			direction := getWinds(directions, windIndex, 1)[0]
			//fmt.Println("Winds", windIndex, getWinds(directions, windIndex, 1))
			windIndex++
			if direction == Left && leftAndRightDist[0] > 1 || (direction == Right && leftAndRightDist[1] > 1) {
				rock.position = rock.position.Move(direction)
			}
		}
		//printCurrentConstellation(currentHeights, &rock)
		blockedPositions := rock.blockedPositions()
		for _, position := range blockedPositions {
			currentHeights[position.X] = position.Y
			if position.Y > currentTop {
				currentTop = position.Y
			}
		}
	}
	return currentHeights, currentTop
}

func AoC17() {
	year := 2022
	day := 17
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInput(17, 2022)
	fmt.Println(input)
	directions := parseDirections(input)
	fmt.Println("Part 1:")
	heights, top := simulate(directions, 2022)
	fmt.Println(heights, top)
	requests.SubmitAnswer(day, year, top, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
