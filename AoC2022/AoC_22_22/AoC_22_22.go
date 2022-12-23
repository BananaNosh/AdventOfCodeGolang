package AoC_22_22

import (
	. "AoC/AoC2022/AoC_22_9"
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type GridType int

const (
	NONE GridType = iota
	FREE
	WALL
)

func gridTypeFromString(str string) GridType {
	switch str {
	case " ":
		return NONE
	case ".":
		return FREE
	case "#":
		return WALL
	default:
		panic("No such GridType: " + str)
	}
}

func createMap(gridString [][]string) (grid [][]GridType, xRanges [][]int, yRanges [][]int) {
	maxWidth := collections.MaxNumber(collections.Map(gridString, func(row []string) int {
		return len(row)
	}))
	height := len(gridString)
	xRanges = make([][]int, height)
	yRanges = make([][]int, maxWidth)
	grid = make([][]GridType, height)
	for i := 0; i < maxWidth; i++ {
		yRanges[i] = []int{-1, -1}
	}
	for y, row := range gridString {
		xRanges[y] = []int{-1, -1}
		grid[y] = make([]GridType, maxWidth)
		for x, value := range row {
			val := gridTypeFromString(value)
			grid[y][x] = val
			if val != NONE {
				if xRanges[y][0] == -1 {
					xRanges[y][0] = x
				}
				xRanges[y][1] = x
				if yRanges[x][0] == -1 {
					yRanges[x][0] = y
				}
				yRanges[x][1] = y
			}
		}
	}
	return grid, xRanges, yRanges
}

func getOrders(str string) ([]int, []Direction) {
	var directions []Direction
	var dists []int
	pattern := regexp.MustCompile("\\d+|L|R")
	allStrings := pattern.FindAllString(str, -1)
	directions = append(directions, Right)
	if _, err := strconv.Atoi(str[:1]); err != nil {
		dists = append(dists, 0)
	}
	for _, submatch := range allStrings {
		if number, err := strconv.Atoi(submatch); err == nil {
			dists = append(dists, number)
		} else {
			newDir := collections.Last(directions).Turn(DirectionFromString(submatch) == Right, 2)
			directions = append(directions, newDir)
		}
	}
	if len(dists) != len(directions) {
		panic("Dists and dirs should be same length")
	}
	return dists, directions
}

func pathTroughGrid(grid [][]GridType, xRanges [][]int, yRanges [][]int, dists []int, directions []Direction, shouldPrint bool) (Position, Direction) {
	currentPos := Position{xRanges[0][0], 0}
	path := []Position{currentPos}
	facing := []Direction{Right}
	for i := 0; i < len(dists); i++ {
		dist := dists[i]
		dir := directions[i]
		for j := 0; j < dist; j++ {
			if shouldPrint {
				printGrid(grid, path, facing)
			}
			newPos := currentPos.Move(dir)
			currentXRange := xRanges[currentPos.Y]
			currentWidth := currentXRange[1] + 1 - currentXRange[0]
			newPos.X = (newPos.X-currentXRange[0]+currentWidth)%currentWidth + currentXRange[0]
			currentYRange := yRanges[currentPos.X]
			currentHeight := currentYRange[1] + 1 - currentYRange[0]
			newPos.Y = (newPos.Y-currentYRange[0]+currentHeight)%currentHeight + currentYRange[0]
			nextField := grid[newPos.Y][newPos.X]
			if nextField == NONE {
				panic("wrapping should prevent NONE")
			}
			if nextField != FREE {
				break
			}
			currentPos = newPos
			if shouldPrint {
				path = append(path, currentPos)
				facing = append(facing, dir)
			}
			fmt.Println(newPos)
		}
	}
	return currentPos, collections.Last(directions)
}

func printGrid(grid [][]GridType, positions []Position, directions []Direction) {
	posMap := make(map[Position]int)
	for i, position := range positions {
		posMap[position] = i
	}
	for y, row := range grid {
		fmt.Println()
		for x, gridType := range row {
			pos := Position{X: x, Y: y}
			if collections.HasKey(posMap, pos) {
				switch directions[posMap[pos]] {
				case Up:
					fmt.Print("^")
				case Right:
					fmt.Print(">")
				case Down:
					fmt.Print("v")
				case Left:
					fmt.Print("<")
				default:
					panic("Wrong dir")
				}
			} else {
				switch gridType {
				case NONE:
					fmt.Print(" ")
				case FREE:
					fmt.Print(".")
				case WALL:
					fmt.Print("#")
				default:
					panic("Wrong gridtype")
				}
			}
		}
	}
}

func valueForEndState(position Position, dir Direction) int {
	dirVal := (int(dir)/2 + 3) % 4
	return 1000*(position.Y+1) + 4*(position.X+1) + dirVal
}

func AoC22() {
	year := 2022
	day := 22
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputAs2DStrings(22, 2022)
	fmt.Println(input)
	fmt.Println("Part 1:")
	grid, xRanges, yRanges := createMap(input[:len(input)-2])
	fmt.Println(grid)
	fmt.Println(xRanges)
	fmt.Println(yRanges)
	dists, directions := getOrders(strings.Join(collections.Last(input), ""))
	fmt.Println(dists)
	fmt.Println(directions)
	position, facing := pathTroughGrid(grid, xRanges, yRanges, dists, directions, false)
	result := valueForEndState(position, facing)
	fmt.Println(result)
	requests.SubmitAnswer(day, year, result, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
