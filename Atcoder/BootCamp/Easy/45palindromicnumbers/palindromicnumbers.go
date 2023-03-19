package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	ans := 0
	for i := a; i <= b; i++ {
		if checkIsKaibun(i) {
			ans++
		}
	}
	fmt.Println(ans)
}

func checkIsKaibun(numeric int) bool {
	text := strconv.Itoa(numeric)
	return (text[0] == text[4]) && text[1] == text[3]
}
