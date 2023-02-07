package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	list := make([]int, n)
	for i := 1; i <= n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		list[tmp-1] = i
	}

	for _, v := range list {
		fmt.Println(v)
	}
}
