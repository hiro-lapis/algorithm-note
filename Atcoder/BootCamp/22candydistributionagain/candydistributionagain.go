package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	list := make([]int, n)

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		list[i] = tmp
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	sort.Ints(list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	var ans int
	tmp2 := x
	for i := 0; i < n; i++ {
		if list[i] > tmp2 {
			break
		}
		if i == n-1 && tmp2 > list[i] {
			break
		}

		tmp2 -= list[i]
		if tmp2 >= 0 {
			ans++
		} else {
			break
		}
	}
	fmt.Println(ans)
}
