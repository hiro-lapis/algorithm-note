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

/**
 * https://atcoder.jp/contests/abc301/tasks/abc301_b
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	n := getNextInt(scanner)

	a := make([]int, 0)
	for i := 0; i < n; i++ {
		tmp := getNextInt(scanner)
		a = append(a, tmp)
	}

	for {
		diffI := check(a)
		if diffI == -1 {
			break
		}
		fr := a[diffI]
		ba := a[diffI+1]
		addFlg := true
		if fr > ba {
			addFlg = false
		}
		for j := 0; j < Abs(fr-ba)-1; j++ {
			num := 0
			if addFlg {
				num = fr + 1 + j
			} else {
				num = fr - 1 - j
			}
			a = insertValue(a, diffI+j+1, num)
		}

	}

	for i, v := range a {
		fmt.Print(v)
		if i != len(a)-1 {
			fmt.Print(" ")
		}
	}

	defer writer.Flush()
}

func check(l []int) int {
	diffI := -1
	for i := 0; i < len(l)-1; i++ {
		cur := l[i]
		nex := l[i+1]
		if Abs(cur-nex) != 1 {
			diffI = i
			break
		}
	}
	return diffI
}

// Abs  5 -> 5, -5 -> 5
func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// InsertValue スライスの特定のインデックスに新しい値を挿入する関数
func InsertValue(slice []int, index, value int) []int {
	// スライスを拡張する
	slice = append(slice, 0)

	// 指定した位置の要素以降をひとつずつ後ろにずらす
	copy(slice[index+1:], slice[index:])

	// 新しい値を指定した位置に挿入する
	slice[index] = value

	return slice
}
