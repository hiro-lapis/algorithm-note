package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc063/tasks/abc063_b
 */
func main() {
	var s string
	fmt.Scan(&s)

	l := make([]int32, 0)
	diff := true
	for _, v := range s {
		if Contain(l, v) {
			diff = false
			break
		} else {
			l = append(l, v)
		}
	}
	if diff {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func Contain(l []int32, target int32) (r bool) {
	for _, v := range l {
		if v == target {
			r = true
		}
	}
	return
}
