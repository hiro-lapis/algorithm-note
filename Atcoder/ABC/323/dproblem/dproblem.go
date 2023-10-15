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

type s struct {
	size  int
	count int
}

// https://atcoder.jp/contests/abc323/tasks/abc323_d
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)

	slimes := map[int]int{}
	for i := 0; i < n; i++ {
		s := getNextInt(scanner)
		c := getNextInt(scanner)
		for s&1 == 0 {
			s >>= 1
			c <<= 1
		}
		slimes[s] += c
	}

	ans := 0
	for _, v := range slimes {
		for v != 0 {
			if (v & 1) == 1 {
				ans++
			}
			v >>= 1
		}
	}
	fmt.Println(ans)
	defer writer.Flush()
}
