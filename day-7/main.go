package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	COMMAND_BASE     = "$"
	CHANGE_DIRECTORY = "cd"
	LIST             = "ls"
	GO_BACK          = ".."
	DIRECTORY        = "dir"
	LIMIT            = 100000
	TOTAL_DISC_SPACE = 70000000
	REQUIRED_SPACE   = 30000000
)

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

type Directory struct {
	Name        string
	Size        int
	Directories []*Directory
	Files       []*File
	Parent      *Directory
}

func (d Directory) String() string {
	return fmt.Sprintf("Name: %s, Size: %d, Directories: %d, Parent: %s", d.Name, d.Size, len(d.Directories), d.Parent)
}

func (f File) String() string {
	return fmt.Sprintf("Name: %s, Size: %d, Parent: %s", f.Name, f.Size, f.Parent)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	result := []int{}

	root := buildTree(data)
	rootSize, sums := calculateDirTotals(&root, &result)
	one, two := partOne(*sums), partTwo(rootSize, *sums)
	fmt.Println(one, two)
}

func partOne(sums []int) int {
	var answer int
	for _, sum := range sums {
		if sum < LIMIT {
			answer += sum
		}
	}
	return answer
}

func partTwo(rootSize int, sums []int) int {
	var answer int
	unusedSpace := TOTAL_DISC_SPACE - rootSize
	neededSpace := REQUIRED_SPACE - unusedSpace
	sort.Ints(sums)

	for _, sum := range sums {
		if neededSpace-sum < 0 {
			answer = sum
			break
		}
	}
	return answer
}

func buildTree(data []byte) Directory {
	var currentDirectory *Directory
	root := Directory{}
	root.Name = "/"
	root.Directories = []*Directory{}
	root.Files = []*File{}
	currentDirectory = &root
	split := strings.Split(string(data), "\n")

	for _, char := range split {
		if len(char) == 0 {
			continue
		}

		str := strings.Split(string(char), " ")

		switch str[0] {
		case COMMAND_BASE:
			switch str[1] {
			case CHANGE_DIRECTORY:
				if str[2] == GO_BACK {
					currentDirectory = currentDirectory.Parent
				} else {
					if str[2] == "/" {
						currentDirectory = &root
					}
					for _, dir := range currentDirectory.Directories {
						if dir.Name == str[2] {
							currentDirectory = dir
						}
					}
				}
			case LIST:
			}
		case DIRECTORY:
			newDirectory := Directory{Name: str[1], Parent: currentDirectory, Directories: []*Directory{}, Files: []*File{}}
			currentDirectory.Directories = append(currentDirectory.Directories, &newDirectory)
		default:
			size, _ := strconv.Atoi(str[0])
			newFile := File{Size: size, Name: str[1], Parent: currentDirectory}
			currentDirectory.Size += size
			currentDirectory.Files = append(currentDirectory.Files, &newFile)
		}

	}
	return root
}

func calculateDirTotals(root *Directory, result *[]int) (int, *[]int) {
	for _, dir := range root.Directories {
		sum, result := calculateDirTotals(dir, result)
		*result = append(*result, sum)
		root.Size += sum
	}
	return root.Size, result
}
