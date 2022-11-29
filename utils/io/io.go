package io

import (
	date2 "AoC/utils/date"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadInput(date ...int) string {
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

	filePath := fmt.Sprintf("AoC%v/input/input_%v.txt", yearString, dayString)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(day, year, filePath)
	}
	//else {
	//	fmt.Println("INFO: File already exists.. Will not create new one")
	//}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(file), "\n")
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
		line_matches := pattern.FindAll([]byte(line), -1)

		var parsed []string
		for _, m := range line_matches {
			parsed = append(parsed, string(m))
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

func createFile(day int, year int, path string) {
	puzzleInput := requests.LoadInput(day, year)

	err := os.WriteFile(path, []byte(puzzleInput), 0644)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("INFO: File successfully created")
	}
}
