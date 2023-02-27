package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getScanner(fp *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
	return scanner
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	var n, k, q int
	fmt.Scan(&n, &k, &q)

	a := make([]int, n)
	for j := 0; j < q; j++ {
		// key毎に加点値を入れる
		a[getNextInt(scanner)-1]++
	}

	// 点数 = 初期点数(K) - a[i]
	// if 点数 > 0

	for _, v := range a {
		// 初期点 + 加減点(加点+減点で最大0)
		ans := k + v - q
		//fmt.Printf("ans>> %v \n", ans)
		if ans > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
	writer.Flush()
}
