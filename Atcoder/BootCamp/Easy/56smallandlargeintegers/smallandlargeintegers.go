package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc093/tasks/abc093_b
 * solve time 16m
 */
func main() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	fLimit := a + n
	bLimit := b - n
	for i := a; i <= b; i++ {
		if i < fLimit || i > bLimit {
			fmt.Println(i)
		}
	}
}
