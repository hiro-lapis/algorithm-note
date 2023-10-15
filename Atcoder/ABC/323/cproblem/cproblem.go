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

type q struct {
	i int
	p int
}

type Questions []q

func (q Questions) Len() int           { return len(q) }
func (q Questions) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q Questions) Less(i, j int) bool { return q[i].p > q[j].p }

// https://atcoder.jp/contests/abc323/tasks/abc323_c
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	m := getNextInt(scanner)
	// points of each problem
	aa := make(Questions, m)
	for i := 0; i < m; i++ {
		aa[i] = q{i, getNextInt(scanner)}
	}

	maxS := 0
	// score of each player
	scores := make([]string, n)
	points := make([]int, n)
	//ssx := false // same score exists
	for j := 0; j < n; j++ {
		scores[j] = getNextString(scanner)
		score := 0
		for k := 0; k < m; k++ {
			if string(scores[j][k]) == "o" {
				// add point
				score += aa[k].p
			}
		}
		points[j] = score + (j + 1) // sum with append bonus
		if points[j] > maxS {
			maxS = points[j]
		}
	}

	ans := make([]int, n)
	// sort point list asc
	sort.Sort(aa)

	for i := 0; i < n; i++ {
		// max Score player skip
		if maxS == points[i] {
			continue
		}
		pp := points[i] // player point
		ox := scores[i] // player score
		// adding not solved quesions point
		// until over max point
		for _, q := range aa {
			// already get points skip
			if string(ox[q.i]) == "o" {
				continue
			}
			// increment answer count
			ans[i]++
			pp += q.p
			if pp > maxS {
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}

	defer writer.Flush()
}
