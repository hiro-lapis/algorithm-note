package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc121/tasks/abc121_b
 */
func main() {
	var h, w int
	fmt.Scan(&h, &w)
	num := 0
	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}
	if h%2 == 1 && w%2 == 1 {
		num = 1
	}
	fmt.Println(h*w/2 + num)
}
