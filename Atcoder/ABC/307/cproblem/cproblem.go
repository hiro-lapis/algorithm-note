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
 * https://atcoder.jp/contests/abc307/tasks/abc307_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	ha := getNextInt(scanner)
	wa := getNextInt(scanner)
	aa := make([]string, ha)
	for i := 0; i < ha; i++ {
		aa[i] = getNextString(scanner)
	}

	hb := getNextInt(scanner)
	wb := getNextInt(scanner)
	bb := make([]string, hb)
	for i := 0; i < hb; i++ {
		bb[i] = getNextString(scanner)
	}

	hx := getNextInt(scanner)
	wx := getNextInt(scanner)
	xx := make([]string, hx)
	for i := 0; i < hx; i++ {
		xx[i] = getNextString(scanner)
	}

	fmt.Println(v)

	defer writer.Flush()
}
