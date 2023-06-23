package main

import (
	"fmt"
	"math"
)

/*
 * https://atcoder.jp/contests/abc113/tasks/abc113_b
 * solve time: 8m
 */
func main() {
	var n, t, a int
	fmt.Scan(&n, &t, &a)
	min := math.MaxFloat64
	ans := 0
	for i := 1; i <= n; i++ {
		var tmp float64
		fmt.Scan(&tmp)
		avg := float64(t) - (tmp * 0.006)
		if min > abs(avg-float64(a)) {
			ans = i
			min = abs(avg - float64(a))
		}
	}
	fmt.Println(ans)
}

func abs(x float64) float64 {
	if x >= 0 {
		return x
	}
	return -x
}
