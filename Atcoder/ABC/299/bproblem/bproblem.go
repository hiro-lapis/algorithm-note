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
 * https://atcoder.jp/contests/abc299/tasks/abc299_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	t := getNextInt(scanner)

	cl := make([]int, n)
	tEx := false
	for i := 0; i < n; i++ {
		cl[i] = getNextInt(scanner)
		if !tEx && cl[i] == t {
			tEx = true
		}
	}

	// decide condition of color
	color := t
	if !tEx {
		color = cl[0]
	}

	rl := make([]int, n)
	for j := 0; j < n; j++ {
		rl[j] = getNextInt(scanner)
	}

	ans := 0
	max := 0
	for i := 0; i < n; i++ {
		c := cl[i]
		if c != color {
			continue
		}
		val := rl[i]
		if val > max {
			max = val
			ans = i + 1
		}
	}
	fmt.Println(ans)

	defer writer.Flush()
}
