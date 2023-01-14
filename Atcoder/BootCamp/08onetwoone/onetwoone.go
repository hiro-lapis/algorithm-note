package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * https://atcoder.jp/contests/abc086/tasks/abc086_b
 */
func main() {
	var a, b int

	fmt.Scan(&a, &b)
	c, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	root := math.Sqrt(float64(c))
	if IsInt(root) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func IsInt(x float64) bool {
	return math.Floor(x) == x
}
