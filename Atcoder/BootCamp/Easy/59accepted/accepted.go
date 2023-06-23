package main

import "fmt"

/*
 * 12m
 * https://atcoder.jp/contests/abc104/tasks/abc104_b
 */
func main() {
	var s string
	fmt.Scan(&s)

	r := true
	c1 := false
	for i, v := range s {
		str := string(v)
		if i == 0 {
			if str != "A" {
				r = false
				break
			}
			continue
		}
		if i >= 2 && i <= (len(s)-1)-1 {
			if str == "C" {
				if !c1 {
					c1 = true
					continue
				} else {
					r = false
					break
				}
			}
		}
		if !isLowerCase(str) {
			r = false
			break
		}
	}
	if r && c1 {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}

}

func isLowerCase(str string) bool {
	for _, char := range str {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
