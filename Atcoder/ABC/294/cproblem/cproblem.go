package main

import (
	"bufio"
	"fmt"
	"os"
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

func getNextInt64(scanner *bufio.Scanner) int64 {
	i, _ := strconv.ParseInt(getNextString(scanner), 10, 64)
	return i
}

func getNextUint64(scanner *bufio.Scanner) uint64 {
	i, _ := strconv.ParseUint(getNextString(scanner), 10, 64)
	return i
}

func getNextFloat64(scanner *bufio.Scanner) float64 {
	i, _ := strconv.ParseFloat(getNextString(scanner), 64)
	return i
}

/**
 * https://atcoder.jp/contests/abc293/tasks/abc293_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)
	/**
	 * 1 << h+w-2 4 => cnt=16
	 * 1 が4回の left shift で 10000になる
	 */
	//cnt := 1 << uint(h+w-2) // ？？
	graph := makeGrid(h, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			graph[i][j] = getNextInt(scanner)
		}
	}

	//ans := bitSearch(cnt, graph, h, w)
	//fmt.Printf("ans>>%v \n", ans)

	st := make(map[int]bool) // 探索隅マス目の数値リスト
	var dfs func(i, j int)
	dAns := 0
	dfs = func(i, j int) {
		if st[graph[i][j]] { // 嬉しくない場合
			return
		}
		if i == h-1 && j == w-1 { // ゴール到達
			dAns++
			return
		}
		st[graph[i][j]] = true // 探索すみにする

		if i != h-1 { // 縦軸がゴール未到達の時
			dfs(i+1, j)
		}
		if j != w-1 { // 横軸がゴール未到達の時
			dfs(i, j+1)
		}
		st[graph[i][j]] = false // 別ルート探索のため過去の探索値をリセット
	}
	dfs(0, 0)
	fmt.Println(dAns)

	defer writer.Flush()
}

func Contain(list []int, num int) bool {
	r := false
	for _, v := range list {
		if v == num {
			r = true
			break
		}
	}
	return r
}

// 縦横指定したindex の二次元配列を作成
// index[h][w]
func makeGrid(h, w int) [][]int {
	index := make([][]int, h)
	data := make([]int, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}

func bitSearch(cnt int, graph [][]int, h int, w int) int {
	ans := 0
	for i := 0; i < cnt; i++ {
		y := 0
		x := 0
		tap := map[int]bool{}
		tap[graph[0][0]] = true // スタート地点を探索済にする

		// 探索ロジック
		// 移動回数はh+w-2回 ex.3*3マスなら4回まで移動
		for j := 0; j < h+w-2; j++ {
			if i>>uint(j)&1 == 1 {
				y++
			} else {
				x++
			}
			if y < h && x < w {
				tap[graph[y][x]] = true
			}
		}
		// tap(移動回数)がゴールまで辿り着いてたら解答ルートに追加
		if len(tap) == h+w-1 {
			ans++
		}
	}
	return ans
}
