package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if IsEven(a, b, c) && IsSameVal(a, b, c) {
		fmt.Println(-1)
		return
	}

	var count int
	for {
		if !IsEven(a, b, c) {
			break
		}
		count++
		aHalf := a / 2
		bHalf := b / 2
		cHalf := c / 2

		a = bHalf + cHalf
		b = aHalf + cHalf
		c = aHalf + bHalf
	}
	fmt.Println(count)
}

func IsEven(n ...int) bool {
	for _, v := range n {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func IsSameVal[T comparable](x ...T) bool {
	for i := 0; i < len(x)-1; i++ {
		if x[i] != x[i+1] {
			return false
		}
	}
	return true
}
