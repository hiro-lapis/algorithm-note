package main

import "fmt"

/*
 * https://atcoder.jp/contests/agc003/tasks/agc003_a
 */
func main() {
	var s string
	fmt.Scan(&s)
	var N, S, W, E bool
	for _, v := range s {
		align := string(v)
		if align == "N" {
			N = true
		}
		if align == "S" {
			S = true
		}
		if align == "E" {
			E = true
		}
		if align == "W" {
			W = true
		}
	}
	if ((N && S) || (!N && !S)) == false {
		fmt.Println("No")
		return
	}
	if ((W && E) || (!W && !E)) == false {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")

}
