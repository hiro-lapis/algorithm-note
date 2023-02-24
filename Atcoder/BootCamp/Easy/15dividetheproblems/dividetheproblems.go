package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

	d := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		d[i] = tmp
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] < d[j]
	})
	ans := Divide(d)
	fmt.Println(ans)
	writer.Flush()
}

func Divide(list []int) int {
	ans := 0
	mid := len(list) / 2
	front := list[mid-1]
	back := list[mid]
	if front == back {
		return ans
	}
	return back - front
}
