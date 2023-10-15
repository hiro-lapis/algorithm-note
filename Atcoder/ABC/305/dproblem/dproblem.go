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
 * https://atcoder.jp/contests/abc305/tasks/abc305_d
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = getNextInt(scanner)
	}

	//各区切りごとの累計睡眠時間
	tl := make([]int, n+1)
	for i := 0; i+1 < n; i++ {
		if i%2 == 1 {
			//　起床 - 就寝時間を足す
			tl[i+1] += a[i+1] - a[i]
		}
		// 奇数回では起きてるので手前の値をそのまま入れる
		tl[i+1] += tl[i]
	}

	q := getNextInt(scanner)
	for j := 0; j < q; j++ {
		l := getNextInt(scanner)
		r := getNextInt(scanner)

		// 睡眠開始時間　値がstart以上の最小の値のindexを取得(upper bound)
		ll := sort.SearchInts(a, l)
		// 起床時間　値がend以上の最小の値のindexを取得(upper bound)
		rr := sort.SearchInts(a, r)
		//　期間での累積睡眠時間算出(起床時点での累積睡眠時間 - 就寝時点の累積就寝時間)
		ans := tl[rr] - tl[ll]

		/*
		 * upper bound index がevenの時 = 起床時間のindexを参照している
		 * upper bound で省かれている時間の足しひきを行う
		 *
		 * ll => 計測開始 ~ 起床時間までを +
		 * rr => 計測開始 ~ 起床時間までを -
		 */
		if ll%2 == 0 {
			ans += a[ll] - l
		}
		if rr%2 == 0 {
			ans -= a[rr] - r
		}

		fmt.Println(ans)
	}

	defer writer.Flush()
}
