package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc061/tasks/abc061_bs
 * solve time: 10m
 */
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	l := make([]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		l[(a-1)]++
		l[(b-1)]++
	}
	for _, v := range l {
		fmt.Println(v)
	}
}
