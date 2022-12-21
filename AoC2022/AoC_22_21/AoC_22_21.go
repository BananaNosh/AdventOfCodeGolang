package AoC_22_21

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
	"strconv"
)

type Op int

const (
	NONE Op = iota
	ADD
	SUB
	MUL
	DIV
)

type Value struct {
	name      string
	value     int
	args      []*Value
	operation Op
}

func opFromString(str string) Op {
	switch str {
	case "+":
		return ADD
	case "-":
		return SUB
	case "*":
		return MUL
	case "/":
		return DIV
	default:
		panic("No such Op")
	}
}

func (op Op) inverseOp() Op {
	switch op {
	case NONE:
		return NONE
	case ADD:
		return SUB
	case SUB:
		return ADD
	case MUL:
		return DIV
	case DIV:
		return MUL
	default:
		panic("No such Op")
	}
}

func (val Value) compute() int {
	switch val.operation {
	case NONE:
		return val.value
	case ADD:
		return val.args[0].compute() + val.args[1].compute()
	case SUB:
		return val.args[0].compute() - val.args[1].compute()
	case MUL:
		return val.args[0].compute() * val.args[1].compute()
	case DIV:
		return val.args[0].compute() / val.args[1].compute()
	default:
		panic("Op not implemented")
	}
}

func parseMonkeys(monkeys [][]string) map[string]*Value {
	nameToOp := make(map[string]*Value)
	for _, monkey := range monkeys {
		args := collections.Map([]string{monkey[0], monkey[2], monkey[4]}, func(inpStr string) *Value {
			if len(inpStr) == 0 {
				return nil
			}
			if collections.HasKey(nameToOp, inpStr) {
				return nameToOp[inpStr]
			} else {
				newVal := &Value{name: inpStr}
				nameToOp[inpStr] = newVal
				return newVal
			}
		})
		if len(monkey[1]) > 0 {
			num, _ := strconv.Atoi(monkey[1])
			args[0].value = num
		} else {
			args[0].args = args[1:]
			args[0].operation = opFromString(monkey[3])
		}
	}
	return nameToOp
}

func getInputToMatch(operations map[string]*Value) int {
	root := operations["root"]
	for _, arg := range root.args {
		start := &Value{name: "start"}
		reverseToHumn(arg, start)
	}
}

func reverseToHumn(arg *Value, start *Value) *Value {
	if arg.name == "humn" {
		return arg
	}
	next := &Value{
		name:      "inv-" + arg.name,
		args:      []*Value{start, arg},
		operation: arg.operation.inverseOp(),
	}

}

func AoC21() {
	year := 2022
	day := 21
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputFromRegexPerLine("(\\w+): (?:(\\d+)|(\\w+) (.) (\\w+))", 21, 2022)
	fmt.Println(input)
	operationsDict := parseMonkeys(input)
	fmt.Println("Part 1:")
	root := operationsDict["root"]
	result := root.compute()
	fmt.Println(result)
	requests.SubmitAnswer(day, year, result, 1)

	fmt.Println("Part 2:")

	// requests.SubmitAnswer(day, year, 0, 2)
}
