package main

import (
	"bufio"
	"fmt"
	"math"
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
 *  https://atcoder.jp/contests/abc309/tasks/abc309_a
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	a := getNextInt(scanner)
	b := getNextInt(scanner)
	if Abs(a-b) > 1 {
		fmt.Println("No")
		return
	}
	max := Max(a, b)
	if max == 4 || max == 7 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
	defer writer.Flush()
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Max(num ...int) int {
	max := math.MinInt32
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
