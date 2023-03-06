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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	num := n * 5
	l := make([]int, num)
	for i := 0; i < num; i++ {
		l[i] = getNextInt(scanner)
	}

	l = Sort(l, "desc")
	l = l[n:]
	l = l[:len(l)-n]

	ans := 0.0
	for _, v := range l {
		ans += float64(v)
	}
	ans = ans / float64(len(l))
	fmt.Println(ans)
	defer writer.Flush()
}

func Sort(l []int, o string) []int {
	if o == "desc" {
		sort.Slice(l, func(i, j int) bool {
			return l[i] > l[j]
		})
	} else {
		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
	}
	return l
}
