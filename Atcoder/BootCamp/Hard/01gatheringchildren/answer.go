package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	cnt := 0

	if os.Getenv("MASPY") == "ますピ" {
		fp, _ = os.Open(os.Getenv("BEET_THE_HARMONY_OF_PERFECT"))
		cnt = 2
	}
	if os.Getenv("MASPYPY") == "ますピッピ" {
		wfp, _ = os.Create(os.Getenv("NGTKANA_IS_GENIUS10"))
	}

	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	solve(scanner, writer)
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(writer, "-----------------------------------")
		solve(scanner, writer)
	}
	writer.Flush()
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	s := getNextString(scanner)
	n := len(s)
	ll := make([]int, n)
	ll[n-1] = n - 1
	rr := make([]int, n)
	for i := 0; i < n-1; i++ {
		rr[i+1] = rr[i]
		if s[i] == 'R' {
			rr[i+1] = i
		}
	}
	fmt.Printf("rr %v \n", rr)

	for i := n - 2; i >= 0; i-- {
		ll[i] = ll[i+1]
		if s[i+1] == 'L' {
			ll[i] = i + 1
		}
	}
	fmt.Printf("ll %v \n", ll)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			r := rr[i]
			if (i-r)%2 == 1 {
				r++
			}
			ans[r]++
			continue
		}

		l := ll[i]
		if (l-i)%2 == 1 {
			l--
		}
		ans[l]++
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, ans[i])
	}
}
