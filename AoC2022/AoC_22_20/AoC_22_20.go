package AoC_22_20

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/math"
	"AoC/utils/requests"
	"AoC/utils/types"
	"fmt"
	"strconv"
	"strings"
)

type ConnectedNumber struct {
	number int
	prev   *ConnectedNumber
	next   *ConnectedNumber
}

func (n *ConnectedNumber) toString() string {
	strs := collections.Map(n.toInts(), func(i int) string {
		return strconv.Itoa(i)
	})
	return strings.Join(strs, ", ")
}

func (n *ConnectedNumber) toInts() []int {
	start := n
	var ints []int
	ints = append(ints, start.number)
	//fmt.Println("prev", start.prev.number)
	current := start.next
	for current != start {
		//fmt.Println("prev", current.prev.number)
		ints = append(ints, current.number)
		current = current.next
	}
	return ints
}

func mixNumbers(first *ConnectedNumber, count int) *ConnectedNumber {
	fmt.Println("Start")
	fmt.Println(first.toString())
	numbers := first.toInts()
	originalNext := make(map[*ConnectedNumber]*ConnectedNumber)
	current := first.next
	originalNext[first] = first.next
	total := 1
	for current != first {
		originalNext[current] = current.next
		current = current.next
		total += 1
	}
	fmt.Println("total", total)

	var zero *ConnectedNumber
	for i := 0; i < count; i++ {
		current = first
		for range numbers {
			//fmt.Println(current.number)
			if current.number == 0 {
				zero = current
			}
			steps := current.number % (total - 1)
			if steps == 0 {
				current = originalNext[current]
				continue
			}

			prev := current.prev
			next := current.next
			prev.next = next
			next.prev = prev

			var newPrev *ConnectedNumber
			if steps > 0 {
				newPrev = current
			} else {
				newPrev = current.prev
			}
			for s := 0; s < math.Abs(steps); s++ {
				if steps > 0 {
					newPrev = newPrev.next
				} else {
					newPrev = newPrev.prev
				}
			}

			newNext := newPrev.next
			current.prev = newPrev
			current.next = newNext
			newNext.prev = current
			newPrev.next = current

			current = originalNext[current]
		}
		fmt.Println(i+1, zero.toString())
	}
	return zero
}

func numbersToConnected(numbers []int) *ConnectedNumber {
	first := &ConnectedNumber{
		number: numbers[0],
		prev:   nil,
		next:   nil,
	}
	prev := first
	for _, number := range numbers[1:] {
		currentNumber := &ConnectedNumber{
			number: number,
			prev:   prev,
			next:   nil,
		}
		prev.next = currentNumber
		prev = currentNumber
	}
	prev.next = first
	first.prev = prev
	return first
}

func keyFromMixed(mixed []int, positions []int) int {
	key := 0
	for _, position := range positions {
		key += mixed[position%len(mixed)]
	}
	return key
}

func AoC20() {
	year := 2022
	day := 20
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	lines := io.ReadInputLines(20, 2022)
	numbers := types.ToIntSlice(lines)

	set := collections.NewSet[int]()
	set.AddMultiple(numbers)
	fmt.Println(len(numbers), set.Size(), collections.MaxNumber(numbers))
	fmt.Println("Part 1:")
	first := numbersToConnected(numbers)
	mixed := mixNumbers(first, 1)
	fmt.Println(mixed.toString())
	key := keyFromMixed(mixed.toInts(), []int{1000, 2000, 3000})
	fmt.Println(key)
	requests.SubmitAnswer(day, year, key, 1)

	fmt.Println("Part 2:")
	decryptionKey := 811589153
	//bigNumbers := ndarray.New1D(numbers).MulScalar(decryptionKey).GetSlice(math.Range(len(numbers)))
	bigNumbers := collections.Map(numbers, func(n int) int {
		return n * decryptionKey
	})
	first = numbersToConnected(bigNumbers)
	mixed = mixNumbers(first, 10)
	key = keyFromMixed(mixed.toInts(), []int{1000, 2000, 3000})
	fmt.Println(key)
	requests.SubmitAnswer(day, year, key, 2)
}
