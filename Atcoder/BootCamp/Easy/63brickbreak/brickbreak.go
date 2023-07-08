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
 * stdIn/Out にバッファを使うパターン
 * 実行時間33ms
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	c := 0
	next := 1
	for i := 0; i < n; i++ {
		if next == getNextInt(scanner) {
			next++
		} else {
			c++
		}
	}

	if next == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(c)
	}

	defer writer.Flush()
}

/*
 * バッファを使わないパターン
 * 実行時間1sくらい
 */
//func main() {
//	var n int
//	fmt.Scan(&n)
//	c := 0
//	next := 1
//	for i := 0; i < n; i++ {
//		var tmp int
//		fmt.Scan(&tmp)
//		if next == tmp {
//			next++
//		} else {
//			c++
//		}
//	}
//
//	if next == 1 {
//		fmt.Println(-1)
//	} else {
//		fmt.Println(c)
//	}
//}
