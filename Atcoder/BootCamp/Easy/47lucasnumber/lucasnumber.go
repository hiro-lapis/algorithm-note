package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(1)
		return
	}
	if n == 2 {
		fmt.Println(3)
		return
	}

	fmt.Println(Lucas(3, 1, n, 3))
}

func Lucas(high, row, mc, c int) int {
	if c == mc {
		return high + row
	}
	return Lucas(high+row, high, mc, c+1)
}
