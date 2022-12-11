package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var result int
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(string(data), "\n")

	for _, pair := range split {
		if len(pair) == 0 {
			continue
		}
		elfs := strings.Split(pair, ",")
		elf1, elf2 := strings.Split(elfs[0], "-"), strings.Split(elfs[1], "-")
		elf1Int, elf2Int := make([]int, len(elf1)), make([]int, len(elf2))

		for i, s := range elf1 {
			elf1Int[i], _ = strconv.Atoi(s)
		}

		for i, s := range elf2 {
			elf2Int[i], _ = strconv.Atoi(s)
		}

		elf1Start, elf1End, elf2Start, elf2End := elf1Int[0], elf1Int[1], elf2Int[0], elf2Int[1]

		if (elf2Start <= elf1Start && elf1Start <= elf2End) || elf1Start <= elf2Start && elf2Start <= elf1End {
			result++
		}
	}
	fmt.Println(result)
}
