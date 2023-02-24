package main

import (
	"fmt"
	"sort"
)

/**
 * https://atcoder.jp/contests/abc088/tasks/abc088_b
 */
func main() {
	var n, sum int
	fmt.Scan(&n)
	var cards = make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		cards[i] = tmp
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	for j := 0; j < n; j++ {
		if j%2 == 0 {
			sum += cards[j]
		} else {
			sum -= cards[j]
		}
	}
	fmt.Println(sum)
}
