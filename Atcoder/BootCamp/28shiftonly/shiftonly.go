package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	l := make([]int, n)
	even := true
	count := 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if !IsEven(tmp) {
			even = false
		}
		l[i] = tmp
	}
	if !even {
		fmt.Println(count)
		return
	}
	for {
		ll := make([]int, n)
		for k, v := range l {
			even = even && IsEven(v)
			ll[k] = v / 2
			if !even {
				break
			}
		}
		if even {
			count++
			l = ll
		} else {
			break
		}
	}
	fmt.Println(count)
}

func IsEven(num int) bool {
	return num%2 == 0
}
