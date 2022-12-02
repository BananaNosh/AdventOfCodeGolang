package requests

import (
	"AoC/secrets"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	url2 "net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	BaseUrl     = "https://adventofcode.com/%v/day/%v"
	InputUrl    = BaseUrl + "/input"
	AnswerUrl   = BaseUrl + "/answer"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorOrange = "\033[34m"
	colorReset  = "\033[0m"
)

func LoadInput(day int, year int) string {
	fmt.Println(day, year)
	url := fmt.Sprintf(InputUrl, year, day)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", secrets.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func SubmitAnswer(day int, year int, answer int, part int) {
	SubmitStringAnswer(day, year, strconv.Itoa(answer), part)
}

func SubmitStringAnswer(day int, year int, answer string, part int) {
	answerState := CheckForAnswerToSubmit(day, year, answer, part)
	defer fmt.Print(colorReset + "\n")
	if answerState == Correct {
		fmt.Printf(colorGreen+"Answer %s already given and correct!\n", answer)
		return
	} else if answerState == Wrong {
		fmt.Printf(colorRed+"Answer %s already given and WRONG!\n", answer)
		return
	}
	isCorrect := SendStringAnswer(day, year, answer, part)
	SaveGivenAnswer(day, year, answer, part, isCorrect)
	if isCorrect {
		fmt.Println(colorGreen + "Puzzle solved!")
	} else {
		fmt.Printf(colorRed+"Wrong Answer %s!\n", answer)
	}
}

func SendAnswer(day int, year int, answer int, part int) {
	SendStringAnswer(day, year, strconv.Itoa(answer), part)
}

func SendStringAnswer(day int, year int, answer string, part int) bool {
	url := fmt.Sprintf(AnswerUrl, year, day)
	fmt.Println(url)
	data := url2.Values{
		"level":  {strconv.Itoa(part)},
		"answer": {answer},
	}
	fmt.Println(data)
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", secrets.Session)
	req.Header.Set("user-agent", "advent-of-code-go")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	message := doc.Find("article").Text()
	fmt.Println("Message: ", message)
	if strings.Contains(message, "That's the right answer") {
		return true
	}
	if strings.Contains(message, "That's not the right answer") {
		return false
	}
	if strings.Contains(message, "Did you already complete it") {
		println(colorYellow + "Already solved this level! Might be wrong or false.")
		return false
	}
	if strings.Contains(message, "You gave an answer too recently") {
		reg := regexp.MustCompile("((\\d+)m)? (\\d+)s")
		matches := reg.FindStringSubmatch(message)
		remainingMin := 0
		if len(matches[2]) > 0 {
			remainingMin, _ = strconv.Atoi(matches[2])
		}
		remainingSek, _ := strconv.Atoi(matches[3])
		remaining := 60*remainingMin + remainingSek
		if remaining > 300 {
			fmt.Println(colorRed + "Has to wait too long, not reattempting")
			return false
		}
		fmt.Printf(colorOrange+"Waiting %ds for submitting next answer:\n", remaining)
		time.Sleep(time.Duration(remaining) * time.Second)
		fmt.Println(colorYellow + "Resubmitting ...")
		return SendStringAnswer(day, year, answer, part)
	}
	panic("Some other problem:\n " + message)
}

//response = requests.post(
//url=url,
//cookies=self.user.auth,
//headers=USER_AGENT,
//data={"level": level, "answer": value},
//)
