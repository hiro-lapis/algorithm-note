package main

import "fmt"

// https: //atcoder.jp/contests/sumitrust2019/tasks/sumitb2019_c

func main() {
	// 100 to 105 x 1000,000
	var x int // 1~100,000
	fmt.Scan(&x)
	ans := false

	dp := make([][]int, 6)
	for i := 0; i < len(dp); i++ {
		for j := 0; j < 6; j++ {
			v := 100
			if i != 0 {
				v *= i
			}
		}
	}

	// 100
	// ok
	// 200
	// 10:ok
	// 17,18,10:ng
	// 18:ng
	// 10

}
