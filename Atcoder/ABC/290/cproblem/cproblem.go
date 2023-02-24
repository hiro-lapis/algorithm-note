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

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	n := getNextInt(scanner)
	k := getNextInt(scanner)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = getNextInt(scanner)
	}

	// 各数値の存在を示すmapを作成
	// length=0 で、追加に応じて拡張
	m := map[int]bool{}
	for i := 0; i < n; i++ {
		m[l[i]] = true
	}

	/**
	 * 7 3
	 * 2 0 2 3 2 1 9
	 * => map[0:true 1:true 2:true 3:true 9:true]
	 * fmt.Println(m[7]) => false
	 */
	ans := 0
	// MEXの開始値は0だが、MEX最大値はN+1なのでそこまでの範囲を探索
	for i := 0; i < n+1 && k > 0; i++ {
		// 探している値がある場合
		if m[i] {
			ans++
			k--
			continue
		}
		fmt.Println(i)
		return
	}
	fmt.Println(ans)
	defer writer.Flush()
}
