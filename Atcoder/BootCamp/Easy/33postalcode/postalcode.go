package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var s string
	fmt.Scan(&s)

	front := s[:a]
	separater := s[a : a+1]
	back := s[a+1:]

	r := false
	// 全ての条件が 揃っている時
	r = !strings.Contains(front, "-") && !strings.Contains(back, "-") && separater == "-"
	if r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
