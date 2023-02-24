package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	min := 999
	for i := 0; i < len(s); i++ {
		if i > len(s)-3 {
			break
		}
		number, _ := strconv.Atoi(s[i : i+3])
		if min > abs(number-753) {
			min = abs(number - 753)
		}
	}
	fmt.Println(min)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
