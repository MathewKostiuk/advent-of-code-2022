package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data), "\n")
	var i int
	counter := make([]int64, len(arr))
	for _, val := range arr {
		if val == "" {
			i++
			continue
		}
		num, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		counter[i] += num
	}
	var largest int64
	for _, val := range counter {
		if largest < val {
			largest = val
		}
	}
	fmt.Println(largest)
}
