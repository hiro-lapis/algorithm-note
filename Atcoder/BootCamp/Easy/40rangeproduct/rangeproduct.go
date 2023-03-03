package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// 0 を挟んだレンジ
	if a < 0 && b >= 0 {
		fmt.Println("Zero")
		return
	}

	// 0以上の値
	if a > 0 {
		fmt.Println("Positive")
		return
	}
	// ここまで来るなら、a,bともに負の値
	count := Abs(a - b)
	if count%2 == 1 {
		fmt.Println("Positive")
		return
	}
	fmt.Println("Negative")
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
