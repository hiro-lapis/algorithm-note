package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)
	list := map[string]int{}
	for i := 0; i < len(w); i++ {
		list[string(w[i])]++
	}
	odd := false
	for _, i := range list {
		if i%2 != 0 {
			fmt.Println("No")
			odd = true
			break
		}
	}
	if odd {
		return
	}
	fmt.Println("Yes")
}
