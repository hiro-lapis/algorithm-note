package main

import (
	"fmt"
)

/**
 * https://atcoder.jp/contests/abc124/tasks/abc124_c
 */
func main() {
	var s string
	fmt.Scan(&s)

	/**
	 * 10010
	 * before: 10 after: 010
	 */
	l := len(s)
	// pattern of 010101...
	c01 := 0
	for i := 0; i < l; i++ {
		if i%2 == 0 {
			if string(s[i]) != "0" {
				c01++
			}
			continue
		}
		//
		if string(s[i]) != "1" {
			c01++
		}
	}

	// pattern of 10101...
	c10 := 0
	for j := 0; j < l; j++ {
		if j%2 == 0 {
			if string(s[j]) != "1" {
				c10++
			}
			continue
		}
		if string(s[j]) != "0" {
			c10++
		}
	}

	if c01 > c10 {
		fmt.Println(c10)
		return
	}
	fmt.Println(c01)
}
