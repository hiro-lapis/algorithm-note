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
 * https://atcoder.jp/contests/abc303/tasks/abc303_a
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	s := getNextString(scanner)
	t := getNextString(scanner)
	if s == t {
		fmt.Println("Yes")
		return
	}

	sim := true
	for i := 0; i < n; i++ {
		ss := string(s[i])
		ss2 := string(t[i])
		if !similar(ss, ss2) {
			sim = false
			break
		}
	}
	if sim {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	defer writer.Flush()
}

func similar(s string, s2 string) bool {
	if s == s2 {
		return true
	}
	if s == "0" && s2 == "o" || s == "o" && s2 == "0" {
		return true
	}
	if s == "1" && s2 == "l" || s == "l" && s2 == "1" {
		return true
	}
	return false
}
