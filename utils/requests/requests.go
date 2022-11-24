package requests

import (
	"AoC/secrets"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	InputUrl = "https://adventofcode.com/%v/day/%v/input"
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

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
