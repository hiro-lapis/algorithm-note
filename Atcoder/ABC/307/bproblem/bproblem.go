package main

import (
	"bufio"
	"fmt"
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

/*
 * https://atcoder.jp/contests/abc307/tasks/abc307_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = getNextString(scanner)
	}

	ans := false
	for i := 0; i < n; i++ {
		ii := s[i]
		for j := i + 1; j < n; j++ {
			if j > n-1 {
				break
			}
			jj := s[j]
			if IsPalindrome(ii+jj) || IsPalindrome(jj+ii) {
				ans = true
				break
			} else {
				ans = false
			}

		}
		if ans {
			break
		}
	}
	if ans {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")

	defer writer.Flush()
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {

			return false
		}
		start++
		end--
	}
	return true
}
