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

type hub struct {
	w int
	x int
}
type hubs []hub

func (h hubs) Len() int      { return len(h) }
func (h hubs) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h hubs) Less(i, j int) bool {
	return h[i].x > h[j].x
}

// https://atcoder.jp/contests/abc325/tasks/abc325_b
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	hh := make(hubs, 0)
	for i := 0; i < n; i++ {
		hh = append(hh, hub{getNextInt(scanner), getNextInt(scanner)})
	}
	max := 0
	tmp := 0
	for i := 0; i < 24; i++ {
		tmp = 0
		for j := 0; j < len(hh); j++ {
			destTime := hh[j].x + i
			destTime = format(destTime)
			if destTime >= 9 && 18 > destTime {
				tmp += hh[j].w
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	fmt.Println(max)

	defer writer.Flush()
}

func format(t int) int {
	if t < 24 {
		return t
	}
	return t - 24
}
