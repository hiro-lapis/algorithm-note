package main

import (
	"competitive-programming/Atcoder/util"
	"fmt"
)

func main() {
	// 10
	d := util.NewDsu(10)
	// &{10 [-1 -1 -1 -1 -1 -1 -1 -1 -1 -1]}
	fmt.Println(d)

	if !d.Same(1, 2) {
		d.Merge(1, 2)
	}
	// &{10 [-1 -2 1 -1 -1 -1 -1 -1 -1 -1]}
	fmt.Printf("%v \n", d)
	d.Merge(1, 3)
	// &{10 [-1 -3 1 1 -1 -1 -1 -1 -1 -1]}
	fmt.Printf("%v \n", d)
	// 1
	fmt.Println(d.Leader(1))
	// 1
	fmt.Println(d.Leader(2))
	// 9
	fmt.Println(d.Leader(9))
	// [[0] [1 2 3] [4] [5] [6] [7] [8] [9]]
	fmt.Println(d.Groups())
	// 8
	fmt.Println(len(d.Groups()))
}
