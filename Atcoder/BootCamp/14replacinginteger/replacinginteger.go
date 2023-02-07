package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc161/tasks/abc161_c
 */
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(Solve(n, k))
}

func Solve(dividend, divisor int) int {
	// for test16, avoid use math.Abs()
	remainder := Abs(dividend % divisor)
	math.Abs(10.00)
	abNum := Abs(divisor - remainder)
	if abNum < remainder {
		return abNum
	}
	return remainder
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
