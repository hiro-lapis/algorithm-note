package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc283/tasks/abc283_a
 */
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	Power(a, b)
}

func Power(a int, b int) {
	result := a
	for i := 1; i < b; i++ {
		result *= a
		// fmt.Printf("%d is %d\n", i, result)
	}
	fmt.Println(result)
}
