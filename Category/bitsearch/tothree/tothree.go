package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc182/tasks/abc182_c
 */
func main() {
	// 正の数
	var n string
	fmt.Scan(&n)
	// ケタ数
	d := len(n)

	ans := 100
	// 2進数をforで作る
	//for i := 1; i < 1<<d; i++ {
	//	nTmp := ""
	//	ansTmp := 0
	//	for shift := 0; shift < d; shift++ {
	//
	//		if 1&i>>shift == 1 {
	//			// uint8のbit型をintへキャスト
	//			nTmp = nTmp + n[int(shift)]
	//		} else {
	//			ansTmp++
	//		}
	//	}
	//	if nTmp == "" {
	//		continue
	//	}
	//}
	fmt.Println(n[0])
	fmt.Println(d)
	fmt.Println(ans)
}
