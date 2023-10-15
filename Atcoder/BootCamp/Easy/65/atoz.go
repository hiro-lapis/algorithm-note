package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc053/tasks/abc053_b
 */
func main() {
	var s string
	fmt.Scan(&s)

	var begin, end int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			begin = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "Z" {
			end = i
			break
		}
	}
	fmt.Println(end - begin + 1)
}
