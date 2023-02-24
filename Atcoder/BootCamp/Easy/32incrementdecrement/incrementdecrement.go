package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc052/submissions/me
 * 回答時間:6m
 */
func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	max := 0
	current := 0
	for i := 0; i < n; i++ {
		char := s[i : i+1]
		if char == "I" {
			current++
			if current > max {
				max = current
			}
			continue
		}
		if char == "D" {
			current--
		}
	}
	fmt.Println(max)
}
