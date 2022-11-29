package requests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AnswerState int

const (
	NotGiven AnswerState = iota // 0
	Correct              = iota // 1
	Wrong                = iota // 2
)

func getDataDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataDir := filepath.Join(dir, ".config", "aocd", "go")

	err = os.MkdirAll(dataDir, 0700)
	if err != nil {
		panic(err)
	}
	return dataDir
}

func SaveGivenAnswer(day int, year int, answer string, part int, correct bool) {
	fileName := fileNameForAnswer(day, year, part, correct)
	flags := os.O_CREATE | os.O_RDWR
	if !correct {
		flags |= os.O_APPEND
		answer += "\n"
	}
	file, err := os.OpenFile(fileName, flags, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(answer))
	if err != nil {
		panic(err)
	}
}

func fileNameForAnswer(day int, year int, part int, correct bool) string {
	timePart := fileTimePart(day, year, part)
	fileName := "answer" + timePart
	if !correct {
		fileName = "answer_wrong" + timePart
	}
	fileName = filepath.Join(getDataDir(), fileName+".txt")
	return fileName
}

func fileTimePart(day int, year int, part int) string {
	part_str := "a"
	if part == 2 {
		part_str = "b"
	}
	time_part := fmt.Sprintf("_%d_%d_%s", year, day, part_str)
	return time_part
}

func CheckForAnswerToSubmit(day int, year int, answer string, part int) AnswerState {
	fileName := fileNameForAnswer(day, year, part, true)
	readAnswer, err := os.ReadFile(fileName)
	if err == nil {
		if string(readAnswer) == answer {
			return Correct
		}
		return Wrong
	} else if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	fileName = fileNameForAnswer(day, year, part, false)
	readAnswer, err = os.ReadFile(fileName)
	if err == nil {
		answers := strings.Split(strings.TrimSpace(string(readAnswer)), "\n")
		for _, givenAnswer := range answers {
			if givenAnswer == answer {
				return Wrong
			}
		}
	} else if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	return NotGiven
}
