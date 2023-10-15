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

// https://atcoder.jp/contests/abc300/tasks/abc300_c
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)
	s := make([]int, Min(h, w))
	cl := make([]string, h)
	for i := 0; i < h; i++ {
		cl[i] = getNextString(scanner)
	}

	for i := 1; i < h-1; i++ {
		for j := 1; j < w-1; j++ {
			if string(cl[i][j]) == "#" {
				size := check(cl, i, j)
				if size > 0 {
					s[size-1]++
				}
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if i >= 1 {
			fmt.Print(" ")
		}
		fmt.Print(s[i])
	}
	defer writer.Flush()
}
func Min(num ...int) int {
	min := math.MaxInt32
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}

func check(l []string, h int, w int) int {
	var size int
	H := len(l)
	W := len(l[0])
	index := 1
	for {
		// check out of range
		if h-index < 0 || index+h > H-1 || index+w > W-1 || w-index < 0 {
			break
		}
		lu := string(l[h-index][w-index])
		ru := string(l[h-index][w+index])
		ld := string(l[h+index][w-index])
		rd := string(l[h+index][w+index])
		if lu == "#" && ru == "#" && ld == "#" && rd == "#" {
			size++
			index++
		} else {
			break
		}
	}
	return size
}

// makeGrid h*wマスのint多次元配列作成
func makeGrid(h, w int) [][]string {
	index := make([][]string, h)
	data := make([]string, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}
