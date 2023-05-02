package main

import (
	"bufio"
	"competitive-programming/Atcoder/util"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

// main
// util パッケージ群の動作仕様テスト
func main() {
	fmt.Println("stack test")
	//StackTest()
	fmt.Println("PriorityQueue test")
	//PriorityQueueTest()
	fmt.Println("Math test")
	MathTest()
}

// StackTest
func StackTest() {
	stack := util.NewStack()
	fmt.Println("initial length")
	stack.Show()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("after length")
	stack.Show()
	// 最後に追加された要素が最初に取り出される
	// popしたタイミングで配列末尾のポインタを削除する
	fmt.Printf("Pop orde [1st:%v] [2nd:%v] [3rd:%v] \n", stack.Pop(), stack.Pop(), stack.Pop())
	fmt.Println("after popped length")
	stack.Show()
}

// PriorityQueueTest
// edge の仕様確認を含めた動作テスト
func PriorityQueueTest() {
	fmt.Println("Edge spec")
	edge1 := util.Edge{U: 1, V: 2, Weight: 10}
	fmt.Printf("edge1 [from:%v to:%v weight:%v]\n", edge1.U, edge1.V, edge1.Weight)

	// edge weight を省略すると一律 0 になる
	// weight がないgraphでは、この使い方をすることで通常のqueueと同じFIFOの処理ができる
	edge2 := util.Edge{U: 444, V: 555}
	fmt.Printf("edge2 [from:%v to:%v weight:%v]\n", edge2.U, edge2.V, edge2.Weight)

	edge3 := util.Edge{U: 6, V: 7, Weight: 5}
	edge4 := util.Edge{U: 10, V: 9}
	// min-heap base のキュー
	pq := util.NewPriorityQueue()

	// edgeのポインタを渡す
	// weight が低い順に並べて格納、参照される
	// weight が同値のものは最初に追加されたものが最初にくる
	pq.Push(&edge1)
	pq.Push(&edge2)
	pq.Push(&edge3)
	pq.Push(&edge4)
	pq.Show()

	popped1 := pq.Pop()
	popped2 := pq.Pop()
	popped3 := pq.Pop()
	fmt.Printf("[%v U:%v V:%v] cost>>%v \n", "popped1", popped1.U, popped1.V, popped1.Weight)
	fmt.Printf("[%v U:%v V:%v] cost>>%v \n", "popped2", popped2.U, popped2.V, popped2.Weight)
	fmt.Printf("[%v U:%v V:%v] cost>>%v \n", "popped3", popped3.U, popped3.V, popped3.Weight)
}

func MathTest() {
	// 素数判定
	fmt.Printf("IsPrime(5): expect:true >> %v\n", util.IsPrime(5))
	fmt.Printf("IsPrime(114): expect:false >> %v\n", util.IsPrime(114))
	// 素数列挙
	fmt.Printf("GetPrimes(5, 40): expect:[1 2 3 5 7 9 11 13 15 17 19] >> %v\n", util.GetPrimes(1, 20))

	// 最大公約数
	fmt.Printf("Gcd(3, 27): expect:9 >> %v\n", util.Gcd(3, 27))
	fmt.Printf("Gcd(3, 29): expect:0 >> %v\n", util.Gcd(3, 29))
	fmt.Printf("Gcd(30, 12): expect:6 >> %v\n", util.Gcd(30, 12))
	// 素因数分解
	fmt.Printf("GetPrimeFactorials(2020): expect:[2:2 5:1 101:1] >> %v\n", util.GetPrimeFactorials(2020))
}
