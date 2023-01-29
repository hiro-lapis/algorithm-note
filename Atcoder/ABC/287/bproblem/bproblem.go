package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s[i] = tmp
	}

	t := make([]string, m)
	for j := 0; j < m; j++ {
		var tmp string
		fmt.Scan(&tmp)
		t[j] = tmp
	}
	var c int
	for k := 0; k < n; k++ {
		tail := s[k][3:]
		for l := 0; l < m; l++ {
			if t[l] == tail {
				c++
				break
			}
		}
	}
	fmt.Println(c)
}
