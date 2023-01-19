package main

import (
	"fmt"
)

/**
 * https://atcoder.jp/contests/abc160/tasks/abc160_c
 */
func main() {
	var k, n int
	fmt.Scan(&k, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a[i] = tmp
	}

	//ans := a[n-1] - a[0]
	//for i := 1; i < n; i++ {
	//	ans = Min(ans, a[i]+a[i-1])
	//}
	//fmt.Println(ans)
	max := k - a[n-1] + a[0]
	for i := 1; i < n; i++ {
		max = Max(max, a[i]-a[i-1])
	}
	fmt.Println(k - max)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
