package main

import (
	"fmt"
)

/*
 * 10m
 * https://atcoder.jp/contests/abc050/tasks/abc050_b
 */
func main() {
	var n int
	fmt.Scan(&n)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		t[i] = tmp
	}

	var m int
	fmt.Scan(&m)
	l := make([]int, m)
	for i := 0; i < m; i++ {
		var tmpP, tmpX int
		fmt.Scan(&tmpP)
		fmt.Scan(&tmpX)

		sum := 0
		for i, v := range t {
			if i+1 == tmpP {
				sum += tmpX
			} else {
				sum += v
			}
		}
		l[i] = sum
	}
	for _, v := range l {
		fmt.Println(v)
	}
}

func SumSlice(l []int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}
