package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc150/tasks/abc150_c
 */
func main() {
	var n int
	fmt.Scan(&n)
	p := make([]int, n)
	q := make([]int, n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		p[i] = tmp
	}
	for j := 0; j < n; j++ {
		var tmp int
		fmt.Scan(&tmp)
		q[j] = tmp
	}

	ans := math.MaxInt32
	pattern := GetPattern(n)
	pVal := GetOrder(p, pattern)
	qVal := GetOrder(q, pattern)
	total := math.Abs(float64(pVal - qVal))
	totalInt := int(total)
	if ans > totalInt {
		ans = totalInt
	}
	fmt.Println(ans)
}

func Pair(n int) (r int) {
	r = n
	for i := n; i > 1; i-- {
		r *= i
	}
	return
}

func GetPattern(n int) []int {
	// 桁数ごとの組み合わせ値slice
	pattern := make([]int, n)
	pattern[0] = 1
	pattern[1] = 1
	for i := 2; i < n; i++ {
		pattern[i] = pattern[i-1] * i
	}
	return pattern
}

func GetOrder(l []int, p []int) int {
	n := len(p)
	ans := 0
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		sum := 0
		// 先頭の桁から順位算出
		for j := l[i] - 1; j >= 0; j-- {
			if used[j] {
				continue
			}
			sum++
		}
		// i回目のループで使われた数値を使用済みフラグを更新
		used[l[i]-1] = true
		// pの値と掛け合わせて桁ごとの影響値を掛けて順位の値を更新
		/**
		 * p[n-1-i]
		 * 3桁数列の1桁目＝1: sum(1) * p[2](value:2) = 2
		 */
		ans += sum * p[n-1-i]
	}
	return ans
}
