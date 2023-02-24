package main

import (
	"fmt"
)

/**
 * https://atcoder.jp/contests/abc157/tasks/abc157_b
 */
func main() {
	var a = make([]int, 9)
	for i := 0; i < 9; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a[i] = tmp
	}
	var n int
	fmt.Scan(&n)

	var hitList = make([]bool, 9)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		k := GetKey(a, tmp)
		if k != -1 {
			hitList[k] = true
		}
	}

	if Collected(hitList) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func GetKey(s []int, e int) int {
	for i, v := range s {
		if e == v {
			return i
		}
	}
	return -1
}

func Exists(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func CheckT(s []bool, i []int) (result bool) {
	result = true
	for _, v := range i {
		if s[v] != true {
			result = false
			break
		}
	}
	return
}
func Collected(l []bool) bool {
	if CheckT(l, []int{0, 1, 2}) || CheckT(l, []int{3, 4, 5}) || CheckT(l, []int{6, 7, 8}) {
		return true
	}
	if CheckT(l, []int{0, 3, 6}) || CheckT(l, []int{1, 4, 7}) || CheckT(l, []int{2, 5, 8}) {
		return true
	}
	if CheckT(l, []int{0, 4, 8}) || CheckT(l, []int{2, 4, 6}) {
		return true
	}

	return false
}
