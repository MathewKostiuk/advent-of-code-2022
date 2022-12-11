package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	SPACES_PER_COLUMN           = 4
	NEWLINES_UNTIL_INSTRUCTIONS = 2
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	initialStack, instructions := parseInput(data)

	counter := initStacks(initialStack)
	result := processInstructions(instructions, counter)
	for key := range result {
		fmt.Println(key, string(result[key][len(result[key])-1]))
	}
}

func parseInput(data []byte) ([]byte, []byte) {
	initialStack := make([]byte, len(data))
	instructions := make([]byte, len(data))
	var spaceCounter, newLineCounter int
	var foundStack, foundInstructions bool
	for i, b := range data {
		if foundInstructions {
			break
		}
		if spaceCounter == 1 && newLineCounter == 1 && !foundStack {
			initialStack = data[0 : i-2]
			foundStack = true
		}
		switch b {
		case 109:
			instructions = data[i:]
			foundInstructions = true
		case 32:
			spaceCounter++
			continue
		case 10:
			newLineCounter++
			continue
		default:
			spaceCounter = 0
			newLineCounter = 0
		}
	}
	return initialStack, instructions
}

func initStacks(data []byte) map[int][]byte {
	var spaceCounter int
	col := 1
	counter := make(map[int][]byte)

	for _, b := range data {
		switch b {
		case 91:
			if spaceCounter == 1 {
				col++
				spaceCounter = 0
			}
			continue
		case 93:
			continue
		case 10:
			col = 1
			spaceCounter = 0
			continue
		case 32:
			spaceCounter++
			if spaceCounter == 4 {
				col++
				spaceCounter = 0
			}
			continue
		default:
			counter[col] = append(counter[col], b)
		}
	}

	for key := range counter {
		reverseByteSlice(counter[key])
	}
	return counter
}

func processInstructions(instructions []byte, counter map[int][]byte) map[int][]byte {
	split := strings.Split(string(instructions), "\n")
	for _, instruct := range split {
		if len(instruct) == 0 {
			continue
		}
		moveInt, fromInt, toInt := getInstructionDetails(instruct)

		toMove := counter[fromInt][len(counter[fromInt])-moveInt:]
		reverseByteSlice(toMove)
		counter[fromInt] = counter[fromInt][:len(counter[fromInt])-moveInt]
		counter[toInt] = append(counter[toInt], toMove...)
	}
	return counter
}

func getInstructionDetails(instruct string) (moveInt, fromInt, toInt int) {
	var move, from, to []byte
	moveStart := strings.Index((instruct), "move ")
	fromStart := strings.Index((instruct), "from ")
	toStart := strings.Index((instruct), "to ")

	for i := moveStart + len("move "); i < len(instruct); i++ {
		if instruct[i] == 32 {
			break
		}
		move = append(move, []byte(instruct)[i])
	}
	for i := fromStart + len("from "); i < len(instruct); i++ {
		if instruct[i] == 32 {
			break
		}
		from = append(from, []byte(instruct)[i])
	}
	for i := toStart + len("to "); i < len(instruct); i++ {
		if instruct[i] == 32 {
			break
		}
		to = append(to, []byte(instruct)[i])
	}
	moveInt, _ = strconv.Atoi(string(move))
	fromInt, _ = strconv.Atoi(string(from))
	toInt, _ = strconv.Atoi(string(to))
	return moveInt, fromInt, toInt
}

func reverseByteSlice(slice []byte) []byte {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
