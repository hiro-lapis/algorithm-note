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
 * https://atcoder.jp/contests/abc298/tasks/abc298_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	al := makeGrid(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			al[i][j] = getNextInt(scanner)
		}
	}
	bl := makeGrid(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			bl[i][j] = getNextInt(scanner)
		}
	}

	same := false
	for i := 0; i < 4; i++ {
		if check(al, bl, n) {
			same = true
			break
		}
		al = rotate(n, al)
	}
	if same {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	defer writer.Flush()
}

func rotate(num int, l [][]int) [][]int {
	ll := makeGrid(num, num)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			ll[i][j] = l[num-1-j][i]
		}
	}
	return ll
}

func check(l, ll [][]int, n int) bool {
	r := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if l[i][j] == 1 && ll[i][j] != 1 {
				r = false
				break
			}
		}
	}
	return r
}

// makeGrid h*wマスのint多次元配列作成
func makeGrid(h, w int) [][]int {
	index := make([][]int, h)
	data := make([]int, h*w)
	for i := 0; i < h; i++ {
		index[i] = data[i*w : (i+1)*w]
	}
	return index
}
