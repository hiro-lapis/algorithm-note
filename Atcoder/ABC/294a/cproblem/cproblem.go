package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
 * https://atcoder.jp/contests/abc294/tasks/abc294_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	m := getNextInt(scanner)
	A := make([]int, 0)

	for i := 0; i < n; i++ {
		A = append(A, getNextInt(scanner))
	}
	B := make([]int, 0)
	for i := 0; i < m; i++ {
		B = append(B, getNextInt(scanner))
	}

	// merge slice A and B
	C := append(A, B...)
	sort.Slice(C, func(i, j int) bool {
		return C[i] < C[j]
	})

	for j := 0; j < n; j++ {
		fmt.Println(1 + BinarySearch(A[j], C, len(C)/2, len(C)/2))
	}
	for k := 0; k < m; k++ {
		fmt.Println(1 + BinarySearch(B[k], C, len(C)/2, len(C)/2))
	}

	defer writer.Flush()
}

func BinarySearch(target int, list []int, next int, half int) int {
	if len(list) == 1 {
		return -1
	}

	if list[next] == target {
		return next
	}

	num := half / 2
	if num == 0 {
		num++
	}
	if list[next] > target {
		return BinarySearch(target, list, next-num, half/2)
	} else {
		return BinarySearch(target, list, next+num, half/2)
	}
}
