package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc290/tasks/abc290_b
 */
func main() {
	var n, k int
	var s string

	fmt.Scan(&n, &k)
	fmt.Scan(&s)

	limit := k
	ans := ""
	index := 1
	for _, v := range s {
		char := string(v)
		if limit == 0 {
			ans += "x"
		} else if char == "o" {
			ans += "o"
			limit--
		} else if char == "x" {
			ans += "x"
		}
		index++
	}
	fmt.Println(ans)
}
