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
	// 改行区切りで入力を受け取り
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
 * https://atcoder.jp/contests/abc294/tasks/abc294_d
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	l := make([]int, n)
	for i := 1; i <= n; i++ {
		l[i-1] = i
	}

	q := getNextInt(scanner)
	for i := 0; i < q; i++ {
		// 空白込みで入力を受け取りできないので、2の時だけ追加入力を受け取り
		event := getNextString(scanner)
		if event == "2" {
			// 2の処理
			fmt.Println(getNextInt(scanner))
			continue
		}
		if event == "1" {

			continue
		}
		if event == "3" {
			fmt.Println("33333")
			continue
		}

	}

	defer writer.Flush()
}

// わからないこと
// sliceの変数をポインタ渡しで渡す関数の書き方
// sliceの先頭の要素を削除する方法
//slice の任意index要素を削除する方法

// EventOne a
func EventOne(list *[]int) int {
	r := list[0]
	list = list[1:]
	return r
}
