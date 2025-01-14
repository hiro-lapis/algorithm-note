package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scan(&n)
	num := 0
	for i := 0; i < len(n); i++ {
		// 先頭の数cは他の桁と異なる処理をしたいのでスルー
		if i == 0 {
			continue
		}
		val, _ := strconv.Atoi(string(n[i]))
		if val != 9 {
			num = -1
			break
		}
	}
	c, _ := strconv.Atoi(string(n[0]))
	fmt.Println(c + 9*(len(n)-1) + num)
}
