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

type score struct {
	i int // index
	w int // win count
	//l string // lose count
	//d string // draw count
}

type Scores []score

func (s Scores) Len() int { return len(s) }

// func (s Scores) Swap(i, j int) int { s[i], s[j] = s[j], s[i] }
func (s Scores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Scores) Less(i, j int) bool {
	// win count is same, index asc
	if s[i].w == s[j].w {
		return s[i].i < s[j].i
	}
	return s[i].w > s[j].w
}

// https://atcoder.jp/contests/abc323/tasks/abc323_b
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	ss := make(Scores, 0)
	for i := 0; i < n; i++ {
		tmp := getNextString(scanner)
		var w int
		for _, v := range tmp {
			s := string(v)
			if s == "o" {
				w++
			}
		}
		ss = append(ss, score{i + 1, w})
	}
	sort.Sort(ss)
	for k := 0; k < len(ss); k++ {
		fmt.Println(ss[k].i)
	}
	defer writer.Flush()
}
