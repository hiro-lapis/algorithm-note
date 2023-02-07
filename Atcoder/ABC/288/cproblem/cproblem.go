package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/* dsu */

// Dsu Data structures and algorithms for disjoint set union problems
type Dsu struct {
	n            int
	parentOrSize []int
}

// NewDsu Constructor
func NewDsu(n int) *Dsu {
	p := make([]int, n)
	// 初期のparent nodeなしの状態は-1という値で表す
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
	// 値が小さい方をy,大きい方をxに代入する
	// 同値の場合はそのまま
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	// x をleader にする形で調整
	// x のサイズを増やす(-1=>-2=>-3...)
	d.parentOrSize[x] += d.parentOrSize[y]
	// 取り込まれる側は親のindex値を入れる
	d.parentOrSize[y] = x
	// 可変長引数なので省略されることあり
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

/* END dsu */

/* receive input func */
func getScanner(fp *os.File) *bufio.Scanner {
	// create buffered scanner
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
	return scanner
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

/* END receive input func */

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// n and m is number of vertices, edges
	n := getNextInt(scanner)
	m := getNextInt(scanner)
	// create disjoint set union of number of n-vertices
	d := NewDsu(n)

	ans := 0
	for i := 0; i < m; i++ {
		// receive edge info
		u := getNextInt(scanner)
		v := getNextInt(scanner)
		// adjust input to slice index
		u--
		v--
		if !d.Same(u, v) {
			d.Merge(u, v)
		} else {
			ans++
		}
	}
	// out put redundant edges
	fmt.Println(ans)
	defer writer.Flush()
}
