package AoC_22_18

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"AoC/utils/search"
	"fmt"
	"os"
	"strconv"
)

func getSidesOfCube(cube []int) [][]int {
	x, y, z := collections.UnpackThree(cube)
	sides := [][]int{
		{x, y, z, 0},
		{x + 1, y, z, 0},
		{x, y, z, 1},
		{x, y + 1, z, 1},
		{x, y, z, 2},
		{x, y, z + 1, 2},
	}
	return sides
}

func getCubeSurface(cubes [][]int) [][]int {
	seenSides := make(map[[4]int]int)
	for _, cube := range cubes {
		sides := getSidesOfCube(cube)
		for _, side := range sides {
			sideArray := arrayFromSide(side)
			seenSides[sideArray] += 1
		}
	}
	surfaceSides := make([][]int, 0)
	for side, count := range seenSides {
		if count == 1 {
			surfaceSides = append(surfaceSides, side[:])
		}
	}
	return surfaceSides
}

func arrayFromSide(side []int) [4]int {
	var sideArray [4]int
	copy(sideArray[:], side)
	return sideArray
}

func getOuterSides(sides [][4]int, zeroSide [4]int) [][4]int {
	outerSides := collections.NewSet[[4]int]()
	allSurfaceSides := collections.NewSet[[4]int]()
	allSurfaceSides.AddMultiple(sides)
	for _, currentSide := range sides {
		if outerSides.Has(currentSide) {
			continue
		}
		path := search.AStar(zeroSide, func(currentSide [4]int) []search.Neighbour[[4]int] {
			dim := currentSide[3]
			var neighbours [][4]int
			for newDim := 0; newDim < 3; newDim++ {
				for offsetInNewDim := -1; offsetInNewDim <= 1; offsetInNewDim += 2 {
					if dim == newDim {
						var neighbour [4]int
						copy(neighbour[:], currentSide[:])
						neighbour[dim] += offsetInNewDim
						neighbours = append(neighbours, neighbour)
					} else {
						for offsetInOldDim := 0; offsetInOldDim <= 1; offsetInOldDim++ {
							var neighbour [4]int
							copy(neighbour[:], currentSide[:])
							neighbour[dim] += offsetInOldDim
							neighbour[newDim] += offsetInNewDim
							neighbours = append(neighbours, neighbour)
						}
					}
				}
			}
			neighbours = collections.Filter(neighbours, func(neighbour [4]int) bool { return !allSurfaceSides.Has(neighbour) || neighbour == currentSide })
			return collections.Map(neighbours, func(neighbour [4]int) search.Neighbour[[4]int] {
				return search.Neighbour[[4]int]{neighbour, 1}
			})
		}, func(side [4]int) bool {
			return side == currentSide
		}, func(side [4]int) int {
			return collections.Sum(collections.Map(side[:], func(number int) int {
				return math.Abs(number)
			}))
		})
		fmt.Println(path)
		os.Exit(1)
	}
	return sides
}

func AoC18() {
	year := 2022
	day := 18
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	cubes := io.ReadInputFromRegexPerLineInt("(\\d+),(\\d+),(\\d+)", 18, 2022)
	fmt.Println(cubes)
	fmt.Println("Part 1:")
	surfaceSides := getCubeSurface(cubes)
	totalSeenSurfaces := len(surfaceSides)
	fmt.Println(totalSeenSurfaces)
	requests.SubmitAnswer(day, year, totalSeenSurfaces, 1)

	fmt.Println("Part 2:")
	getOuterSides(collections.Map(surfaceSides, func(side []int) [4]int {
		var array [4]int
		copy(array[:], side)
		return array
	}), [4]int{0, 0, 0, 0})
	// requests.SubmitAnswer(day, year, 0, 2)
}
