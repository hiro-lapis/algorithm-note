package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc049/tasks/abc049_b
 */
func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	l := make([]string, 0)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Scan(&tmp)
		l = append(l, tmp, tmp)
	}

	for _, s := range l {
		fmt.Println(s)
	}
}
