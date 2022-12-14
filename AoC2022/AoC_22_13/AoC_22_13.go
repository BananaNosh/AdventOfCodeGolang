package AoC_22_13

import (
	"AoC/utils/collections"
	"AoC/utils/date"
	"AoC/utils/io"
	"AoC/utils/requests"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ListTreeItem struct {
	children []*ListTreeItem
	value    *int
	parent   *ListTreeItem
}

//func lineToNumber(line string) float32 {
//	numbers := strings.Split(strings.Trim(line, "[]\n "), ",")
//}

type TreeWithIndex struct {
	index int
	tree  *ListTreeItem
}

func (t *ListTreeItem) toString() string {
	str := ""
	if t.value != nil {
		str += strconv.Itoa(*t.value)
	} else {
		str += "["
		var children []string
		for _, child := range t.children {
			children = append(children, child.toString())
		}
		str += strings.Join(children, ",")
		str += "]"
	}
	return str
}

func parseLine(line string) *ListTreeItem {
	current := new(ListTreeItem)
	root := current
	pattern := regexp.MustCompile("\\d+|\\[|]")
	for _, sign := range pattern.FindAllString(line, -1) {
		if sign == "[" {
			newChild := new(ListTreeItem)
			current.children = append(current.children, newChild)
			newChild.parent = current
			current = newChild
		} else if sign == "]" {
			current = current.parent
		} else {
			newChild := new(ListTreeItem)
			number, _ := strconv.Atoi(sign)
			newChild.value = &number
			newChild.parent = current
			current.children = append(current.children, newChild)
		}
	}
	// unpack from outermost brackets
	if len(root.children) != 1 {
		panic("Wrong structure")
	}
	return root.children[0]
}

func compareTrees(tree1 *ListTreeItem, tree2 *ListTreeItem) int {
	if tree1.value != nil && tree2.value != nil {
		return *tree1.value - *tree2.value
	}
	if tree1.value != nil {
		tempTree := new(ListTreeItem)
		tempTree.children = append(tempTree.children, tree1)
		return compareTrees(tempTree, tree2)
	}
	if tree2.value != nil {
		tempTree := new(ListTreeItem)
		tempTree.children = append(tempTree.children, tree2)
		return compareTrees(tree1, tempTree)
	}
	for i, subTree := range tree1.children {
		if i >= len(tree2.children) {
			return 1
		}
		otherSubTree := tree2.children[i]
		comparison := compareTrees(subTree, otherSubTree)
		if comparison < 0 {
			return -1
		}
		if comparison > 0 {
			return 1
		}
	}
	if len(tree1.children) < len(tree2.children) {
		return -1
	}
	return 0
}

func printPair(pair []*ListTreeItem) {
	fmt.Println(pair[0].toString())
	fmt.Println(pair[1].toString())
	fmt.Println()
}

func getDividerPackets(numbers []int) []*ListTreeItem {
	var packets []*ListTreeItem
	for _, number := range numbers {
		line := fmt.Sprintf("[[%d]]", number)
		packets = append(packets, parseLine(line))
	}
	return packets
}

func printTrees(trees []*ListTreeItem) {
	fmt.Println()
	for _, tree := range trees {
		println(tree.toString())
	}
}

func AoC13() {
	year := 2022
	day := 13
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	input := io.ReadInput(13, 2022)
	input = strings.ReplaceAll(input, "\n\n", "\n")
	lines := strings.Split(input, "\n")
	fmt.Println(lines)
	trees := collections.Map(lines, func(line string) *ListTreeItem {
		return parseLine(line)
	})
	fmt.Println("Part 1:")
	pairs := collections.ReduceWithIndex(trees, func(i int, result [][]*ListTreeItem, tree *ListTreeItem) [][]*ListTreeItem {
		if i%2 == 0 {
			result = append(result, []*ListTreeItem{tree})
		} else {
			lastPair := result[len(result)-1]
			result[len(result)-1] = append(lastPair, tree)
		}
		return result
	}, make([][]*ListTreeItem, 0))
	var correctIndices []int
	for i, pair := range pairs[0:] {
		printPair(pair)
		comparison := compareTrees(pair[0], pair[1])
		if comparison <= 0 {
			correctIndices = append(correctIndices, i+1)
		}
	}
	indexSum := collections.Sum(correctIndices)
	fmt.Println(indexSum)
	requests.SubmitAnswer(day, year, indexSum, 1)

	fmt.Println("Part 2:")
	dividerPackets := getDividerPackets([]int{2, 6})
	trees = append(trees, dividerPackets...)
	printTrees(trees)
	sort.Slice(trees, func(i, j int) bool {
		return compareTrees(trees[i], trees[j]) < 0
	})
	printTrees(trees)
	indicesAndTrees := collections.Filter(collections.MapWithIndex(trees, func(index int, tree *ListTreeItem) TreeWithIndex {
		return TreeWithIndex{index + 1, tree}
	}), func(treeWithIndex TreeWithIndex) bool {
		return collections.Contains(dividerPackets, treeWithIndex.tree)
	})
	indices := collections.Map(indicesAndTrees, func(treeWithIndex TreeWithIndex) int {
		return treeWithIndex.index
	})
	fmt.Println(indices)
	requests.SubmitAnswer(day, year, collections.Prod(indices), 2)
}
