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
 * https://atcoder.jp/contests/abc303/tasks/abc303_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	m := getNextInt(scanner)

	memo := make([]string, 0)
	for i := 0; i < m; i++ {
		l := make([]int, n)
		for j := 0; j < n; j++ {
			l[j] = getNextInt(scanner)
			if j > 0 {
				p := pair(l[j], l[j-1])
				if !Contains(memo, p) {
					memo = append(memo, p)
				}
			}
		}
	}
	ans := countPairs(n) - len(memo)
	if ans < 0 {
		ans = 0
	}
	fmt.Println(ans)
	defer writer.Flush()
}

func pair(n int, n2 int) string {
	if n > n2 {
		return strconv.Itoa(n2) + strconv.Itoa(n)
	}
	return strconv.Itoa(n) + strconv.Itoa(n2)
}

func countPairs(n int) int {
	return n * (n - 1) / 2
}

func Contains(l []string, s string) bool {
	r := false
	for i := 0; i < len(l); i++ {
		if l[i] == s {
			r = true
			break
		}
	}
	return r
}
