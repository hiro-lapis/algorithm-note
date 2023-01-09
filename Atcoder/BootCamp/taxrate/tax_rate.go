package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	Solve(n)
}

func Solve(num float64) {
	price := int(math.Ceil(num / 1.08))
	taxedPrice := math.Floor(float64(price) * 1.08)
	if num != taxedPrice {
		fmt.Println(":(")
	} else {
		fmt.Println(price)
	}
}
