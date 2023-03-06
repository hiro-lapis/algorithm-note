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
 * https://atcoder.jp/contests/abc291/tasks/abc291_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	ans := 0
	for x := 1; x < n; x++ {
		ans += factorialCount(x) * factorialCount(n-x)
	}
	fmt.Println(ans)

	defer writer.Flush()
}

func factorialCount(x int) (count int) {
	// X の値を構成するA,Bの範囲 √X を探索する
	// X 自体がX の約数となるため探索範囲に含める
	for a := 1; a*a <= x; a++ {
		// 約数でないときは組み合わせになり得ないため次のループへ
		if x%a != 0 {
			continue
		}
		b := x / a
		if a == b {
			count++
			// A B ともに異なる約数の場合は組み合わせが2通りあるため足す
		} else {
			count += 2
		}
	}
	return
}
