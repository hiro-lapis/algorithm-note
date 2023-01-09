package main

import "fmt"

/**
 * https://atcoder.jp/contests/code-festival-2016-qualb/tasks/codefestival_2016_qualB_b
 */
func main() {
	var n, a, b int
	var s string
	fmt.Scan(&n, &a, &b)
	fmt.Scan(&s)
	Solve(n, a, b, s)
}

func Solve(n int, a int, b int, s string) {
	limit := a + b
	bn := 1
	count := 0
	for i := 0; i < len(s); i++ {
		student := s[i : i+1]
		switch student {
		case "c":
			fmt.Println("No")
		case "a":
			if count < limit {
				fmt.Println("Yes")
				count++
			} else {
				fmt.Println("No")
			}
		case "b":
			if count < limit && bn <= b {
				fmt.Println("Yes")
				bn++
				count++
			} else {
				fmt.Println("No")
			}
		}
	}
}
