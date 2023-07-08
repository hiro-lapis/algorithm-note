package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	aa := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		aa = append(aa, tmp)
	}

	c := 0
	ans := -1
	visited := make([]int, 0)
	for _, v := range aa {
		c++
		if v == 2 {
			ans = c
			break
		} else {
			if Contains(visited, v) {
				break
			}
			visited = append(visited, v)
		}
	}
	fmt.Println(ans)
}

func Contains(l []int, i int) bool {
	r := false
	for _, v := range l {
		if v == i {
			r = true
			break
		}
	}
	return r
}
