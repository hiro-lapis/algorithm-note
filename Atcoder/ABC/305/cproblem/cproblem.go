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
 * https://atcoder.jp/contests/abc305/tasks/abc305_c
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
		l[i] = getNextString(scanner)
	}

	hcl := make([]int, h)
	wcl := make([]int, w)

	for j := 0; j < h; j++ {
		str := l[j]
		for k := 0; k < w; k++ {
			if string(str[k]) == "#" {
				wcl[k]++
				hcl[j]++
			}
		}
	}

	var bef int
	ans := make([]int, 2)
	for i, v := range hcl {
		if v == 0 {
			continue
		}
		cur := v
		if bef != 0 {
			if bef != cur {
				if bef > cur {
					ans[0] = i + 1
				} else {
					ans[0] = i
				}
				break
			}
		} else {
			bef = cur
		}
	}
	var bef2 int
	for ii, vv := range wcl {
		if vv == 0 {
			continue
		}
		cur := vv
		if bef2 != 0 {
			if bef2 != cur {
				if bef2 > cur {
					ans[1] = ii + 1
				} else {
					ans[1] = ii
				}
				break
			}
		} else {
			bef2 = cur
		}
	}
	fmt.Println(ans[0])
	fmt.Println(ans[1])

	defer writer.Flush()
}
