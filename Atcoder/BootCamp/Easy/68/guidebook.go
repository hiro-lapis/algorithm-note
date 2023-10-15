package main

import (
	"fmt"
	"sort"
)

type Restraunt struct {
	i int
	s string
	p int
}
type Rs []Restraunt

func (r Rs) Len() int      { return len(r) }
func (r Rs) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r Rs) Less(i, j int) bool {
	// 市名が同じならポイントのDESC
	if r[i].s == r[j].s {
		return r[i].p > r[j].p
	}
	// 市名の降順
	return r[i].s < r[j].s
}

// https://atcoder.jp/contests/abc128/tasks/abc128_b
func main() {
	var (
		n  int
		s  []string
		rs Rs
	)

	fmt.Scan(&n)
	rs = make(Rs, 0)
	s = make([]string, n)
	for i := 0; i < n; i++ {
		var tmpS string
		var tmpP int
		fmt.Scan(&tmpS)
		fmt.Scan(&tmpP)
		rs = append(rs, Restraunt{i + 1, tmpS, tmpP})
		s[i] = tmpS
	}

	// 市の名前の昇順でソート(int:asc)
	sort.Sort(rs)
	for i := 0; i < len(rs); i++ {
		fmt.Println(rs[i].i)
	}
}
