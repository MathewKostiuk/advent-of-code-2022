package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	WIN      = 6
	DRAW     = 3
	LOSE     = 0
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

func main() {
	var ms, os int

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Split(string(data), "\n")

	for i := 0; i < len(arr)-1; i++ {
		opp := string(arr[i][0])
		me := string(arr[i][2])
		switch opp {
		case "A":
			switch me {
			case "X":
				ms += LOSE + SCISSORS
				os += WIN + ROCK
			case "Y":
				ms += DRAW + ROCK
				os += DRAW + ROCK
			case "Z":
				ms += WIN + PAPER
				os += LOSE + ROCK
			}
		case "B":
			switch me {
			case "X":
				ms += LOSE + ROCK
				os += WIN + PAPER
			case "Y":
				ms += DRAW + PAPER
				os += DRAW + PAPER
			case "Z":
				ms += WIN + SCISSORS
				os += LOSE + PAPER
			}
		case "C":
			switch me {
			case "X":
				ms += LOSE + PAPER
				os += WIN + SCISSORS
			case "Y":
				ms += DRAW + SCISSORS
				os += DRAW + SCISSORS
			case "Z":
				ms += WIN + ROCK
				os += LOSE + SCISSORS
			}
		}
	}
	fmt.Println(ms, os)
}
