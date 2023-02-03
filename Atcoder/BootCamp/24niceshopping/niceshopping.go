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
 * https://atcoder.jp/contests/hitachi2020/tasks/hitachi2020_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	//var x, y, m int
	//fmt.Scan(&x, &y, &m)
	x := getNextInt(scanner)
	y := getNextInt(scanner)
	m := getNextInt(scanner)

	xx := make([]int, x)
	xMin := math.MaxInt32
	for i := 0; i < x; i++ {
		//var tmp int
		//fmt.Scan(&tmp)
		tmp := getNextInt(scanner)
		xx[i] = tmp
		if xMin > tmp {
			xMin = tmp
		}
	}

	yy := make([]int, y)
	yMin := math.MaxInt32
	for j := 0; j < y; j++ {
		//var tmp int
		//fmt.Scan(&tmp)
		tmp := getNextInt(scanner)
		yy[j] = tmp
		if yMin > tmp {
			yMin = tmp
		}
	}
	ans := xMin + yMin
	for k := 0; k < m; k++ {
		//var a, b, discount int
		//fmt.Scan(&a, &b, &discount)
		a := getNextInt(scanner)
		b := getNextInt(scanner)
		discount := getNextInt(scanner)
		cost := xx[a-1] + yy[b-1] - discount
		if ans > cost {
			ans = cost
		}
	}
	fmt.Println(ans)

	writer.Flush()
}
