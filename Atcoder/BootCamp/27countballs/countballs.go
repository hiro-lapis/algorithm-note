package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if a == 0 {
		fmt.Println(0)
		return
	}

	ans := (n / (a + b)) * a
	remain := n % (a + b)
	if remain >= a {
		ans += a
	} else {
		ans += remain
	}
	fmt.Println(ans)
}
