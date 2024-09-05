package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		h[i] = tmp
	}
	// sort desc
	sort.Slice(h, func(i, j int) bool {
		return h[i] > h[j]
	})
	min := math.MaxInt
	for i := 0; i < len(h); i++ {
		// avoid over index
		if i+k-1 > len(h)-1 {
			break
		}
		diff := h[i] - h[i+k-1]
		if min > diff {
			min = diff
		}
	}
	fmt.Println(min)
}
