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

type GridType uint8

const sandStartX = 500

const (
	Air GridType = iota
	Sand
	RockT
)

type RockLine struct {
	startX int
	startY int
	endX   int
	endY   int
}

func (g GridType) toString() string {
	switch g {
	case Air:
		return "."
	case Sand:
		return "o"
	case RockT:
		return "#"
	default:
		panic("Wrong grid type")
	}
}

func parseRockLines(lines []string) []RockLine {
	var rockLines []RockLine
	for _, line := range lines {
		linePointsString := strings.Split(line, " -> ")
		linePoints := collections.Map(linePointsString, func(str string) []int {
			return types.ToIntSlice(strings.Split(str, ","))
		})
		for i, point := range linePoints[:len(linePoints)-1] {
			startX, startY := collections.UnpackTwo(point)
			endX, endY := collections.UnpackTwo(linePoints[i+1])
			x := []int{startX, endX}
			y := []int{startY, endY}
			sort.Ints(x)
			sort.Ints(y)
			rockLines = append(rockLines, RockLine{x[0], y[0], x[1], y[1]})
		}
	}
	return rockLines
}

func createGrid(rockLines []RockLine, addBottom bool) (grid [][]GridType, shift int) {
	maxX := collections.MaxNumber(collections.Map(rockLines, func(line RockLine) int {
		return line.endX
	}))
	maxY := collections.MaxNumber(collections.Map(rockLines, func(line RockLine) int {
		return line.endY
	}))
	minX := collections.MinNumber(collections.Map(rockLines, func(line RockLine) int {
		return line.startX
	}))
	minX = math.Min(minX, sandStartX) - 1
	if addBottom {
		maxY += 2
		maxX += 1
		minX = math.Min(minX, sandStartX-maxY-1)
		maxX = math.Max(maxX, sandStartX+maxY)
		rockLines = append(rockLines, RockLine{minX, maxY, maxX + 1, maxY})
	}
	grid = make([][]GridType, maxY+1)
	maxX -= minX
	for i := range grid {
		grid[i] = make([]GridType, maxX+2)
	}
	for _, line := range rockLines {
		for j := line.startY; j < line.endY+1; j++ {
			for k := line.startX; k < line.endX+1; k++ {
				grid[j][k-minX] = RockT
			}
		}
	}
	return grid, -minX
}

func printGrid_(grid [][]GridType) {
	fmt.Println()
	for _, row := range grid {
		fmt.Println(strings.Join(collections.Map(row, func(gridType GridType) string {
			return gridType.toString()
		}), ""))
	}
}

func simulateSand(grid [][]GridType, shift int, show bool, stopAt50 bool) [][]GridType {
	newGrid := make([][]GridType, len(grid))
	for i, row := range grid {
		newGrid[i] = make([]GridType, len(grid[i]))
		copy(newGrid[i], row)
	}
	for true {
		sandPosition := Position{sandStartX + shift, -1}
		for true {
			var newPos *Position
			for _, direction := range []Direction{Down, DownLeft, DownRight} {
				possiblePos := sandPosition.move(direction)
				if possiblePos.y >= len(newGrid) {
					newGrid[sandPosition.y][sandPosition.x] = Air
					if show {
						printGrid_(newGrid)
					}
					return newGrid
				}
				if newGrid[possiblePos.y][possiblePos.x] == Air {
					newPos = &possiblePos
					break
				}
			}
			if newPos == nil || *newPos == sandPosition {
				if stopAt50 {
					if sandPosition.x == sandStartX+shift && sandPosition.y == 0 {
						if show {
							printGrid_(newGrid)
						}
						return newGrid
					}
				}
				break
			}
			if sandPosition.y >= 0 {
				newGrid[sandPosition.y][sandPosition.x] = Air
			}
			sandPosition = *newPos
			newGrid[sandPosition.y][sandPosition.x] = Sand
			if show {
				printGrid_(newGrid)
			}
		}
	}
	return newGrid
}

func countSand(grid [][]GridType) int {
	return collections.Sum(collections.Map(grid, func(line []GridType) int {
		return len(collections.Filter(line, func(value GridType) bool {
			return value == Sand
		}))
	}))
}

func AoC14() {
	year := 2022
	day := 14
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputLines(14, 2022)
	fmt.Println(input)
	rockLines := parseRockLines(input)
	fmt.Println(rockLines)
	grid, shift := createGrid(rockLines, false)
	fmt.Println("Part 1:")
	printGrid_(grid)
	grid = simulateSand(grid, shift, false, false)
	totalSand := countSand(grid)
	fmt.Println(totalSand)
	requests.SubmitAnswer(day, year, totalSand, 1)

	fmt.Println("Part 2:")
	grid, shift = createGrid(rockLines, true)
	printGrid_(grid)
	grid = simulateSand(grid, shift, true, true)
	totalSand = countSand(grid)
	fmt.Println(totalSand)
	requests.SubmitAnswer(day, year, totalSand, 2)
}
