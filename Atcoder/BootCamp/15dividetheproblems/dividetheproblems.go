package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		d[i] = tmp
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	ans := Divide(d)
	fmt.Println(ans)
}

func Divide(list []int) int {
	ans := 0
	mid := len(list) / 2
	front := list[mid-1]
	back := list[mid]
	if front == back {
		return ans
	}
	return back - front
}
