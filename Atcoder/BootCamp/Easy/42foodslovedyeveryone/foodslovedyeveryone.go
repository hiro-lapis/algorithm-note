package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * https://atcoder.jp/contests/abc118/tasks/abc118_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	n := getNextInt(scanner)
	m := getNextInt(scanner)
	l := make([]int, m)
	for i := 1; i <= m; i++ {
		l[i-1] = i
	}

	for i := 0; i < n; i++ {
		k := getNextInt(scanner)
		tmpL := make([]int, 0)
		for j := 0; j < k; j++ {
			tmp := getNextInt(scanner)
			if Contain(l, tmp) {
				tmpL = append(tmpL, tmp)
				continue
			}
		}
		l = tmpL
	}
	fmt.Println(len(l))

	defer writer.Flush()
}

func getScanner(fp *os.File) *bufio.Scanner {
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

func Contain(list []int, target int) (r bool) {
	for _, v := range list {
		if v == target {
			r = true
			break
		}
	}
	return
}
