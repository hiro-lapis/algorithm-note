package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc123/tasks/abc123_b
 */
func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	list := []int{a, b, c, d, e}

	// 条件1 10の倍数の値は途中で注文する
	// 条件2 最も時間のかかる注文は、最後に注文するのはNG
	// 条件3 最後の注文だけは届く時間が10の切り上げにならない
	// => 1の位が1以上で、かつ最も低い注文を最後にするのが最短である
	minI := 0
	min := 9
	// 最後にorderすべき注文(1の位が最も低い)を探索
	for i := 0; i < len(list); i++ {
		mod := list[i] % 10
		if mod > 0 && mod < min {
			min = mod
			minI = i
		}
	}
	// ラストオーダーの注文を除いて時間を10単位で合算
	t := 0
	for i, v := range list {
		if i != minI {
			t += FloorTen(v)
		}
	}
	fmt.Println(t + list[minI])
}

// FloorTen 10単位の切り上げ
func FloorTen(num int) int {
	x := num / 10
	y := num % 10
	if y != 0 {
		return 10 * (x + 1)
	}
	return 10 * x
}
