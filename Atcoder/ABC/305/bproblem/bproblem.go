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
 * https://atcoder.jp/contests/abc305/tasks/abc305_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	p := getNextString(scanner)
	q := getNextString(scanner)

	pV := getV(p)
	qV := getV(q)

	fmt.Println(Abs(pV - qV))
	defer writer.Flush()
}

func getV(str string) int {
	if str == "A" {
		return 0
	}
	if str == "B" {
		return 3
	}
	if str == "C" {
		return 4
	}
	if str == "D" {
		return 8
	}
	if str == "E" {
		return 9
	}
	if str == "F" {
		return 14
	}
	return 23
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
