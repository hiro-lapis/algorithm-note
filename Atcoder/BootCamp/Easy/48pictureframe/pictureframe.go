package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	l := make([]string, h+1)
	for i := 0; i <= h+1; i++ {
		s := "#"
		if IsFrame(i, h) {
			for j := 0; j < w; j++ {
				s += "#"
			}
		} else {
			var tmp string
			fmt.Scan(&tmp)
			s += tmp
		}
		s += "#"
		l = append(l, s)
	}
	for _, str := range l {
		fmt.Println(str)
	}
}

func IsFrame(hi, h int) bool {
	if hi == 0 || hi-1 >= h {
		return true
	}
	return false
}
