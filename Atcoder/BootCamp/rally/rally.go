package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc156/tasks/abc156_c
 */
func main() {
	// 1 <= N <= 100 なので、会議を開く座標P も100 まで

	// n = 人数
	var n, tmp int
	fmt.Scan(&n)
	// x = 各自の座標
	x := make([]int, n)
	// 人数分の位置情報を受け取る
	// TODO: 関数化する場合に、fmt を参照渡しした方がいい？
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		x[i] = tmp
	}

	// brute force
	ans := int(1e18)
	for i := 0; i < 100; i++ {
		var life int
		for j := 0; j < n; j++ {
			life += int(math.Pow(float64(x[j]-i), 2.0))
		}
		ans = min(ans, life)
	}
	fmt.Println(ans)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
