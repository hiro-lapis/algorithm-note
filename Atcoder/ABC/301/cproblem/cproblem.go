package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
 * https://atcoder.jp/contests/abc301/tasks/abc301_c
 */
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	s := getNextString(scanner)
	t := getNextString(scanner)

	// count each letter appearance
	sA := 0
	sM := CreateAlphabetMap()
	for _, v := range s {
		if string(v) != "@" {
			sM[string(v)]++
		} else {
			sA++
		}
	}

	ans := "Yes"
	tA := 0
	for _, v := range t {
		if string(v) == "@" {
			tA++
		} else {
			// decrement correspond t str count
			sM[string(v)]--
		}
	}

	sRep, tRep := 0, 0
	for i, c := range sM {
		if !checkAtcoder(i) && c != 0 {
			ans = "No"
			break
		}
		if c < 0 {
			sRep -= c
		} else {
			tRep += c
		}
	}
	// sRep is need to S replace, tRep is ...
	if sRep > sA || tRep > tA {
		ans = "No"
	}
	fmt.Println(ans)

	defer writer.Flush()
}

// CreateAlphabetMap a~zまでの英小文字をキーとし、int型の値を持つマップを作成する関数
func CreateAlphabetMap() map[string]int {
	alphabetMap := make(map[string]int)
	// a~zまでの英小文字を走査してマップに追加
	for i := 'a'; i <= 'z'; i++ {
		alphabetMap[string(i)] = 0
	}
	return alphabetMap
}

func CountAtSymbol(str string) int {
	count := strings.Count(str, "@")
	return count
}

func checkAtcoder(s string) bool {
	if s == "a" || s == "t" || s == "c" || s == "o" || s == "d" || s == "e" || s == "r" {
		return true
	}
	return false
}
