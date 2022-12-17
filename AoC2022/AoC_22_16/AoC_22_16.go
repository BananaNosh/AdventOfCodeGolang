package AoC_22_16

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/search"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Valve struct {
	name       string
	rate       int
	neighbours []*Valve
}

type State struct {
	opened        *collections.Set[*Valve]
	currentRate   int
	totalReleased int
	pos           *Valve
}

func createValveSystem(valveValues [][]string) *Valve {
	var start *Valve
	created := make(map[string]*Valve)
	for _, valveStrings := range valveValues {
		name := valveStrings[0]
		valve := &Valve{}
		if collections.HasKey(created, name) {
			valve = created[name]
		}
		valve.name = name
		rate, _ := strconv.Atoi(valveStrings[1])
		valve.rate = rate
		neighbours := strings.Split(valveStrings[2], ", ")
		valve.neighbours = make([]*Valve, len(neighbours))
		for i, neighbourName := range neighbours {
			if !collections.HasKey(created, neighbourName) {
				created[neighbourName] = &Valve{name: neighbourName}
			}
			valve.neighbours[i] = created[neighbourName]
		}
		created[name] = valve
		if start == nil {
			start = valve
		}
	}
	return start
}

func getNeighbourFunc() func(state State) []State {
	return func(state State) []State {
		neighbours := make([]State, 0)
		currentValve := state.pos
		if !state.opened.Has(currentValve) && currentValve.rate > 0 {
			nowOpened := state.opened.Copy()
			nowOpened.Add(currentValve)
			neighbours = append(neighbours, State{
				opened:        nowOpened,
				currentRate:   state.currentRate + currentValve.rate,
				totalReleased: state.totalReleased + state.currentRate,
				pos:           currentValve,
			})
		}
		for _, neighbour := range currentValve.neighbours {
			neighbours = append(neighbours, State{
				opened:        state.opened,
				currentRate:   state.currentRate,
				totalReleased: state.totalReleased + state.currentRate,
				pos:           neighbour,
			})
		}
		return neighbours
	}
}

func getNeighbourWithDistFunc() func(state State) []search.Neighbour[State] {
	return func(state State) []search.Neighbour[State] {
		neighbours := getNeighbourFunc()(state)
		return collections.Map(neighbours, func(heigbour State) search.Neighbour[State] {
			return search.Neighbour[State]{
				Value:    heigbour,
				Distance: 1,
			}
		})
	}
}

func AoC16() {
	year := 2022
	day := 16
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	valveValues := io.ReadInputFromRegexPerLine("Valve (\\w+) (?:\\w+ )+rate=(\\d+); (?:\\w+ )+valves? ([\\w, ]+)", 16, 2022)
	fmt.Println(valveValues)
	firstValve := createValveSystem(valveValues)
	fmt.Println(firstValve)

	fmt.Println("Part 1:")
	startState := State{
		opened:        collections.NewSet[*Valve](),
		currentRate:   0,
		totalReleased: 0,
		pos:           firstValve,
	}
	maxDepth := 30
	heuristic := func(state State) int {
		return 0 //-state.totalReleased
	}
	identify := func(state State) string {
		setAsStrings := collections.Map(state.opened.AsSlice(), func(v *Valve) string {
			return v.name
		})
		sort.Strings(setAsStrings)
		return state.pos.name + ":" + strings.Join(setAsStrings, "-") + ":" + strconv.Itoa(state.totalReleased)
	}
	path := search.FindBestWithIdentify(startState, getNeighbourWithDistFunc(), func(state State) int { return state.totalReleased }, maxDepth+1, identify)
	for _, state := range path {
		fmt.Print(identify(state) + ", ")
	}
	fmt.Println()
	totalReleased := collections.Last(path).totalReleased
	fmt.Println(totalReleased)

	// requests.SubmitAnswer(day, year, 0, 1)

	fmt.Println("Part 2:")
	// requests.SubmitAnswer(day, year, 0, 2)
}
