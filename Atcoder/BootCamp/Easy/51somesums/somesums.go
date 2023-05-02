package main

import (
	"fmt"
	"strconv"
)

/**
 * https://atcoder.jp/contests/abc083/tasks/abc083_b
 * solve time: 20m
 */
func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	// between 1~N of int, output i of sum of each digit is A<= i <= B
	ans := 0
	for i := 1; i <= n; i++ {
		tmp := strconv.Itoa(i)
		sum := getSum(tmp)
		if sum >= a && b >= sum {
			ans += i
		}
	}
	fmt.Println(ans)
}

// getSum receive numeric string, return each sum of each numeric
func getSum(s string) int {
	r := 0
	for _, i := range s {
		numericS := string(i)
		num, _ := strconv.Atoi(numericS)
		r += num
	}
	return r
}
