package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc136/tasks/abc136_d
 */
func main() {
	var s string
	fmt.Scan(&s)
	r := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		c := i
		//fmt.Printf("i is %d \n", c)
		for j := 0; j < 10^100; j++ {
			a := s[c : c+1]
			var b string
			//fmt.Printf("current %s \n", a)
			if a == "L" {
				b = s[c-1 : c]
			} else {
				b = s[c+1 : c+2]
			}
			if a != b && j%2 == 0 {
				r[c]++
				break
			}
			if a == "L" {
				c--
			} else {
				c++
			}
		}
	}
	for _, v := range r {
		fmt.Println(v)
	}
}
