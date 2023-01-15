package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 1; i < n; i++ {
		l := 0
		//fmt.Printf("l の初期値は %d \n", l)
		for j := 0; j < n; j++ {
			if i+j >= n {
				break
			}
			front := s[j : j+1]
			back := s[j+i : j+i+1]
			if front == back {
				break
			}
			l++
		}
		fmt.Println(l)
	}
}
