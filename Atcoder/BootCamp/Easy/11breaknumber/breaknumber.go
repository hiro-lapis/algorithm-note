package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc068/tasks/abc068_b
 */
func main() {
	var n int
	fmt.Scan(&n)
	a := 1
	for i := 1; i < n; {
		i *= 2
		if i <= n {
			a = i
		}
	}
	fmt.Println(a)
	//var n, max int
	//a := 1
	//fmt.Scan(&n)
	//
	//for i := 1; i <= n; i++ {
	//	tmp := 0
	//	Divide(i, &tmp)
	//	if tmp >= max {
	//		max = tmp
	//		a = i
	//	}
	//}
	//fmt.Println(a)
}

func Divide(n int, tmp *int) int {
	if n%2 != 0 {
		return 0
	}
	if n == 2 {
		*tmp++
		return 1
	}
	*tmp++
	return Divide(n/2, tmp)
}
