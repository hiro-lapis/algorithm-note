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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	m := getNextInt(scanner)
	s := make([]string, 3)
	for i := 0; i < 3; i++ {
		s[i] = getNextString(scanner)
	}

	const INF = math.MaxInt
	ans := INF

	// i=reel1, j=reel2, k-reel3
	for i := 0; i < m*3; i++ {
		for j := 0; j < m*3; j++ {
			// 同じ秒数にリールを同時押しはできないためスキップ
			if i == j {
				continue
			}
			for k := 0; k < m*3; k++ {
				// 同じ秒数にリールを同時押しはできないためスキップ
				if i == k || j == k {
					continue
				}
				// 1,2,3 それぞれがあっている場合
				if s[0][i%m] == s[1][j%m] && s[0][i%m] == s[2][k%m] {
					// 各種リールが揃った秒数のうち最小の値と、それまでの計算での最小値を比較、小さい方を解答値にする
					ans = Min(ans, Max(i, Max(j, k)))
				}
			}
		}
	}
	if ans == INF {
		ans = -1
	}
	fmt.Println(ans)
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

func Max(num ...int) int {
	max := math.MinInt32
	for _, v := range num {
		if v > max {
			max = v
		}
	}
	return max
}
