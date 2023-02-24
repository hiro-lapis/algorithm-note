package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc074/tasks/abc074_b
 */
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var sum int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if (x * 2) > moveB(x, k) {
			sum += moveB(x, k)
		} else {
			sum += x * 2
		}
	}
	fmt.Println(sum)
}

func moveB(x int, k int) int {
	return (k - x) * 2
}
