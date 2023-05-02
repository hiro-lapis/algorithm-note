package util

/**
 * Atcoder ではGenerics を使えないので注意
 */

type Graph[T any] struct {
	vertices  []*Vertex[T]
	edges     []*Edge
	adjacency [][]*Edge
}

//	func (g *Graph) AddVertex(vertex int) error {
//		if contains(g.vertices, vertex) {
//			err := fmt.Errorf("Vertex %d already exists", vertex)
//			return err
//		} else {
//			v := &Vertex{
//				key: vertex,
//			}
//			g.vertices = append(g.vertices, v)
//		}
//		return nil
//	}

/**
 * disjoint set union with union find
 * how to use
 *
 * d := Dsu(10) => d.n: 10, parentOrSize:[-1 -1 -1 -1 -1 -1 -1 -1 -1 -1]
 * d.Same(1, 2)  // judge 1v and 2v are same group
 * d.Leader(1) // get 1v's root node
 * d.Size(1) // get 1v's group's vertex count
 * d.Merge(1, 2) => &{10 [-1 -2 1 -1 -1 -1 -1 -1 -1 -1]}
 *
 * https://atcoder.jp/contests/abc287/submissions/38383558
 */

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

// Alternative implementation

//type UnionFind struct {
//	par  []int
//	size []int
//}
//
//func newUnionFind(n int) *UnionFind {
//	rs := new(UnionFind)
//	rs.par = make([]int, n+1)
//	rs.size = make([]int, n+1)
//
//	for i := 0; i < n+1; i++ {
//		rs.par[i] = -1
//		rs.size[i] = 1
//	}
//	return rs
//}
//
//func (uf UnionFind) root(x int) int {
//	if uf.par[x] == -1 {
//		return x
//	}
//	uf.par[x] = uf.root(uf.par[x])
//	return uf.par[x]
//}
//
//func (uf UnionFind) union(x, y int) {
//	x = uf.root(x)
//	y = uf.root(y)
//
//	if x == y {
//		return
//	}
//	if uf.getsize(x) < uf.getsize(y) {
//		x, y = y, x
//	}
//	uf.par[y] = x
//	uf.size[x] += uf.size[y]
//}
//
//func (uf UnionFind) find(x, y int) bool {
//	return uf.root(x) == uf.root(y)
//}
//
//func (uf UnionFind) getsize(x int) int {
//	return uf.size[uf.root(x)]
//}

/** Dsu **/

type Vertex[T any] struct {
	key T
}
