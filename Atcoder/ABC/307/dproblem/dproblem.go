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

/*
 * https://atcoder.jp/contests/abc307/tasks/abc307_d
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	s := getNextString(scanner)

	// open&close parentheses index list
	oL := make([]int, 0)
	cL := make([]int, 0)
	for i := 0; i < n; i++ {
		if string(s[i]) == "(" {
			oL = append(oL, i)
		}
		if string(s[i]) == ")" {
			cL = append(cL, i)
		}
	}

	// 置換する文字列数だけズレをカウント
	diff := 0
	// 1番深い()から取り除いていく
	for j := len(oL) - 1; j >= 0; j-- {
		tmpI := oL[j]
		if tmpI < 0 {
			break
		}
		cI := -1
		for k := 0; k < len(cL); k++ {
			if cL[k] > tmpI {
				cI = cL[k] - diff
				break
				//cL[k] = -1
			}
		}
		if cI != -1 {
			s = s[:tmpI] + s[cI+1:]
			diff += cI - tmpI
		}
	}
	fmt.Println(s)

	defer writer.Flush()
}
