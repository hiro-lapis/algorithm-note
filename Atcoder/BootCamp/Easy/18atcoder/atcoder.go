package main

import (
	"fmt"
	"strings"
)

/**
 * https://atcoder.jp/contests/abc122/tasks/abc122_b
 */
func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)

	var max int
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			partial := s[i : n-j]
			count := Acgt(partial)
			if count > max {
				max = count
			}
		}
	}
	fmt.Println(max)
}

func Acgt(s string) (c int) {
	for _, char := range s {
		if strings.ContainsAny(string(char), "ACGT") {
			c++
		} else {
			c = 0
		}
	}
	return
}
