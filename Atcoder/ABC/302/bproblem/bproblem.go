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

/**
 * https://atcoder.jp/contests/abc302/tasks/abc302_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	defer writer.Flush()

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)

	l := make([]string, h)
	for i := 0; i < h; i++ {
		l[i] = getNextString(scanner)
	}

	target := "snuke" // target
	ans := make([]int, 2)
	r := false
	find := false

	// horizon
	for i := 0; i < h; i++ {
		for j := 0; j < w-4; j++ {
			str := l[i][j : j+5]
			if str == target || ReverseString(str) == target {
				ans[0] = i + 1
				ans[1] = j + 1
				find = true
				if ReverseString(str) == target {
					r = true
				}
				break
			}
		}
		if find {
			break
		}
	}

	if find {
		for i := 0; i < 5; i++ {
			if !r {
				fmt.Printf("%d %d\n", ans[0], ans[1]+i)
			} else {
				fmt.Printf("%d %d\n", ans[0], ans[1]+4-i)
			}
		}
		return
	}

	// vert
	for i := 0; i < w; i++ {
		for j := 0; j < h-4; j++ {
			str := string(l[j][i]) + string(l[j+1][i]) + string(l[j+2][i]) + string(l[j+3][i]) + string(l[j+4][i])
			if str == target || ReverseString(str) == target {
				ans[0] = j + 1
				ans[1] = i + 1
				find = true
				if ReverseString(str) == target {
					r = true
				}
				break
			}
		}
		if find {
			break
		}
	}

	if find {
		for i := 0; i < 5; i++ {
			if !r {
				fmt.Printf("%d %d\n", ans[0]+i, ans[1])
			} else {
				fmt.Printf("%d %d\n", ans[0]+4-i, ans[1])
			}
		}
		return
	}

	left := true
	for i := 0; i < h-4; i++ {
		for j := 0; j < w; j++ {
			var strR, strL string

			if j-4 >= 0 {
				strL = string(l[i][j]) + string(l[i-1][j-1]) + string(l[i-2][j-2]) + string(l[i-3][j-3]) + string(l[i-4][j-4])
			}
			if j+4 <= w {
				strR = string(l[i][j]) + string(l[i+1][j+1]) + string(l[i+2][j+2]) + string(l[i+3][j+3]) + string(l[i+4][j+4])
			}
			if strL == target || ReverseString(strL) == target {
				ans[0] = i + 1
				ans[1] = j + 1
				find = true
				if ReverseString(strL) == target {

					r = true
				}
				break
			}
			if strR == target || ReverseString(strR) == target {
				left = false
				ans[0] = i + 1
				ans[1] = j + 1
				find = true
				if ReverseString(strR) == target {
					r = true
				}
				break
			}
		}
		if find {
			break
		}
	}
	for i := 0; i < 5; i++ {
		hh := i
		if r {
			hh = 4 - i
		}
		ww := i
		if !left {
			ww = 4 - i
		}
		fmt.Printf("%d %d\n", ans[0]+hh, ans[1]+ww)
	}
}

func ReverseString(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
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
