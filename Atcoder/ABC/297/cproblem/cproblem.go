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
 * https://atcoder.jp/contests/abc297/tasks/abc297_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)
	l := make([]string, h)
	for i := 0; i < h; i++ {
		si := getNextString(scanner)
		flg := false
		for j := 0; j < w; j++ {
			letter := string(si[j])
			if letter == "T" {
				if flg {
					//fmt.Printf("front >> %v \n", si[:j-1])
					//fmt.Printf("back >> %v \n", si[j+1:])
					si = si[:j-1] + "PC" + si[j+1:]
					flg = false
				} else {
					flg = true
				}
				continue
			}
			flg = false
		}
		// assign replaced string...
		l[i] = si
	}

	for _, s := range l {
		fmt.Println(s)
	}
	defer writer.Flush()
}
