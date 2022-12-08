package AoC22

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"fmt"
)

func scanMap(grid [][]int, reverse bool) ([][]int, [][]int) {
	scanLeft := make([][]int, len(grid))
	scanTop := make([][]int, len(grid))
	if reverse {
		grid = reversed(grid)
	}
	columnMax := make([]int, len(grid[0]))
	for i, row := range grid {
		rowMax := -1
		for j, tree := range row {
			if tree > rowMax {
				rowMax = tree
				scanLeft[i] = append(scanLeft[i], 1)
			} else {
				scanLeft[i] = append(scanLeft[i], 0)
			}
			if tree > columnMax[j] || columnMax[j] == 0 {
				columnMax[j] = tree
				scanTop[i] = append(scanTop[i], 1)
			} else {
				scanTop[i] = append(scanTop[i], 0)
			}
		}
	}
	if reverse {
		return reversed(scanLeft), reversed(scanTop)
	}
	return scanLeft, scanTop
}

func reversed[T any](grid [][]T) [][]T {
	reversedGrid := make([][]T, len(grid))
	for i := range grid {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			reversedGrid[i] = append(reversedGrid[i], grid[len(grid)-1-i][j])
		}
	}
	return reversedGrid
}

func printGrid[T any](grid [][]T) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
}

func ewiseOr(grid1 [][]int, grid2 [][]int) [][]int {
	result := make([][]int, len(grid1))
	for i, row1 := range grid1 {
		for j, value1 := range row1 {
			result[i] = append(result[i], math.Min(1, value1+grid2[i][j]))
		}
	}
	return result
}

func AoC8() {
	year := 2022
	day := 8
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputAs2DInts(8, 2022)
	printGrid(input)
	fmt.Println("Part 1:")
	leftScan, topScan := scanMap(input, false)
	rightScan, bottomScan := scanMap(input, true)
	printGrid(leftScan)
	printGrid(topScan)
	printGrid(rightScan)
	printGrid(bottomScan)
	seenTrees := ewiseOr(leftScan, topScan)
	seenTrees = ewiseOr(rightScan, seenTrees)
	seenTrees = ewiseOr(bottomScan, seenTrees)
	fmt.Println("SeenTrees")
	printGrid(seenTrees)
	seenCount := collections.Sum(collections.Map(seenTrees, func(row []int) int {
		return collections.Sum(row)
	}))
	fmt.Println("seen", seenCount)
	requests.SubmitAnswer(day, year, seenCount, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
