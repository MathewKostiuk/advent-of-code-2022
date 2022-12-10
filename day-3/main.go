package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

const (
	LOWERCASE_MODIFIER = 96
	LOWERCASE_START    = 97
	LOWERCASE_END      = 122
	UPPERCASE_MODIFiER = 38
	UPPERCASE_START    = 65
	UPPERCASE_END      = 90
)

func main() {
	var sum int
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	split := bytes.Split(data, []byte("\n"))

	for i, j, k := 0, 1, 2; k < len(split); i, j, k = i+3, j+3, k+3 {
		str1 := split[i]
		str2 := split[j]
		str3 := split[k]
		map1 := make(map[byte]int)
		map2 := make(map[byte]int)
		map3 := make(map[byte]int)
		var longest map[byte]int

		for _, b := range str1 {
			map1[b]++
		}

		for _, b := range str2 {
			map2[b]++
		}

		for _, b := range str3 {
			map3[b]++
		}

		if len(map1) < len(map2) {
			longest = map2
		}

		if len(longest) < len(map3) {
			longest = map3
		}

		for key := range longest {
			_, ok1 := map1[key]
			_, ok2 := map2[key]
			_, ok3 := map3[key]
			if ok1 && ok2 && ok3 {
				var result int
				key := int(key)
				if LOWERCASE_START <= key && key <= LOWERCASE_END {
					result = key - LOWERCASE_MODIFIER
				} else {
					result = key - UPPERCASE_MODIFiER
				}
				sum += result
			}
		}
	}
	fmt.Println(sum)
}
