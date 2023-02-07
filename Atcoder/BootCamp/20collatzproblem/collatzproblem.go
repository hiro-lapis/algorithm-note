package main

import "fmt"

// https://atcoder.jp/contests/abc116/tasks/abc116_b
func main() {
	var s int
	fmt.Scan(&s)
	count := 1
	var encount bool

	ans := F(s, &count, &encount)
	fmt.Println(ans)
}

func F(n int, c *int, e *bool) int {
	if n == 4 || n == 2 || n == 1 {
		return *c + 3
	}
	*c++
	if n%2 == 0 {
		return F(n/2, c, e)
	} else {
		return F(3*n+1, c, e)
	}
}

//func F(n int, c *int, e *bool) int {
//	if n == 4 {
//		if *e {
//			return *c
//		}
//		*e = true
//	}
//	*c++
//	if n%2 == 0 {
//		return F(n/2, c, e)
//	} else {
//		return F(3*n+1, c, e)
//	}
//}
