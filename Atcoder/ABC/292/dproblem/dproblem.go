package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

/** util func */
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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	n := getNextInt(scanner)
	dsu := NewDsu(n)

	m := getNextInt(scanner)
	// edgeがない場合は Yes になり得ない
	if m == 0 {
		fmt.Println("No")
		return
	}

	// edgeの入力値を受け取り、即merge
	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		u := getNextInt(scanner) - 1
		v := getNextInt(scanner) - 1
		edges[i][0] = u
		edges[i][1] = v
		// loop-edgeは処理省略
		if dsu.Same(edges[i][0], edges[i][1]) {
			continue
		}
		dsu.Merge(edges[i][0], edges[i][1])
	}

	// 各edgeのグループも算出し l に設定
	l := map[int]int{}
	for j := 0; j < m; j++ {
		// 制約 u < v により、edge の u を参照することで
		// edgeのリーダーを設定できる
		l[dsu.Leader(edges[j][0])]++
	}
	/**
	 * 3 3
	 * 2 3
	 * 1 1
	 * 2 3
	 *
	 * l: map[0:1 1:2]
	 * edges: [[1 2] [0 0] [1 2]]
	 *
	 */
	fmt.Println(l)
	fmt.Println(edges)
	for v, count := range l {
		if dsu.Size(v) != count {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

	defer writer.Flush()
}

func Solve(dsu *Dsu, edges []Edge) (r bool) {
	for _, edge := range edges {
		u := edge.u
		v := edge.v
		dsu.Merge(u, v)
	}
	/**
	 * N = dsu.n (vertex数)
	 * M = len(edges) (edge数)
	 */
	n := dsu.n
	m := len(edges)
	// 各vertexのリーダーを設定
	vs := make([]int, n)
	for i := 0; i < n; i++ {
		vs[dsu.Leader(i)]++
	}
	// edge set(disjoint setを形成するedgeの数)
	es := make([]int, m)
	for j := 0; j < m; j++ {
		es[dsu.Leader(edges[j].u)]++
	}
	r = reflect.DeepEqual(vs, es)
	return
}

type Edge struct {
	u, v int
}

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
// -1 の場合は
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
