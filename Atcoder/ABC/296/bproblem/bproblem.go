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

/*
 * https://atcoder.jp/contests/abc295/tasks/abc295_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// height
	r := getNextInt(scanner)
	// width
	c := getNextInt(scanner)
	ll := make([]string, r)
	bombIj := make([]Idx, 0)
	// 爆発前のmap作成
	for i := 0; i < r; i++ {
		tmp := getNextString(scanner)
		ll = append(ll, tmp)
		// bomb index mapping
		bombIdx := findIndexesBomb(tmp)
		if len(bombIdx) > 0 {
			for _, idx := range bombIdx {
				// i:縦、j:横で爆弾のポジションを記録
				bombIj = append(bombIj, Idx{i, idx})
			}
		}
	}
	// 爆弾によるマップの加工
	for _, idx := range bombIj {
		//y := idx.i
		//x := idx.j
		ran, _ := strconv.Atoi(ll[idx.i])
		// x,yを起点に周囲ranマスを加工
		for y := 0; y < idx.i; y++ {
			for x := 0; x < idx.j; x++ {
				str := ll[idx.j]

			}

		}

	}
	// print ans
	for _, v := range ll {
		fmt.Println(v)
	}

	defer writer.Flush()
}

type Idx struct {
	i, j int
}

// findIndexesBomb
func findIndexesBomb(s string) []int {
	r := make([]int, 0)
	for i, v := range s {
		_, err := strconv.Atoi(string(v))
		if err == nil {
			r = append(r, i)
		}
	}
	return r
}

// makeMatrix 縦:y 横:x の2次元sliceを作成
func makeMatrix(x, y int) [][]string {
	ll := make([][]string, y)
	for i := range ll {
		ll[i] = make([]string, x)
	}
	return ll
}
