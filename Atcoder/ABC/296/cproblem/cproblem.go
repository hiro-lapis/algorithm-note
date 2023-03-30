package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
 * https://atcoder.jp/contests/abc295/tasks/abc295_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	unique := make([]bool, 1000000000)
	ans := 0
	for i := 0; i < n; i++ {
		tmp := getNextInt(scanner) - 1
		if !unique[tmp] {
			unique[tmp] = true
		} else {
			ans++
			unique[tmp] = false
		}
	}
	fmt.Println(ans)

	defer writer.Flush()
}

// DeleteV delete slice element by searching value
func DeleteV(l []int, val int) []int {
	index := -1
	for i, v := range l {
		if v == val {
			index = i
			break
		}
	}
	if index == -1 {
		return l
	}
	return append(l[:index], l[index+1:]...)
}

func Contain(l []int, num int) (r bool) {
	for _, v := range l {
		if v == num {
			r = true
			break
		}
	}
	return
}

func Sort(l []int, o string) []int {
	if o == "desc" {
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
	} else {
		sort.Slice(l, func(i, j int) bool {
			return l[i] > l[j]
		})
	}
	return l
}
