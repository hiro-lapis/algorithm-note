package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc286/tasks/abc286_b
 */
func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	nExist := false
	var ans string
	for i := 0; i < n; i++ {
		if nExist && s[i:i+1] == "a" {
			ans += "y"
			nExist = false
		} else if s[i:i+1] == "n" {
			nExist = true
		} else {
			nExist = false
		}
		ans += s[i : i+1]
	}
	fmt.Println(ans)
}
