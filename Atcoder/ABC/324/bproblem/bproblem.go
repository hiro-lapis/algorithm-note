package main

import (
	"bufio"
	"fmt"
	"math"
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

// https://atcoder.jp/contests/abc324/tasks/abc324_b
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)

	for {
		if n%2 == 0 {
			n /= 2
		} else {
			break
		}
	}
	for {
		if n%3 == 0 {
			n /= 3
		} else {
			break
		}
	}

	ans := n == 1
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	defer writer.Flush()
}

func getValue(x, y int) int {
	return int(math.Pow(2, float64(x))) * int(math.Pow(3, float64(y)))
}
