package main

import (
	"fmt"
	"math"
	"sort"
)

/*
 * https://atcoder.jp/contests/abc138/tasks/abc138_c
 */
func main() {
	var n int
	fmt.Scan(&n)
	list := make([]float64, n)
	for i := 0; i < n; i++ {
		var tmp float64
		fmt.Scan(&tmp)
		list[i] = tmp
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	mix := list[0]
	for j := 1; j < n; j++ {
		mix = (mix + list[j]) / 2
	}
	if IsInt(mix) {
		fmt.Println(int(mix))
	} else {
		fmt.Println(mix)
	}
}

func IsInt(n float64) bool {
	return int(n) == int(math.Ceil(n))
}
