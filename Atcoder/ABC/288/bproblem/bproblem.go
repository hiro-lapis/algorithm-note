package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	s := make([]string, k)
	for i := 0; i < k; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s[i] = tmp
	}
	ans := StringSort(s, k)
	for _, v := range ans {
		fmt.Println(v)
	}
}

func StringSort(s []string, num int) []string {
	r := make([]string, num)
	sort.Strings(s)

	for i := 0; i < num; i++ {
		r[i] = s[i]
	}
	return r
}
