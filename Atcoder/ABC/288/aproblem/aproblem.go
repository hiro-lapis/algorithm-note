package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)

	l := make([]int64, n)
	for i := 0; i < int(n); i++ {
		var tmp, tmp2 int64
		fmt.Scan(&tmp, &tmp2)
		l[i] = tmp + tmp2
	}
	for _, v := range l {
		fmt.Println(v)
	}
}
