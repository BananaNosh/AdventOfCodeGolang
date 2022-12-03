package io

import (
	"AoC/utils/collections"
	date2 "AoC/utils/date"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	ExampleOsVariableName = "EXAMPLE_%d_%d"
)

func ReadInput(date ...int) string {
	day, year, _, _ := getDayAndYearString(date)
	shouldUseExample, ok := os.LookupEnv(fmt.Sprintf(ExampleOsVariableName, year, day))
	if ok && shouldUseExample == strconv.FormatBool(true) {
		return readExample(date...)
	}
	return readInput(date...)
}

func readInput(date ...int) string {
	day, year, yearString, dayString := getDayAndYearString(date)

	filePath := fmt.Sprintf("AoC%v/input/input_%v.txt", yearString, dayString)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(filePath, requests.LoadInput(day, year))
	}
	//else {
	//	fmt.Println("INFO: File already exists.. Will not create new one")
	//}

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(file), "\n")
}

func readExample(date ...int) string {
	day, year, yearString, dayString := getDayAndYearString(date)

	filePath := fmt.Sprintf("AoC%v/example/example_%v.txt", yearString, dayString)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(filePath, requests.LoadExample(day, year))
	}
	//else {
	//	fmt.Println("INFO: File already exists.. Will not create new one")
	//}

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(file), "\n")
}

func getDayAndYearString(date []int) (int, int, string, string) {
	var day, year int
	day = 0
	year = 0
	for _, d := range date {
		if d > 25 {
			year = d
		} else {
			day = d
		}
	}
	if day == 0 {
		day = date2.CurrentDay()
	}
	if year == 0 {
		year = date2.CurrentYear()
	}

	yearString := strconv.Itoa(year)
	dayString := strconv.Itoa(day)
	return day, year, yearString, dayString
}

func ReadAndSplitInput(delimiter string, date ...int) []string {
	input := ReadInput(date...)
	return strings.Split(input, delimiter)
}

func ReadInputLines(date ...int) []string {
	return ReadAndSplitInput("\n", date...)
}

func ReadInputFromRegex(regex string, date ...int) [][]string {
	lines := ReadInputLines(date...)
	pattern, err := regexp.Compile(regex)
	if err != nil {
		panic(err)
	}
	var result [][]string
	for _, line := range lines {
		line_matches := pattern.FindAllStringSubmatch(line, -1)
		submatches := collections.Reduce(line_matches, func(acc []string, match []string) []string {
			if len(match) > 1 {
				return append(acc, match[1:]...)
			}
			return append(acc, match...)
		}, []string{})
		fmt.Println(line_matches)

		var parsed []string
		for _, m := range submatches {
			parsed = append(parsed, m)
		}
		result = append(result, parsed)
	}
	return result
}

func ReadInputFromRegexInt(regex string, date ...int) [][]int {
	resultStrings := ReadInputFromRegex(regex, date...)
	var result [][]int
	for _, line := range resultStrings {
		result = append(result, types.ToIntSlice(line))
	}
	return result
}

func createFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("INFO: File successfully created")
	}
}
