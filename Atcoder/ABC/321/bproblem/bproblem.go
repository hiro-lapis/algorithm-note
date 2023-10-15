package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// https://atcoder.jp/contests/abc321/tasks/abc321_b
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	x := getNextInt(scanner)
	aa := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		aa[i] = getNextInt(scanner)
	}

	aa = Sort(aa, "asc")
	bb := Copy(aa)
	ans := -1
	for last := 0; last < 101; last++ {
		bb = append(bb, last)
		bb = Sort(bb, "asc")
		if Sum(bb) > x {
			ans = last
			break
		}
	}
	fmt.Println(ans)
	defer writer.Flush()
}

func Sort(l []int, o string) []int {
	if o == "desc" {
		sort.Slice(l, func(i, j int) bool {
			return l[i] > l[j]
		})
	} else {
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
	}
	return l
}

func Copy(l []int) (r []int) {
	Show(r)
	for i := 0; i < len(l); i++ {
		r[i] = l[i]
	}
	return
}

func Show(l []int) {
	fmt.Printf("len=%d cap=%d", len(l), cap(l))
}

func Sum(l []int) (r int) {
	for i := 0; i < len(l); i++ {
		r += l[i]
	}
	return
}
