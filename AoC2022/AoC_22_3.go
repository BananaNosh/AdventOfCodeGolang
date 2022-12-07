package AoC22

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
)

func findWrongItemInRucksack(rucksack string) string {
	items1 := collections.NewSet[string]()
	items2 := collections.NewSet[string]()
	for i, item := range rucksack {
		if i < len(rucksack)/2 {
			items1.Add(string(item))
		} else {
			items2.Add(string(item))
		}
	}
	common := items1.Intersect(items2)
	fmt.Println(items1, items2, common)
	return common.GetRandom()
}

func valueForItem(item string) int {
	ascii := int(item[0])
	if ascii >= 97 {
		return ascii - 96
	}
	return ascii - 38
}

func findValueOfBatch(group []string) int {
	rucksacks := make([]collections.Set[string], len(group))
	for i, rucksack := range group {
		rucksacks[i] = collections.NewSet[string]()
		for _, item := range rucksack {
			rucksacks[i].Add(string(item))
		}
	}
	common := rucksacks[0]
	for _, rucksackSet := range rucksacks[1:] {
		common = common.Intersect(rucksackSet)
	}
	return valueForItem(common.GetRandom())
}

func AoC3() {
	year := 2022
	day := 3
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")
	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))
	lines := io.ReadInputLines(3, 2022)
	fmt.Println("Part 1:")
	wrongItems := collections.Map(lines, findWrongItemInRucksack)
	fmt.Println(wrongItems)
	total := collections.Reduce(wrongItems, func(acc int, item string) int {
		return valueForItem(item) + acc
	}, 0)
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 1)
	fmt.Println("Part 2:")
	groupSize := 3
	total = 0
	for i := 0; i < len(lines)-2; i += groupSize {
		group := lines[i : i+groupSize]
		value := findValueOfBatch(group)
		total += value
	}
	fmt.Println(total)
	requests.SubmitAnswer(day, year, total, 2)
}
