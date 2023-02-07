package main

import "fmt"

func main() {
	var n, p, q, r, s int
	fmt.Scan(&n, &p, &q, &r, &s)

	list := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		list[i] = tmp
	}
	j := 0
	for i := p; i <= q; i++ {
		tmp2 := list[i-1]
		list[i-1] = list[r-1+j]
		list[r-1+j] = tmp2
		j++
	}
	for _, v := range list {
		fmt.Println(v)
	}
}

//8 1 3 5 7
//1 2 3 4 5 6 7 8
