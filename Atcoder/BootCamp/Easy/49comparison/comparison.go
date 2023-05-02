package main

import (
	"fmt"
)

/*
 * https://atcoder.jp/contests/abc059/tasks/abc059_b
 */
func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if len(a) < len(b) {
		fmt.Println("LESS")
		return
	}
	if len(a) > len(b) {
		fmt.Println("GREATER")
		return
	}

	eq := true
	for i := 0; i < len(a); i++ {
		aNum := a[i]
		bNum := b[i]
		if aNum > bNum {
			fmt.Println("GREATER")
			eq = false
			break
		}
		if aNum < bNum {
			fmt.Println("LESS")
			eq = false
			break
		}
	}
	if eq {
		fmt.Println("EQUAL")
	}
}
