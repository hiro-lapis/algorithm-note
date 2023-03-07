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

/**
 * https://atcoder.jp/contests/abc152/tasks/abc152_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	min := math.MaxInt64
	answer := 0
	for i := 0; i < n; i++ {
		tmp := getNextInt(scanner)
		if min > tmp {
			answer++
			min = tmp
		}
	}
	fmt.Println(answer)

	defer writer.Flush()
}
