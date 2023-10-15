package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var iost *Iost

type Iost struct {
	Scanner *bufio.Scanner
	Writer  *bufio.Writer
}

func NewIost(fp io.Reader, wfp io.Writer) *Iost {
	const BufSize = 2000005
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, BufSize), BufSize)
	return &Iost{Scanner: scanner, Writer: bufio.NewWriter(wfp)}
}
func (i *Iost) Text() string {
	if !i.Scanner.Scan() {
		panic("scan failed")
	}
	return i.Scanner.Text()
}
func (i *Iost) Atoi(s string) int                 { x, _ := strconv.Atoi(s); return x }
func (i *Iost) GetNextInt() int                   { return i.Atoi(i.Text()) }
func (i *Iost) Atoi64(s string) int64             { x, _ := strconv.ParseInt(s, 10, 64); return x }
func (i *Iost) GetNextInt64() int64               { return i.Atoi64(i.Text()) }
func (i *Iost) Atof64(s string) float64           { x, _ := strconv.ParseFloat(s, 64); return x }
func (i *Iost) GetNextFloat64() float64           { return i.Atof64(i.Text()) }
func (i *Iost) Print(x ...interface{})            { fmt.Fprint(i.Writer, x...) }
func (i *Iost) Printf(s string, x ...interface{}) { fmt.Fprintf(i.Writer, s, x...) }
func (i *Iost) Println(x ...interface{})          { fmt.Fprintln(i.Writer, x...) }
func isLocal() bool                               { return os.Getenv("NICKEL") == "BACK" }
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	if isLocal() {
		fp, _ = os.Open(os.Getenv("WELL_EVERYBODY_LIES_TOO_MUCH"))
	}
	iost = NewIost(fp, wfp)
	defer func() {
		iost.Writer.Flush()
	}()
	solve()
}
func solve() {
	n := iost.GetNextInt()
	t := iost.Text()
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		s := iost.Text()
		if s == t {
			ans = append(ans, i)
			continue
		}
		if isInsert(s, t) {
			ans = append(ans, i)
			continue
		}
		if isDeleted(s, t) {
			ans = append(ans, i)
			continue
		}
		if isChanged(s, t) {
			ans = append(ans, i)
		}
	}
	iost.Println(len(ans))
	for i := 0; i < len(ans); i++ {
		iost.Println(ans[i] + 1)
	}
}
func isInsert(s, t string) bool {
	// 足した結果同じ文字数にならないならfalse
	// Insertで文字を増やす捜査をTに対して行うため、
	// 文字数においてTが1文字少ないことをチェック
	if len(s) != len(t)+1 {
		return false
	}
	l := 0
	r := len(t) - 1
	// デフォルトを文字数Tにする
	cnt := len(t)
	// 左からl文字めがマッチする限りカウントを減らす
	for l < len(t) && s[l] == t[l] {
		l++
		cnt--
	}
	for r >= l && s[r+1] == t[r] {
		r--
		cnt--
	}
	// 左右それぞれから走査してマッチしないとこまできたらカウントダウン終了
	// 最大1文字の誤差を許容してカウントダウンし、値が0なら一致(合わない文字1だけを削除して一致させられる)
	return cnt == 0
}
func isDeleted(s, t string) bool {
	// ある文字1つを削除して一致するということは、ある文字を挿入することで一致するのと同義
	// 引数だけ入れ替えて実行
	return isInsert(t, s)
}
func isChanged(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	l := 0
	r := len(t) - 1
	cnt := len(t)
	for l < len(t) && s[l] == t[l] {
		l++
		cnt--
	}
	for r >= l && s[r] == t[r] {
		r--
		cnt--
	}
	return cnt == 1
}
