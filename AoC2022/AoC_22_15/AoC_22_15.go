package AoC_22_15

import (
	_9 "AoC/AoC2022/AoC_22_9"
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"fmt"
	"sort"
)

type SensorWithBeacon struct {
	sensor _9.Position
	beacon _9.Position
}

func readSensorsAndBeacons(lines [][]int) []SensorWithBeacon {
	var result []SensorWithBeacon
	for _, lineNumbers := range lines {
		result = append(result, SensorWithBeacon{_9.Position{X: lineNumbers[0], Y: lineNumbers[1]}, _9.Position{X: lineNumbers[2], Y: lineNumbers[3]}})
	}
	return result
}

func coveredPositionsInLineCount(sensorsAndBeacons []SensorWithBeacon, line int) int {
	ranges, blockedPositions := coveredPositionsInLine(sensorsAndBeacons, line, false, nil, nil)
	if len(ranges) == 0 {
		return 0
	}
	total := 0
	for _, r := range ranges {
		total += r[1] - r[0] + 1
		total -= collections.Count(blockedPositions[line], func(xPos int) bool { return xPos >= r[0] && xPos <= r[1] })
	}
	return total
}

func coveredPositionsInLine(sensorsAndBeacons []SensorWithBeacon, line int, countSensorsAndBeacons bool, lowBound *int, upBound *int) ([][]int, map[int]*collections.Set[int]) {
	blockedPositions := make(map[int]*collections.Set[int])
	ranges := make([][]int, 0)
	for _, sensorAndBeacon := range sensorsAndBeacons {
		if !countSensorsAndBeacons {
			for _, pos := range []_9.Position{sensorAndBeacon.sensor, sensorAndBeacon.beacon} {
				yPos := pos.Y
				xPos := pos.X
				if !collections.HasKey(blockedPositions, yPos) {
					blockedPositions[yPos] = collections.NewSet[int]()
				}
				blockedPositions[yPos].Add(xPos)
			}
		}
		sensor, beacon := sensorAndBeacon.sensor, sensorAndBeacon.beacon
		distX := math.Abs(sensor.X - beacon.X)
		distY := math.Abs(sensor.Y - beacon.Y)
		totalDist := distX + distY
		distToLine := math.Abs(sensor.Y - line)
		remaining := totalDist - distToLine
		if remaining > 0 {
			min := sensor.X - remaining
			max := sensor.X + remaining
			if lowBound != nil && upBound != nil {
				min = math.Max(min, *lowBound)
				max = math.Min(max, *upBound)
			}
			ranges = append(ranges, []int{min, max})
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0] || (ranges[i][0] == ranges[j][0] && ranges[i][1] < ranges[j][1])
	})
	index := 0
	for index < len(ranges)-1 {
		current := ranges[index]
		next := ranges[index+1]
		if next[0] <= current[1] { // overlapping
			ranges[index] = []int{current[0], math.Max(current[1], next[1])}
			ranges = append(ranges[:index+1], ranges[index+2:]...)
		} else {
			index += 1
		}
	}
	return ranges, blockedPositions
}

func findSignalPos(sensorsAndBeacons []SensorWithBeacon, max int) _9.Position {
	var foundPos _9.Position
	min := 0
	for i := 0; i <= max; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		ranges, _ := coveredPositionsInLine(sensorsAndBeacons, i, true, &min, &max)
		if len(ranges) > 1 {
			foundPos = _9.Position{ranges[0][1] + 1, i}
			break
		}
	}
	return foundPos
}

func AoC15() {
	year := 2022
	day := 15
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	lineToInspect := 2000000
	max := 4000000
	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))
	//lineToInspect = 10
	//max = 20

	lines := io.ReadInputFromRegexPerLineInt("(?:\\w+ ){2}x=(-?\\d+), y=(-?\\d+): (?:\\w+ ){4}x=(-?\\d+), y=(-?\\d+)", 15, 2022)
	fmt.Println(lines)
	sensorsAndBeacons := readSensorsAndBeacons(lines)
	fmt.Println("Part 1:")
	//coveredPositions := findCoveredPositions(sensorsAndBeacons)
	//fmt.Println(coveredPositions)
	//for _, sensorsAndBeacon := range sensorsAndBeacons {
	//	for _, position := range []_9.Position{sensorsAndBeacon.beacon, sensorsAndBeacon.sensor} {
	//		if collections.HasKey(coveredPositions, position.Y) {
	//			coveredPositions[position.Y].Remove(position.X)
	//		}
	//	}
	//}
	//coveredCount := coveredPositions[lineToInspect].Size()
	coveredCount := coveredPositionsInLineCount(sensorsAndBeacons, lineToInspect)
	fmt.Println(coveredCount)
	requests.SubmitAnswer(day, year, coveredCount, 1)

	fmt.Println("Part 2:")
	signalPos := findSignalPos(sensorsAndBeacons, max)
	freq := signalPos.X*4000000 + signalPos.Y
	fmt.Println(signalPos, freq)
	requests.SubmitAnswer(day, year, freq, 2)
}
