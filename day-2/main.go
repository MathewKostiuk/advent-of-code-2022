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
		fmt.Println(string(arr[i][0]), string(arr[i][2]))
		opp := string(arr[i][0])
		me := string(arr[i][2])
		switch opp {
		case "A":
			switch me {
			case "X":
				ms += DRAW + ROCK
				os += DRAW + ROCK
			case "Y":
				ms += WIN + PAPER
				os += LOSE + ROCK
			case "Z":
				ms += LOSE + SCISSORS
				os += WIN + ROCK
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
				ms += WIN + ROCK
				os += LOSE + SCISSORS
			case "Y":
				ms += LOSE + PAPER
				os += WIN + SCISSORS
			case "Z":
				ms += DRAW + SCISSORS
				os += DRAW + SCISSORS
			}
		}
	}
	fmt.Println(ms, os)
}
