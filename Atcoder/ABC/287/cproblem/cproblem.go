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
	m := iost.GetNextInt()
	d := NewDsu(n)
	for i := 0; i < m; i++ {
		u := iost.GetNextInt() - 1
		v := iost.GetNextInt() - 1
		if d.Same(u, v) {
			continue
		}
		d.Merge(u, v)
	}
	if d.Size(0) != n {
		iost.Println("No")
		return
	}
	if n != m+1 {
		iost.Println("No")
		return
	}
	iost.Println("Yes")
}

// Dsu Data structures and algorithms for disjoint set union problems
type Dsu struct {
	n            int
	parentOrSize []int
}

// NewDsu Constructor
func NewDsu(n int) *Dsu {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = -1
	}
	return &Dsu{parentOrSize: p, n: n}
}

// Merge adds an edge (a, b).
func (d *Dsu) Merge(a, b int, meld ...func(int, int)) int {
	x := d.Leader(a)
	y := d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	for _, f := range meld {
		f(x, y)
	}
	return x
}

// Same returns whether the vertices a and b are in the same connected component.
func (d *Dsu) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader returns the representative of the connected component that contains the vertex a.
func (d *Dsu) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size returns the size of the connected component that contains the vertex a.
func (d *Dsu) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups divides the graph into connected components and returns the list of them.
func (d *Dsu) Groups() [][]int {
	result := make([][]int, d.n)
	groups := make([][]int, 0)
	for i := 0; i < d.n; i++ {
		l := d.Leader(i)
		result[l] = append(result[l], i)
	}
	for i := 0; i < d.n; i++ {
		if result[i] == nil {
			continue
		}
		groups = append(groups, result[i])
	}
	return groups
}
