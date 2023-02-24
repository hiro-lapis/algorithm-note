package main

import "fmt"

func main() {
	var h uint64
	fmt.Scan(&h)

	fmt.Println(h)
	h >>= 1
	fmt.Println(h)
	var ans, c uint64
	ans = 0
	c = 1
	for h > 0 {
		ans += c
		h /= 2
		c *= 2
	}

	// bit shift style
	ans = 0
	bit := 0
	for h > 0 {
		h >>= 1
		ans += 1 << uint(bit)
		bit++
	}

	fmt.Println(ans)
}

func Pow(n uint64, count int) uint64 {
	ans := n
	if count == 1 {
		return ans
	}
	for i := 1; i < count; i++ {
		ans *= ans
	}
	return ans
}
