package main

import (
	"fmt"
	"math"
)

/**
 * https://atcoder.jp/contests/abc158/tasks/abc158_c
 * solve time: 48m
 */
func main() {
	const PRICE8 = 12.5
	const PRICE10 = 10
	var a, b int
	fmt.Scan(&a, &b)
	// 2 2 => 25
	// 2 * 12.5(¥13)  ¥25~37, 2 * 10 ¥20~29
	/*
	 *  8 10
	 * [100 112] (12.5 => 12 without tax, float is floored)
	 * [100 109]
	 */

	aRange := [2]int{int(math.Ceil(PRICE8 * float64(a))), int(math.Ceil(PRICE8*float64(a+1)) - 1)}
	bRange := [2]int{int(math.Ceil(PRICE10 * float64(b))), int(math.Ceil(PRICE10*float64(b+1) - 1))}

	//fmt.Println(aRange)
	//fmt.Println(bRange)
	ans := -1
	for i := aRange[0]; i <= aRange[1]; i++ {
		if ans != -1 {
			break
		}
		for j := bRange[0]; j <= bRange[1]; j++ {
			if i == j {
				ans = i
				break
			}
		}
	}
	fmt.Println(ans)
}
