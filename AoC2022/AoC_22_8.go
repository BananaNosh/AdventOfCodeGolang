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
	for i := range columnMax {
		columnMax[i] = -1
	}
	for i, row := range grid {
		rowMax := -1
		for j, tree := range row {
			if tree > rowMax {
				rowMax = tree
				scanLeft[i] = append(scanLeft[i], 1)
			} else {
				scanLeft[i] = append(scanLeft[i], 0)
			}
			if tree > columnMax[j] {
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
	return ewiseFunc(grid1, grid2, func(l int, r int) int {
		return math.Min(1, l+r)
	})
}

func ewiseMul(grid1 [][]int, grid2 [][]int) [][]int {
	return ewiseFunc(grid1, grid2, func(l int, r int) int {
		return l * r
	})
}

func ewiseFunc(grid1 [][]int, grid2 [][]int, f func(int, int) int) [][]int {
	result := make([][]int, len(grid1))
	for i, row1 := range grid1 {
		for j, value1 := range row1 {
			result[i] = append(result[i], f(value1, grid2[i][j]))
		}
	}
	return result
}

func countSeenTrees(grid [][]int) int {
	leftScan, topScan := scanMap(grid, false)
	rightScan, bottomScan := scanMap(grid, true)
	//printGrid(leftScan)
	//printGrid(topScan)
	//printGrid(rightScan)
	//printGrid(bottomScan)
	seenTrees := ewiseOr(leftScan, topScan)
	seenTrees = ewiseOr(rightScan, seenTrees)
	seenTrees = ewiseOr(bottomScan, seenTrees)
	//fmt.Println("SeenTrees")
	//printGrid(seenTrees)
	seenCount := collections.Sum(collections.Map(seenTrees, func(row []int) int {
		return collections.Sum(row)
	}))
	return seenCount
}

func scanForTreeHouse(grid [][]int, reverse bool) ([][]int, [][]int) {
	scanLeft := make([][]int, len(grid))
	scanTop := make([][]int, len(grid))
	if reverse {
		grid = reversed(grid)
	}
	columnMaxs := make([]map[int]int, len(grid[0]))
	for i := range columnMaxs {
		columnMaxs[i] = make(map[int]int)
	}
	for i, row := range grid {
		rowMax := make(map[int]int)
		for j, tree := range row {
			if collections.HasKey(rowMax, tree) {
				scanLeft[i] = append(scanLeft[i], j-rowMax[tree])
			} else {
				scanLeft[i] = append(scanLeft[i], j)
			}
			if collections.HasKey(columnMaxs[j], tree) {
				scanTop[i] = append(scanTop[i], i-columnMaxs[j][tree])
			} else {
				scanTop[i] = append(scanTop[i], i)
			}
			for k := 0; k <= tree; k++ {
				rowMax[k] = j
				columnMaxs[j][k] = i
			}
		}
	}
	if reverse {
		return reversed(scanLeft), reversed(scanTop)
	}
	return scanLeft, scanTop
}

func findTreeHouse(grid [][]int) int {
	left, top := scanForTreeHouse(grid, false)
	right, bottom := scanForTreeHouse(grid, true)
	printGrid(left)
	printGrid(top)
	printGrid(right)
	printGrid(bottom)
	treeScore := ewiseMul(left, top)
	treeScore = ewiseMul(treeScore, right)
	treeScore = ewiseMul(treeScore, bottom)
	printGrid(treeScore)
	bestScore := collections.MaxNumber(collections.Map(treeScore, func(row []int) int {
		return collections.MaxNumber(row)
	}))
	return bestScore
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
	seenCount := countSeenTrees(input)
	fmt.Println("seen", seenCount)
	requests.SubmitAnswer(day, year, seenCount, 1)

	fmt.Println("Part 2:")
	score := findTreeHouse(input)
	requests.SubmitAnswer(day, year, score, 2)
}
