package main

import "fmt"

// https://atcoder.jp/contests/sumitrust2019/tasks/sumitb2019_c

func main() {
	// 100 to 105 x 1000,000
	var x int // 1~100,000
	fmt.Scan(&x)
	dp := make([]bool, x+1)
	dp[0] = true
	for i := 0; i < 6; i++ {
		for j := 0; j+i+100 <= x; j++ {
			if dp[j] {
				dp[j+i+100] = true
			}
		}
	}

	if dp[x] {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
