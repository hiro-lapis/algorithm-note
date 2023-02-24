package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	cl := makeCharList()

	for _, v := range s {
		cl.check(string(v))
	}
	fmt.Println(string(cl.getMin()))
}

func Contain(list []int32, target int32) bool {
	r := true
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			r = false
			break
		}
	}
	return r
}

type charList struct {
	list           []string
	appearanceList []bool
}

func makeCharList() *charList {
	c := charList{
		list: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
			"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
		appearanceList: make([]bool, 26)}
	return &c
}

func (c *charList) check(target string) bool {
	r := false
	for i, s := range c.list {
		if s == target {
			c.appearanceList[i] = true
			r = true
			break
		}
	}
	return r
}

func (c *charList) getMin() string {
	s := "None"
	for i, b := range c.appearanceList {
		if b == false {
			s = c.list[i]
			break
		}
	}
	return s
}
