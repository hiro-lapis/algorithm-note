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

// https://atcoder.jp/contests/abc300/tasks/abc300_b
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)
	al := make([]string, h)
	for i := 0; i < h; i++ {
		al[i] = getNextString(scanner)
	}
	bl := make([]string, h)
	for i := 0; i < h; i++ {
		bl[i] = getNextString(scanner)
	}
	exist := false
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if check(al, bl) {
				exist = true
				break
			} else {
				al = wShift(al)
			}
		}
		if check(al, bl) {
			exist = true
			break
		} else {
			al = xShift(al)
		}
	}

	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	defer writer.Flush()
}

// check hの範囲で配列同士の値が同値かチェック
func check(l, ll []string) bool {
	r := true
	h := len(l)
	for i := 0; i < h; i++ {
		if l[i] != ll[i] {
			r = false
			break
		}
	}
	return r
}

func wShift(l []string) []string {
	h := len(l)
	length := len(l[0])
	tmp := make([]string, h)
	for i := 0; i < h; i++ {
		iIndex := i
		if i-1 == h {
			iIndex = 0
		}
		prefix := string(l[i][length-1])
		body := l[i][0 : length-1]
		tmp[iIndex] = prefix + body
	}
	return tmp
}

func xShift(l []string) []string {
	h := len(l)
	tmp := make([]string, h)
	for i := 0; i < h; i++ {
		index := i + 1
		if index == h {
			index = 0
		}
		tmp[index] = l[i]
	}
	return tmp
}
