package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

/** util func */
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

/**
 * https://atcoder.jp/contests/abc297/tasks/abc297_d
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	iost = NewIost(fp, wfp)
	a := iost.GetNextInt()
	b := iost.GetNextInt()
	// 割り算の商のINT値のoverflow(RE)を避けるため公約数で割っておく(計算回数には影響なし)
	g := gcd(a, b)
	a /= g
	b /= g
	// 割り算して計算回数を省略するため、b/aの前提を作る
	if a < b {
		a, b = b, a
	}
	c := makeEven(a, b, 0)
	fmt.Println(c)
	defer iost.Writer.Flush()
}

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

func makeEven(a, b, c int) int {
	if a == b {
		return c
	} else if a == 1 || b == 1 {
		if a == 1 {
			return c + b - a
		} else {
			return c + a - b
		}
	}
	if a > b {
		a, b = b, a
	}
	c += b / a
	return makeEven(a, b%a, c)
}
