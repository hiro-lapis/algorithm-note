package main

import (
	"fmt"
	"sort"
)

type kv struct {
	k string
	v int
}

type pairs []kv

// sort処理に必要な最低限の関数を定義してあげる
func (a pairs) Len() int           { return len(a) }
func (a pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pairs) Less(i, j int) bool { return a[i].v > a[j].v }

// https://atcoder.jp/contests/abc155/tasks/abc155_c
func main() {
	var n int
	fmt.Scan(&n)

	// 出現数をカウントしつつ文字列取得
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		m[tmp]++
	}

	// for sort move new data structure
	s := make(pairs, 0)
	for k, v := range m {
		s = append(s, kv{k, v})
	}
	// 昇順でソート(int:asc)
	sort.Sort(s)
	ans := make([]string, 0)
	for i := 0; i < len(m) && s[i].v == s[0].v; i++ {
		ans = append(ans, s[i].k)
	}
	// sort increasing order
	sort.Strings(ans)
	for _, an := range ans {
		fmt.Println(an)
	}
}
