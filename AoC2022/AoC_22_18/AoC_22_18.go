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
	"sort"
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

func getCubeSurface(cubes [][]int) [][4]int {
	seenSides := make(map[[4]int]int)
	for _, cube := range cubes {
		sides := getSidesOfCube(cube)
		for _, side := range sides {
			sideArray := arrayFromSide(side)
			seenSides[sideArray] += 1
		}
	}
	surfaceSides := make([][4]int, 0)
	for side, count := range seenSides {
		if count == 1 {
			surfaceSides = append(surfaceSides, side)
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

func printFlooding(cubes *collections.Set[[3]int], flooded *collections.Set[[3]int], minBound [3]int, maxBound [3]int) {
	for z := minBound[2]; z <= maxBound[2]; z++ {
		fmt.Println()
		fmt.Println("Z:", z)
		printFloodingSlice(z, cubes, flooded, minBound, maxBound)
	}
}

func printFloodingSlice(zDim int, cubes *collections.Set[[3]int], flooded *collections.Set[[3]int], minBound [3]int, maxBound [3]int) {
	for y := minBound[1]; y <= maxBound[1]; y++ {
		for x := minBound[0]; x <= maxBound[0]; x++ {
			cube := [3]int{x, y, zDim}
			if cubes.Has(cube) {
				print("#")
			} else if flooded.Has(cube) {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func floodCubes(cubes [][3]int, minCube [3]int, maxCube [3]int) *collections.Set[[4]int] {
	cubeSet := collections.NewSet[[3]int]()
	touchedSides := collections.NewSet[[4]int]()
	flooded := collections.NewSet[[3]int]()
	cubeSet.AddMultiple(cubes)
	minBound := [3]int{minCube[0] - 1, minCube[1] - 1, minCube[2] - 1}
	maxBound := [3]int{maxCube[0] + 1, maxCube[1] + 1, maxCube[2] + 1}

	queue := collections.NewQueue[[3]int]()
	queue.Enqueue(minCube)
	flooded.Add(minCube)
	neighbourOffset := [][3]int{
		{-1, 0, 0},
		{+1, 0, 0},
		{0, -1, 0},
		{0, +1, 0},
		{0, 0, -1},
		{0, 0, +1},
	}
	for !queue.IsEmpty() {
		//if flooded.Size()%6 == 0 {
		//	printFlooding(cubeSet, flooded, minBound, maxBound)
		//	fmt.Print("")
		//}
		currentCube := queue.Dequeue()
		for i, offset := range neighbourOffset {
			movedDim := i / 2
			nextCube := [3]int{}
			for dim := 0; dim < 3; dim++ {
				nextCube[dim] = currentCube[dim] + offset[dim]
			}
			if cubeSet.Has(nextCube) {
				var touchedSide [4]int
				touchedSide[3] = movedDim
				for dim := 0; dim < 3; dim++ {
					if offset[dim] < 0 {
						touchedSide[dim] = currentCube[dim]
					} else {
						touchedSide[dim] = nextCube[dim]
					}
				}
				touchedSides.Add(touchedSide)
			} else if !flooded.Has(nextCube) {
				outOfBounds := false
				for dim := 0; dim < 3; dim++ {
					if nextCube[dim] < minBound[dim] || nextCube[dim] > maxBound[dim] {
						outOfBounds = true
						break
					}
				}
				if !outOfBounds {
					queue.Enqueue(nextCube)
					flooded.Add(nextCube)
				}
			}
		}
	}
	return touchedSides
}

func AoC18() {
	year := 2022
	day := 18
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	cubes := io.ReadInputFromRegexPerLineInt("(\\d+),(\\d+),(\\d+)", 18, 2022)
	fmt.Println(cubes)
	fmt.Println("Part 1:")
	surfaceSides := getCubeSurface(cubes)
	totalSeenSurfaces := len(surfaceSides)
	fmt.Println(totalSeenSurfaces)
	requests.SubmitAnswer(day, year, totalSeenSurfaces, 1)

	fmt.Println("Part 2:")
	//getOuterSides(collections.Map(surfaceSides, func(side []int) [4]int {
	//	var array [4]int
	//	copy(array[:], side)
	//	return array
	//}), [4]int{0, 0, 0, 0})
	minMaxCube := collections.Reduce(cubes, func(acc [][3]int, cube []int) [][3]int {
		for i, dimValue := range cube {
			if dimValue < acc[0][i] || acc[0][i] == -1 {
				acc[0][i] = dimValue
			}
			if dimValue > acc[1][i] {
				acc[1][i] = dimValue
			}
		}
		return acc
	}, [][3]int{{-1, -1, -1}, {0, 0, 0}})
	fmt.Println(minMaxCube)
	outerSides := floodCubes(collections.Map(cubes, func(cube []int) [3]int {
		var array [3]int
		copy(array[:], cube)
		return array
	}), minMaxCube[0], minMaxCube[1])
	fmt.Println(outerSides)
	fmt.Println(outerSides.Size())
	surfaceSidesSet := collections.NewSet[[4]int]()
	surfaceSidesSet.AddMultiple(surfaceSides)
	diff := surfaceSidesSet.Difference(outerSides)
	diffSlice := diff.AsSlice()
	sort.Slice(diffSlice, func(i, j int) bool {
		if diffSlice[i][3] < diffSlice[j][3] {
			return true
		}
		if diffSlice[i][3] == diffSlice[j][3] {
			for dim := 0; dim < 3; dim++ {
				if diffSlice[i][dim] < diffSlice[j][dim] {
					return true
				}
				if diffSlice[i][dim] > diffSlice[j][dim] {
					return false
				}
			}
		}
		return false
	})
	//fmt.Println(diffSlice)
	requests.SubmitAnswer(day, year, outerSides.Size(), 2)
}
