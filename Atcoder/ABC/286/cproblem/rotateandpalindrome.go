package main

import (
	"fmt"
	"math"
)

/*
 * https://atcoder.jp/contests/abc286/tasks/abc286_c
 */
func main() {
	//test := "test"
	//fmt.Printf("test type is %t", test[5])
	//return
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var s string
	fmt.Scan(&s)
	fmt.Println(s)

	// 回答値が膨大になった時のことを考えてcast
	ans := math.MaxInt64
	// 文字数分だけループ
	for i := 0; i < n; i++ {
		// l = 頭、r = 右から i 番目
		l := i
		r := i + n - 1
		sum := a * i
		fmt.Printf("l is %d : r is %d \n", l, r)
		for l < r {
			// ループ
			fmt.Printf("  l[%d] is %v : r[%d] is %v \n", l%n, s[l%n], r%n, s[r%n])
			if s[l%n] != s[r%n] {
				sum += b
			}
			l++
			r--
		}
		// 0~N-1 回のA操作 を行ってからB操作をした時のtotal costを比較
		// cost が低い方を回答値として更新
		fmt.Printf("%d 回A 操作をした時の cost %d \n", i, sum)
		ans = min(ans, sum)
	}
	fmt.Println(ans)
}

func Replace(s string, from int, to int) (result string) {
	fv := s[from : from+1]
	tv := s[to : to+1]
	for i := 0; i < len(s); i++ {
		if i+1 == from {
			result += tv
		} else if i+1 == to {
			result += fv
		} else {
			result = s[i : i+1]
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
