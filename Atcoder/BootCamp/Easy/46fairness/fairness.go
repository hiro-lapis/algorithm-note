package main

import "fmt"

func main() {
	var a, b, c, k int
	fmt.Scan(&a, &b, &c, &k)

	if k%2 == 0 {
		fmt.Println(a - b)
		return
	}
	fmt.Println(b - a)
}
