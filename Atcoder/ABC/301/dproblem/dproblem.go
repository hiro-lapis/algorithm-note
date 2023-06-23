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
 * https://atcoder.jp/contests/abc301/tasks/abc301_d
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	s := getNextString(scanner)
	n := getNextInt(scanner)

	// calc max
	max, min := -1, -1
	for i := 0; i < len(s); i++ {
		v := s[i]
		if string(v) == "1" || string(v) == "?" {
			max = getMax(len(s) - i)
			break
		}
	}

	fmt.Println(strLen)
	fmt.Println(n)

	defer writer.Flush()
}

func getMax(idx int) (max int) {
	for i := 0; i <= idx; i++ {
		max *= 2
	}
	max--
	return
}
