package main

import (
	"fmt"
	"os"
)

const (
	START_OF_PACKET  = 4
	START_OF_MESSAGE = 14
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	partOne, partTwo := partOne(data), partTwo(data)
	fmt.Println(partOne, partTwo)
}

func partOne(input []byte) int {
	counter := make(map[byte]int)
	var left, right int

	for i := 0; i < START_OF_PACKET; i++ {
		counter[input[i]]++
	}

	for right = START_OF_PACKET; right < len(input); right++ {
		isUnique := checkUniqueness(counter)
		if isUnique {
			break
		}
		counter[input[left]]--
		left++
		counter[input[right]]++
	}
	return right
}

func partTwo(input []byte) int {
	counter := make(map[byte]int)
	var left, right int

	for i := 0; i < START_OF_MESSAGE; i++ {
		counter[input[i]]++
	}

	for right = START_OF_MESSAGE; right < len(input); right++ {
		isUnique := checkUniqueness(counter)
		if isUnique {
			break
		}
		counter[input[left]]--
		left++
		counter[input[right]]++
	}
	return right
}

func checkUniqueness(counter map[byte]int) bool {
	for _, value := range counter {
		if value > 1 {
			return false
		}
	}
	return true
}
