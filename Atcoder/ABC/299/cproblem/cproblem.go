package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
 * https://atcoder.jp/contests/abc299/tasks/abc299_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	s := getNextString(scanner)
	if n == 1 {
		fmt.Println(-1)
		return
	}

	r1 := getMaxDango(s, n)
	alt := reverse(s)
	//fmt.Printf("alt %v \n", alt)
	r2 := getMaxDango(alt, n)

	//fmt.Printf("r1 %v r2 %v \n", r1, r2)
	max := max(r1, r2)
	fmt.Println(max)
	defer writer.Flush()
}

// Split 文字分割
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func getMaxDango(s string, n int) int {
	max := -1
	level := 0
	for i := 0; i < n; i++ {
		if string(s[i]) == "-" {
			if level != 0 && level > max {
				max = level
			}
			level = 0
			continue
		}
		level++
	}
	return max
}

func max(num ...int) int {
	max := math.MinInt32
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
