package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	// 全体範囲を扱う変数定義
	rangeMap := make([]int, n+1)

	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		l--
		/**
		 * index=0 を 1 として各範囲をマッピング
		 * ex.l=1 => range[0]++
		 * ex.r=3 => range[3]-- (4番目が各範囲から出るindex)
		 */
		rangeMap[l]++
		rangeMap[r]--
	}
	ans := 0
	for i := 0; i < n; i++ {
		// 累積和(ゲート通過できる範囲の重複数)計算
		rangeMap[i+1] += rangeMap[i]
		// 全てのゲートを通過できるカードの時は、カウントをincrement
		if m == rangeMap[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
