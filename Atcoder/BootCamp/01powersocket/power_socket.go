package main

import (
	"fmt"
)

/**
 * https://atcoder.jp/contests/abc139/tasks/abc139_b
 */
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := 1
	answer := 0
	for c < b {
		answer++
		c--
		c += a
	}
	fmt.Println(answer)
}
