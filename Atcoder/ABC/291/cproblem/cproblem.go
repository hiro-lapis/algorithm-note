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

/**
 * https://atcoder.jp/contests/abc291/tasks/abc291_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	s := getNextString(scanner)

	xy := make([]int, 2)

	ans := false
	vs := map[string]bool{"0,0": true}
	for i := 0; i < n; i++ {
		v := s[i]
		if string(v) == "R" {
			xy[0]++
		}
		if string(v) == "L" {
			xy[0]--
		}
		if string(v) == "U" {
			xy[1]++
		}
		if string(v) == "D" {
			xy[1]--
		}

		pos := strconv.Itoa(xy[0]) + "," + strconv.Itoa(xy[1])
		if _, ok := vs[pos]; !ok {
			vs[pos] = true
		} else {
			ans = true
			break
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	defer writer.Flush()
}

func NotMatch(target string, needle string) bool {
	return strings.Index(target, needle) == -1
}

func NotContain(list []string, target string) bool {
	r := true
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			r = false
			break
		}
	}
	return r
}
