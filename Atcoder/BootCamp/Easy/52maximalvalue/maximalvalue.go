package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc140/tasks/abc140_c
 * solve time: 40m
 */
func main() {
	var n int
	fmt.Scan(&n)

	ans := 0
	before := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			ans += before
			break
		}
		var bi int
		fmt.Scan(&bi)
		if i == 0 {
			ans += bi
			before = bi
			continue
		}
		if bi > before {
			ans += before
			before = bi
			continue
		}
		ans += bi
		before = bi
	}
	fmt.Println(ans)
}
