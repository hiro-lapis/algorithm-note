package main

import "fmt"

/*
 * https://atcoder.jp/contests/abc109/tasks/abc109_b
 */
func main() {
	var n int
	fmt.Scan(&n)

	l := make([]string, 0)
	exists := false
	violated := false
	appendix := ""
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		if appendix != "" && appendix != string(tmp[0]) {
			violated = true
			break
		}
		appendix = string(tmp[len(tmp)-1])
		if Contain(l, tmp) {
			exists = true
			break
		}
		l = append(l, tmp)
	}
	if exists {
		fmt.Println("No")
		return
	}
	if violated {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}

func Contain(l []string, target string) (r bool) {
	for _, word := range l {
		if word == target {
			r = true
			break
		}
	}
	return
}
