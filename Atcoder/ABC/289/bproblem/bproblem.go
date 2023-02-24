package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var n, m int
	//fmt.Scan(&n, &m)

	l := make([]int, 0)
	fmt.Println(l)
	fmt.Printf("len=%v cap=%v \n", len(l), cap(l))
	fmt.Println("next \n")

	l = append(l, 10)
	l = append(l, 10)
	l = append(l, 10)
	fmt.Printf("len=%v cap=%v \n", len(l), cap(l))
	//for i := 0; i < m; i++ {
	//	var tmp int
	//	fmt.Scan(&tmp)
	//	l[i] = tmp
	//}
	//ans := make([]int, n)
	//stack := make([]int, n)
	//stackI := 0
	//for j := 0; j <= n; j++ {
	//	if Contain(l, j+1) {
	//		stack[stackI] = j + 1
	//		stackI++
	//	} else {
	//		ans = Insert(ans, j)
	//		for i := stackI; i >= 0; i-- {
	//			ans = Insert(ans, stack[stackI])
	//			stack[stackI] = 0
	//			if i == 0 {
	//				stackI = 0
	//			}
	//		}
	//	}
	//}
	//fmt.Println(ans)
}

func Contain(list []int, target int) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	
	return false
}

func Insert(list []int, val int) []int {
	r := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			r[i] = list[i]
		} else {
			r[i] = val
			break
		}
	}
	return r
}

type Stack struct {
	v   *list.List
	len int
}

func NewStack() *Stack {
	return &Stack{v: list.New(), len: 0}
}

func (s *Stack) Push(v int) {
	s.v.PushBack(v)
	s.len++
}

func (s *Stack) Pop() any {
	if s.Length() == 0 {
		return 0
	}
	b := s.v.Back()
	//if b == nil {
	//	return nil
	//}
	s.len--
	return s.v.Remove(b)
}

func (s *Stack) Length() int {
	return s.len
}
