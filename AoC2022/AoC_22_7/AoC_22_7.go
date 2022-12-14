package AoC_22_7

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

type TreeItem interface {
	hasChildren() bool
	size() int
}

type Directory struct {
	TreeItem
	name             string
	parent           *Directory
	childDirectories map[string]*Directory
	childFiles       map[string]File
}

type File struct {
	TreeItem
	name     string
	parent   *Directory
	fileSize int
}

func (d Directory) allChildren() map[string]TreeItem {
	all := make(map[string]TreeItem, len(d.childDirectories)+len(d.childFiles))
	for name, directory := range d.childDirectories {
		all[name] = directory
	}
	for name, file := range d.childFiles {
		all[name] = file
	}
	return all
}

func (d Directory) hasChildren() bool {
	return len(d.allChildren()) > 0
}

func (d Directory) size() int {
	totalSize := 0
	for _, child := range d.allChildren() {
		totalSize += child.size()
	}
	return totalSize
}

func (d Directory) findDirectories(cond func(directory *Directory) bool) []*Directory {
	var found []*Directory
	if cond(&d) {
		found = append(found, &d)
	}
	for _, directory := range d.childDirectories {
		found = append(found, directory.findDirectories(cond)...)
	}
	return found
}

func (d Directory) toString() string {
	indent := "   "
	str := d.name
	for _, file := range d.childFiles {
		str += "\n" + indent + file.toString()
	}
	for _, directory := range d.childDirectories {
		for _, line := range strings.Split(directory.toString(), "\n") {
			str += "\n" + indent + line
		}
	}
	return str
}

func (f File) hasChildren() bool {
	return false
}

func (f File) size() int {
	return f.fileSize
}

func (f File) toString() string {
	return fmt.Sprintf("%s: %d", f.name, f.fileSize)
}

func makeDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, parent: parent, childDirectories: make(map[string]*Directory), childFiles: make(map[string]File)}
}

func parseFileTree(lines []string) *Directory {
	commandReg := regexp.MustCompile("\\$ (\\w+) ?([^\\n]+)?")
	root := makeDirectory("root", nil)
	var currentDir *Directory
	for _, line := range lines {
		submatches := commandReg.FindStringSubmatch(line)
		//fmt.Println(submatches)
		isCommand := len(submatches) > 0
		if isCommand {
			if submatches[1] == "cd" {
				argument := submatches[2]
				if argument == "/" {
					currentDir = root
				} else if argument == ".." {
					currentDir = currentDir.parent
					//if currentDir != root {
					//}
				} else {
					currentDir = currentDir.childDirectories[argument]
				}
			}
		} else { // assuming ls command was read before
			split := strings.Split(line, " ")
			name := split[1]
			if split[0] == "dir" {
				currentDir.childDirectories[name] = makeDirectory(name, currentDir)
			} else {
				size, _ := strconv.Atoi(split[0])
				currentDir.childFiles[name] = File{name: name, parent: currentDir, fileSize: size}
			}

		}
	}
	return root
}

func calculateSpaceToFree(d *Directory, neededSpace int) int {
	return d.size() - neededSpace
}

func AoC7() {
	year := 2022
	day := 7
	fmt.Println("On " + date.DateStringForDay(year, day) + ":")

	// setting EXAMPLE variable
	//_ = os.Setenv(fmt.Sprintf(io.ExampleOsVariableName, year, day), strconv.FormatBool(true))

	lines := io.ReadInputLines(7, 2022)
	fmt.Println(lines)
	fmt.Println("Part 1:")
	d := parseFileTree(lines)
	dirsWithMaxSize := d.findDirectories(func(d *Directory) bool {
		return d.size() < 100000
	})
	foundSizes := collections.Map(dirsWithMaxSize, func(d *Directory) int {
		return d.size()
	})
	fmt.Println(foundSizes)
	sum := collections.Sum(foundSizes)
	fmt.Println(sum)
	fmt.Println(d.toString())
	requests.SubmitAnswer(day, year, sum, 1)
	fmt.Println("Part 2:")
	spaceToFree := calculateSpaceToFree(d, 70000000-30000000)
	fmt.Println(spaceToFree)
	possibleDeletionDirectories := d.findDirectories(func(d *Directory) bool {
		return d.size() > spaceToFree
	})
	sizesOfDeletion := collections.Map(possibleDeletionDirectories, func(d *Directory) int {
		return d.size()
	})
	sort.Ints(sizesOfDeletion)
	fmt.Println(sizesOfDeletion)
	requests.SubmitAnswer(day, year, sizesOfDeletion[0], 2)
}
