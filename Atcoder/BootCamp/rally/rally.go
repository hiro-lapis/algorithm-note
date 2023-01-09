package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc156/tasks/abc156_c
 */
func main() {
	var n, tmp int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		x[i] = tmp
	}

	// Average approach
	ans := AvgApproach(n, x)
	// brute force
	//ans := BruteForce(n, x)
	fmt.Println(ans)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func AvgApproach(n int, x []int) (ans int) {
	var sum, avg int
	for i := 0; i < n; i++ {
		sum += x[i]
	}
	// for decimal calculation, dividing with cast
	avg = int(math.Round(float64(sum) / float64(n)))
	for j := 0; j < n; j++ {
		ans += int(math.Pow(float64(x[j]-avg), 2.0))
	}
	return
}

func BruteForce(n int, x []int) (ans int) {
	ans = int(1e18)
	for i := 0; i < 100; i++ {
		var life int
		for j := 0; j < n; j++ {
			life += int(math.Pow(float64(x[j]-i), 2.0))
		}
		ans = min(ans, life)
	}
	return
}
