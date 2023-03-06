package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/**
 * https://atcoder.jp/contests/abc134/tasks/abc134_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	// 入力受付処理
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	defer writer.Flush()
	n := getNextInt(scanner)

	l := map[int]int{}
	for i := 1; i <= n; i++ {
		l[i] = getNextInt(scanner)
	}
	// mapの降順に参照するkeyを取得
	keys := mapSortedKeys(l, false)

	// j=各ループで除外したいindex
	for j := 1; j <= n; j++ {
		// 各除外indexのループ内で、値の降順にmapを参照していく
		for k := 0; k < len(keys); k++ {
			mapIdx := keys[k]
			if j == mapIdx {
				mapIdx = keys[k+1]
			}
			fmt.Println(l[mapIdx])
			//値の降順は、最初orその次のindexが解答になるので1回目でループを抜ける
			break
		}
	}
}

func getScanner(fp *os.File) *bufio.Scanner {
	// バッファinput/outputから入力スキャナ作成
	scanner := bufio.NewScanner(fp)
	// スキャナの入力分割単位を、単語単位分割(デフォルトはnew line)に設定
	scanner.Split(bufio.ScanWords)
	// バッファサイズを設定
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

// mapSortedKeys map INT値の昇順に参照するkey配列を取得 asc/desc指定必須
func mapSortedKeys(list map[int]int, asc bool) []int {
	keys := make([]int, 0, len(list))
	for i := range list {
		keys = append(keys, i)
	}

	if asc {
		sort.SliceStable(keys, func(i, j int) bool {
			return list[keys[i]] < list[keys[j]]
		})
	} else {
		sort.SliceStable(keys, func(i, j int) bool {
			return list[keys[i]] > list[keys[j]]
		})
	}
	return keys
}
