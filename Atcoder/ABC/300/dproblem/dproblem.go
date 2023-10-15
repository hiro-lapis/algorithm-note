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
	ans := solve1(n)

	fmt.Println(ans)

	defer writer.Flush()
}

func solve1(n int) int {
	const MAX = 300001
	pl := make([]int, MAX)
	dl := make([]int, 0)
	for i := 2; i < MAX; i++ {
		if pl[i] > 0 {
			continue
		}
		pl[i] = i
		for j := i; j < MAX; j += i {
			if pl[j] > 0 {
				continue
			}
			pl[j] = i
		}
	}

	for i := 2; i < MAX; i++ {
		if pl[i] == i {
			dl = append(dl, i)
		}
	}
	ans := 0
	for i := 0; i < len(dl) && dl[i] < n; i++ {
		for j := 0; j < i && dl[i]*dl[j]*dl[j] < n; j++ {
			k := sort.Search(len(dl), func(x int) bool {
				return n/(dl[i]*dl[j]*dl[j]) < dl[x]*dl[x]
			})

			if k-i-1 < 0 {
				continue
			}
			ans += k - i - 1
		}
	}
	return ans
}
