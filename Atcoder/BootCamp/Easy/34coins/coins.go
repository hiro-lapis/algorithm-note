package main

import "fmt"

const aPrice = 500
const bPrice = 100
const cPrice = 50

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var x int
	fmt.Scan(&x)

	// 100 1枚で1通り
	// 100 * 1
	// ↓
	// 100 * 0 + 50 * 2
	ans := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if i*aPrice+j*bPrice+k*cPrice == x {
					ans++
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
