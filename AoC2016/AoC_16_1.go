package AoC16

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	math2 "AoC/utils/math"
	"AoC/utils/requests"
	"fmt"
	"strconv"
)

func AoC1() {
	fmt.Println("On " + date.DateStringForDay(2016, 1) + ":")
	input := io.ReadAndSplitInput(", ", 1, 2016)
	fmt.Println("Part 1:")
	pos := findDest(input, false)
	fmt.Println(math2.Abs(pos.x) + math2.Abs(pos.y))
	fmt.Println("Part 2:")
	pos = findDest(input, true)
	sol2 := math2.Abs(pos.x) + math2.Abs(pos.y)
	fmt.Println(sol2)
	requests.SubmitAnswer(1, 2016, sol2, 2)
}

type Position struct {
	x int
	y int
}

func findDest(input []string, returnOnDouble bool) Position {
	x, y := 0, 0
	face := 0
	seen := collections.NewSet[Position]()
	for _, order := range input {
		count, err := strconv.Atoi(order[1:])
		if err != nil {
			panic(err)
		}
		if order[0] == 'L' {
			face = (face + 3) % 4
		} else {
			face = (face + 1) % 4
		}
		for i := 0; i < count; i++ {
			if face == 0 {
				y += 1
			} else if face == 1 {
				x += 1
			} else if face == 2 {
				y -= 1
			} else {
				x -= 1
			}
			if returnOnDouble {
				currentPos := Position{x, y}
				if ok := seen.Has(currentPos); ok {
					return currentPos
				}
				seen.Add(currentPos)
			}
		}
	}
	return Position{x, y}
}
