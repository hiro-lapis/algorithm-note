package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc103/tasks/abc103_b
 */
func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	ans := false
	for i := 0; i < len(s); i++ {
		if s == t {
			ans = true
			break
		}
		s = s[1:] + string(s[0])
	}
	if ans {
		fmt.Printf("Yes")
		return
	}
	fmt.Printf("No")
}
