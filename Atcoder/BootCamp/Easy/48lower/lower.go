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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	n := getNextInt(scanner)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = getNextInt(scanner)
	}
	max := 0
	next := 0
	for {
		// 比較開始indexの決定
		begin := next
		if begin >= n-1 {
			break
		}
		//
		if l[begin] >= l[begin+1] {
			count := 1
			a := begin + 1
			b := begin + 2
			for {
				// 配列レンジをはみ出る時はbreak
				if b >= n {
					break
				}
				if l[a] >= l[b] {
					count++
					a++
					b++
				} else {
					break
				}
			}
			next = b
			if count > max {
				max = count
			}
		} else {
			next++
		}
	}
	fmt.Println(max)
	defer writer.Flush()
}
