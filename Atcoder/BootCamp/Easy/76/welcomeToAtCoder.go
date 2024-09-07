package main

import "fmt"

const (
	sOK = "AC"
)

/*
*
case1
2 5
1 WA
1 AC
2 WA
2 AC
2 WA

case2
100000 3
7777 AC
7777 AC
7777 AC
*/
//tests := []struct {
//name string
//args args
//want int
//}{

/*
*
answer time 20m
https://atcoder.jp/contests/abc151/tasks/abc151_c
*/
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	penalty := make([]int, n)
	ac := make([]bool, n)

	// s: AC/WA, pNum:problem count
	for i := 0; i < m; i++ {
		var pNum int
		var result string
		fmt.Scan(&pNum, &result)

		if result == sOK {
			ac[pNum-1] = true
			//	if result is WA and don't AC yet
		} else if !ac[pNum-1] {
			penalty[pNum-1]++
		}
	}
	acC := 0
	penaltyC := 0
	for i, r := range ac {
		if r {
			acC++
			penaltyC += penalty[i]
		}
	}
	fmt.Println(acC, penaltyC)
}
