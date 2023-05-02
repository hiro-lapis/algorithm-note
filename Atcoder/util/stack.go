package util

import "fmt"

type Stack struct {
	list []int
}

func NewStack() *Stack {
	tmp := make([]int, 0)
	return &Stack{list: tmp}
}

func (s *Stack) Push(v int) {
	s.list = append(s.list, v)
	return
}

func (s *Stack) Pop() (v int) {
	tail := len(s.list) - 1
	v = s.list[tail]
	s.list = s.list[:tail]
	return
}

func (s *Stack) Show() {
	fmt.Printf("cap:%v len:%v \n", cap(s.list), len(s.list))
	for k, v := range s.list {
		fmt.Printf("%v:%v, ", k, v)
	}
}
