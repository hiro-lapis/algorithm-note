package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc058/tasks/abc058_b
 * 2023/7/7 10:06~10:20
 */
func main() {
	var o, e string
	fmt.Scan(&o)
	fmt.Scan(&e)
	l := len(o + e)

	var ans string
	var it int
	for i := 0; i < l; i++ {
		// odd
		if i%2 == 0 {
			if it < len(o) {
				ans += string(o[it])
			}
		} else {
			ans += string(e[it])
			it++
		}
	}
	fmt.Println(ans)
}
