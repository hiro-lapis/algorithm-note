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

/**
 * https://atcoder.jp/contests/abc297/tasks/abc297_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	s := getNextString(scanner)

	var fB, secB, kI, fR, secR int
	for i, si := range s {
		index := i + 1
		letter := string(si)
		if letter == "B" {
			// save the first B index
			if fB == 0 {
				fB = index
			} else {
				secB = index
			}
			continue
		}
		if letter == "K" {
			kI = index
			continue
		}
		if letter == "R" {
			if fR == 0 {
				fR = index
			} else {
				secR = index
			}
			continue
		}
	}
	if isSameEvenOdd(fB, secB) {
		fmt.Println("No")
		return
	}
	if !isRIncludeK(kI, fR, secR) {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")

	defer writer.Flush()
}

func isSameEvenOdd(i, j int) bool {
	if i%2 == 0 && j%2 == 0 {
		return true
	}
	if i%2 == 1 && j%2 == 1 {
		return true
	}
	return false
}

func isRIncludeK(kI, fR, secR int) bool {
	return kI > fR && secR > kI
}
