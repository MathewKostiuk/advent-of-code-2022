package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	sort.Slice(counter, func(i, j int) bool {
		return counter[j] < counter[i]
	})
	fmt.Println(counter[0] + counter[1] + counter[2])
}
