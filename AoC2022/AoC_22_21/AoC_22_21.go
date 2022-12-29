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

var opStrings = []string{"+", "-", "*", "/"}

func opFromString(str string) Op {
	for i, opString := range opStrings {
		if str == opString {
			return Op(i + 1)
		}
	}
	panic("No such Op")
}

func (op Op) toString() string {
	return opStrings[int(op)-1]
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

func (op Op) inverseOpForArgIndex(argIndex int) Op {
	if argIndex == 1 && (op == DIV || op == SUB) {
		return op
	} else {
		return op.inverseOp()
	}
}

func (val *Value) compute() int {
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

//func (val *Value) toString() string {
//	if val.operation == NONE {
//		return val.name
//	}
//
//}

func (val *Value) hasHumnInput() bool {
	if val.name == "humn" {
		return true
	}
	if val.operation == NONE {
		return false
	}
	return collections.Reduce(val.args, func(has bool, arg *Value) bool {
		return arg.hasHumnInput() || has
	}, false)
}

func (val *Value) backward(start *Value) *Value {
	if val.operation == NONE {
		return start
	}
	humnIndex := -1
	for i, arg := range val.args {
		if arg.hasHumnInput() {
			humnIndex = i
			break
		}
	}
	if humnIndex == -1 {
		panic("Humn should be in this path")
	}
	op := val.operation.inverseOpForArgIndex(humnIndex)
	newArgs := make([]*Value, len(val.args))
	copy(newArgs, val.args)
	newArgs[humnIndex] = start
	if humnIndex == 1 && (val.operation == ADD || val.operation == MUL) {
		newArgs = collections.Reverse(newArgs)
	}
	backVal := &Value{
		name:      "inv-" + val.name,
		value:     0,
		args:      newArgs,
		operation: op,
	}
	return val.args[humnIndex].backward(backVal)

}

func getInputToMatch(operations map[string]*Value) int {
	root := operations["root"]
	humnIndex := 0
	for i, arg := range root.args {
		if arg.hasHumnInput() {
			humnIndex = i
			break
		}
	}
	fmt.Println("Should be ", root.args[1-humnIndex].compute())
	fmt.Println("Is ", root.args[humnIndex].compute())
	backward := root.args[humnIndex].backward(root.args[1-humnIndex])
	return backward.compute()
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

//
//func reverseToHumn(arg *Value, start *Value) *Value {
//	if arg.name == "humn" {
//		return arg
//	}
//	next := &Value{
//		name:      "inv-" + arg.name,
//		args:      []*Value{start, arg},
//		operation: arg.operation.inverseOp(),
//	}
//
//}

func AoC21() {
	year := 2022
	day := 21
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInputFromRegexPerLine("(\\w+): (?:(-?\\d+)|(\\w+) (.) (\\w+))", 21, 2022)
	fmt.Println(input)
	operationsDict := parseMonkeys(input)
	fmt.Println("Part 1:")
	root := operationsDict["root"]
	result := root.compute()
	fmt.Println(result)
	requests.SubmitAnswer(day, year, result, 1)

	fmt.Println("Part 2:")
	inputToMatch := getInputToMatch(operationsDict)
	fmt.Println("res ", inputToMatch)
	//fmt.Println(operationsDict["sjmn"].compute())
	//fmt.Println(94625185243550 - 236694194244295)
	requests.SubmitAnswer(day, year, inputToMatch, 2)
}
