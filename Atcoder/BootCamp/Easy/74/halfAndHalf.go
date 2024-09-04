package main

import (
	"fmt"
)

// solve time: 10m
func main() {
	var a, b, c, x, y int
	fmt.Scan(&a, &b, &c, &x, &y)
	var sum int
	// compare prices between each single pizza and 2 half pizza
	commonPrice := a + b
	if a+b > c*2 {
		commonPrice = c * 2
	}
	// buy x,y same counts
	commonCount := x
	largerPizzaPrice := b
	if x > y {
		commonCount = y
		largerPizzaPrice = a
	}
	sum += commonPrice * commonCount

	// buy remains of larger pizzas
	diffCount := abs(x - y)
	if largerPizzaPrice > c*2 {
		sum += c * 2 * diffCount
	} else {
		sum += largerPizzaPrice * diffCount
	}
	fmt.Println(sum)
}

// Abs  5 -> 5, -5 -> 5
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
