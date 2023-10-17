package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var l, n int
	fmt.Scan(&n)
	fmt.Scan(&l)
	ss := make([]string, n)

	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		ss[i] = tmp
	}
	sort.Strings(ss)
	fmt.Println(strings.Join(ss, ""))
}
