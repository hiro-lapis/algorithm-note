package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc133/tasks/abc133_b
 */
func main() {
	var n, d int
	fmt.Scan(&n, &d)

	l := make([][]int, n)
	for i := 0; i < n; i++ {
		ll := make([]int, d)
		for j := 0; j < d; j++ {
			var tmp int
			fmt.Scan(&tmp)
			ll[j] = tmp
		}
		l[i] = ll
	}

	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			y := i
			z := j
			sum := GetSum(l[y], l[z])

		}

		// 距離を求め、整数かどうかを判断
		if IsInt(math.Sqrt(float64(sum))) {
			count++
		}
	}
	fmt.Println(count)
}

// IsInt float64型の値が整数かどうかを判断
func IsInt(n float64) bool {
	return int(n) == int(math.Ceil(n))
}

// GetSum スライスの値の差分を合算した値を返却 (l, ll 要素数は同じであること)
func GetSum(l, ll []int) int {
	sum := 0
	for j := 0; j < len(l); j++ {
		// 差分値を二乗して座標値の合計値にたす
		diff := l[j] - ll[j]
		sum += diff * diff
	}
	return sum
}
