package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var d, x int
	fmt.Scan(&d, &x)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		list[i] = tmp
	}

	ans := x
	for j := 0; j < len(list); j++ {
		span := list[j]
		i := 0
		for {
			next := 1 + span*i
			if next <= d {
				ans++
				i++
			} else {
				break
			}
		}
	}
	fmt.Println(ans)
}
