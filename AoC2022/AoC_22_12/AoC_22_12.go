package AoC_22_12

import (
	_9 "AoC/AoC2022/AoC_22_9"
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"AoC/utils/search"
	"fmt"
)

func findPosOfValue(grid [][]string, toFind string) _9.Position {
	for rowIndex, row := range grid {
		for column, value := range row {
			if value == toFind {
				return _9.Position{column, rowIndex}
			}
		}
	}
	panic("No start")
}

func getNeighboursFunc(grid [][]string) func(pos _9.Position) []search.Neighbour[_9.Position] {
	return func(pos _9.Position) []search.Neighbour[_9.Position] {
		possibleDirs := []_9.Direction{_9.Up, _9.Right, _9.Left, _9.Down}
		possibleNeighbourPositions := collections.Map(possibleDirs, func(dir _9.Direction) _9.Position {
			return pos.Move(dir)
		})
		var neighbours []search.Neighbour[_9.Position]
		currentHeight := getHeight(grid, pos)
		for _, neighbourPos := range possibleNeighbourPositions {
			if neighbourPos.X < 0 || neighbourPos.X >= len(grid[0]) || neighbourPos.Y < 0 || neighbourPos.Y >= len(grid) {
				continue
			}
			neighbourHeight := getHeight(grid, neighbourPos)
			if neighbourHeight-currentHeight <= 1 {
				// height diff is not bigger than 1 or going down
				neighbours = append(neighbours, search.Neighbour[_9.Position]{neighbourPos, 1})
			}
		}
		return neighbours
	}
}

func getHeight(grid [][]string, position _9.Position) int {
	stringHeight := grid[position.Y][position.X]
	if stringHeight == "S" {
		stringHeight = "a"
	}
	if stringHeight == "E" {
		stringHeight = "z"
	}
	return int(stringHeight[0])
}

func getHeuristic(goalPos _9.Position, grid [][]string) func(_9.Position) int {
	return func(pos _9.Position) int {
		horizontalDist := math.Abs(pos.X-goalPos.X) + math.Abs(pos.Y-goalPos.Y)
		verticalDist := math.Max(0, getHeight(grid, goalPos)-getHeight(grid, pos))
		return math.Max(horizontalDist, verticalDist)
	}
}

func findShortestPathFromAnyLowest(grid [][]string) []_9.Position {
	var shortestPath []_9.Position
	goal := findPosOfValue(grid, "E")
	for rowIndex, row := range grid {
		for column, value := range row {
			if value == "a" || value == "S" {
				start := _9.Position{column, rowIndex}
				fmt.Println("Try", start)
				path := search.AStar(start, getNeighboursFunc(grid), func(pos _9.Position) bool { return pos == goal }, getHeuristic(goal, grid))
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
	path := search.AStar(start, getNeighboursFunc(grid), func(pos _9.Position) bool { return pos == goal }, getHeuristic(goal, grid))
	fmt.Println(path)
	fmt.Println(len(path))
	requests.SubmitAnswer(day, year, len(path)-1, 1)
	fmt.Println("Part 2:")
	path = findShortestPathFromAnyLowest(grid)
	fmt.Println(len(path))
	requests.SubmitAnswer(day, year, len(path)-1, 2)
}
