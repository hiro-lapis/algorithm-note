package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	xA := x1 - x2
	yA := y1 - y2

	x3 := x2 + yA
	y3 := y2 - xA

	xB := x2 - x3
	yB := y2 - y3

	x4 := x3 + yB
	y4 := y3 - xB
	fmt.Println(x3, y3, x4, y4)
}
