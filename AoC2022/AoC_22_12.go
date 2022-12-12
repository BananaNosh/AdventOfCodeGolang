package AoC22

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"AoC/utils/search"
	"fmt"
)

func findPosOfValue(grid [][]string, toFind string) Position {
	for rowIndex, row := range grid {
		for column, value := range row {
			if value == toFind {
				return Position{column, rowIndex}
			}
		}
	}
	panic("No start")
}

func getNeighboursFunc(grid [][]string) func(pos Position) []search.Neighbour[Position] {
	return func(pos Position) []search.Neighbour[Position] {
		possibleDirs := []Direction{Up, Right, Left, Down}
		possibleNeighbourPositions := collections.Map(possibleDirs, func(dir Direction) Position {
			return pos.move(dir)
		})
		var neighbours []search.Neighbour[Position]
		currentHeight := getHeight(grid, pos)
		for _, neighbourPos := range possibleNeighbourPositions {
			if neighbourPos.x < 0 || neighbourPos.x >= len(grid[0]) || neighbourPos.y < 0 || neighbourPos.y >= len(grid) {
				continue
			}
			neighbourHeight := getHeight(grid, neighbourPos)
			if neighbourHeight-currentHeight <= 1 {
				// height diff is not bigger than 1 or going down
				neighbours = append(neighbours, search.Neighbour[Position]{neighbourPos, 1})
			}
		}
		return neighbours
	}
}

func getHeight(grid [][]string, position Position) int {
	stringHeight := grid[position.y][position.x]
	if stringHeight == "S" {
		stringHeight = "a"
	}
	if stringHeight == "E" {
		stringHeight = "z"
	}
	return int(stringHeight[0])
}

func getHeuristic(goalPos Position, grid [][]string) func(Position) int {
	return func(pos Position) int {
		horizontalDist := math.Abs(pos.x-goalPos.x) + math.Abs(pos.y-goalPos.y)
		verticalDist := math.Max(0, getHeight(grid, goalPos)-getHeight(grid, pos))
		return math.Max(horizontalDist, verticalDist)
	}
}

func findShortestPathFromAnyLowest(grid [][]string) []Position {
	var shortestPath []Position
	goal := findPosOfValue(grid, "E")
	for rowIndex, row := range grid {
		for column, value := range row {
			if value == "a" || value == "S" {
				start := Position{column, rowIndex}
				fmt.Println("Try", start)
				path := search.AlphaStar(start, getNeighboursFunc(grid), func(pos Position) bool { return pos == goal }, getHeuristic(goal, grid))
				if len(path) == 0 {
					continue
				}
				if len(shortestPath) == 0 || len(shortestPath) > len(path) {
					shortestPath = path
					fmt.Println("Found at ", start)
				}
			}
		}
	}
	return shortestPath
}

func AoC12() {
	year := 2022
	day := 12
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	grid := io.ReadInputAs2DStrings(12, 2022)
	fmt.Println(grid)
	start := findPosOfValue(grid, "S")
	goal := findPosOfValue(grid, "E")
	fmt.Println(start, goal)
	fmt.Println("Part 1:")
	path := search.AlphaStar(start, getNeighboursFunc(grid), func(pos Position) bool { return pos == goal }, getHeuristic(goal, grid))
	fmt.Println(path)
	fmt.Println(len(path))
	requests.SubmitAnswer(day, year, len(path)-1, 1)
	fmt.Println("Part 2:")
	path = findShortestPathFromAnyLowest(grid)
	fmt.Println(len(path))
	requests.SubmitAnswer(day, year, len(path)-1, 2)
}
