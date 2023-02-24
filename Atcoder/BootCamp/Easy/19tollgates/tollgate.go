package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc094/tasks/abc094_b
 */
func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)
	//list := make([]int, m)

	var above, before int
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp > x {
			above++
		} else {
			before++
		}
	}
	fmt.Println(Min(before, above))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
