package main

import (
	"fmt"
	"sort"
)

// https://atcoder.jp/contests/agc012/tasks/agc012_a
func main() {
	var n int
	fmt.Scan(&n)
	m := n * 3
	a := make([]int, m)
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a[i] = tmp

	}
	// sort desc
	//sort.Slice(a, func(i, j int) bool {
	//	return a[i] > a[j]
	//})
	sort.Ints(a)
	sum := 0

	for i := 0; i < n; i++ {
		// asc sort した配列を奥の方から参照していく
		// ↓の-1部分は 配列のlength,配列のindexの差分調整のための1
		// ↓の+1部分は 2番目に大きい値を参照するための1.ループ毎に参照するindexを飛ばすため
		v := a[m-(2*i+1)-1]
		// ↓のようにしてもいいが、異なる意味をもつ値を別に扱うため↑のようにすることを推奨
		//v := a[m-(2*i)-2]
		sum += v
	}
	fmt.Println(sum)
}
