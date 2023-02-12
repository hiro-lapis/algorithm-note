package main

import "fmt"

/*
 * bit:10 = 2
 * bit:110 = 6
 * bit:1000 = 8
 */
const (
	A uint = 10 // bit: 1010
	B uint = 12 // bit: 1100
)

func main() {
	var bits uint

	// AND: each bit in A&B is true, then 1
	bits = A & B // bit:1000 = 8
	fmt.Println(bits)

	// OR: unless bit in A or B is true, then 1
	bits = A | B // bit: 1110 = 14
	fmt.Println(bits)

	// XOR: unless bit in A or B is false, then 1.
	// both false case, then 0
	bits = A ^ B // 0110
	fmt.Println(bits)

	// AND NOT: both bit in A or B is false, then 1.
	bits = A &^ B // 0010
	fmt.Println(bits)
	
	// 左シフト演算
	bits = 1 << uint64(3) // 1000 : 2の3乗かかる
	// 右シフト演算
	bits = 8 >> uint64(3) // 0001 : 2の(-3)乗かかる
}
