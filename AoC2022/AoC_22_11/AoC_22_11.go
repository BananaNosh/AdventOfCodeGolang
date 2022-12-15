package AoC_22_11

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	name           int
	items          *collections.Queue[int]
	operation      func(int, int) int
	divCond        int
	trueTo         int
	falseTo        int
	playedItems    int
	makeSmallerMod int
}

func createMonkey(name int, match []string) *Monkey {
	startingItems := types.ToIntSlice(strings.Split(match[0], ", "))
	items := collections.NewQueue[int]()
	items.EnqueueMultiple(startingItems...)
	pattern := regexp.MustCompile("new = old ([+*]) (\\d+|old)")
	if !pattern.MatchString(match[1]) {
		panic("wrong operation")
	}
	var operation func(int, int) int
	submatch := pattern.FindAllStringSubmatch(match[1], 1)[0]
	opString := submatch[1]
	arg := submatch[2]
	if arg == "old" {
		if opString == "*" {
			operation = func(i int, i2 int) int {
				return i * i2
			}
		} else {
			operation = func(i int, i2 int) int {
				return i + i2
			}
		}
	} else {
		intArg, _ := strconv.Atoi(arg)
		if opString == "*" {
			operation = func(i int, i2 int) int {
				return i * intArg
			}
		} else {
			operation = func(i int, i2 int) int {
				return i + intArg
			}
		}
	}
	intProperties := types.ToIntSlice(match[2:])
	return &Monkey{name: name, items: items, operation: operation, divCond: intProperties[0], trueTo: intProperties[1], falseTo: intProperties[2]}
}

func (m *Monkey) play(isRelieved bool) map[int]*collections.Queue[int] {
	outMonkeys := make(map[int]*collections.Queue[int])
	m.items.DequeAll(func(item int) {
		newItem := m.operation(item, item)
		if isRelieved {
			newItem /= 3
		}
		if m.makeSmallerMod == 0 {
			fmt.Println(m)
		}
		newItem = newItem % m.makeSmallerMod
		var toMonkey int
		if newItem%m.divCond == 0 {
			toMonkey = m.trueTo
		} else {
			toMonkey = m.falseTo
		}
		if outMonkeys[toMonkey] == nil {
			outMonkeys[toMonkey] = collections.NewQueue[int]()
		}
		if newItem < 0 {
			fmt.Println(newItem, item)
			os.Exit(1)
		}
		outMonkeys[toMonkey].Enqueue(newItem)
		m.playedItems += 1
	})
	return outMonkeys
}

func play(monkeys []*Monkey, rounds int, isRelieved bool) {
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			outItems := monkey.play(isRelieved)
			for otherMonkeyName, outQueue := range outItems {
				monkeys[otherMonkeyName].items.EnqueueQueue(outQueue)
			}
		}
	}
}

func getMonkeyBusiness(monkeys []*Monkey) int {
	activities := collections.Map(monkeys, func(m *Monkey) int {
		return m.playedItems
	})
	fmt.Println(activities)
	sort.Ints(activities)
	monkeyBusiness := collections.Prod(activities[len(activities)-2:])
	return monkeyBusiness
}

func createMonkeys(monkeyMatches [][]string) []*Monkey {
	divMult := 1
	monkeys := collections.MapWithIndex(monkeyMatches, func(i int, match []string) *Monkey {
		monkey := createMonkey(i, match)
		divMult *= monkey.divCond
		return monkey
	})
	fmt.Println(divMult)
	for _, monkey := range monkeys {
		monkey.makeSmallerMod = divMult
	}
	return monkeys
}

func AoC11() {
	year := 2022
	day := 11
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	reg := "Monkey \\d+:\n +Starting items: ((?:\\d+(?:, )?)+)\n +Operation: ([^\n]+)\n + Test: divisible by (\\d+)\n +If true: throw to monkey (\\d+)\n +If false: throw to monkey (\\d+)"
	monkeyMatches := io.ReadInputFromRegex(reg, 2022, 11)
	fmt.Println(monkeyMatches)
	monkeys := createMonkeys(monkeyMatches)
	fmt.Println("Part 1:")
	play(monkeys, 20, true)
	monkeyBusiness := getMonkeyBusiness(monkeys)
	fmt.Println(monkeyBusiness)
	requests.SubmitAnswer(day, year, monkeyBusiness, 1)
	monkeys = createMonkeys(monkeyMatches)
	play(monkeys, 10000, false)
	monkeyBusiness = getMonkeyBusiness(monkeys)
	fmt.Println(monkeyBusiness)
	fmt.Println("Part 2:")
	requests.SubmitAnswer(day, year, monkeyBusiness, 2)
}
