package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	r := Solve(a, b)

	if r {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func Solve(a int, b int) bool {
	if a*2 == b || a*2+1 == b {
		return true
	}
	return false
}
