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

	// receive input
	n := getNextInt(scanner)
	s := getNextString(scanner)

	t := 0
	a := 0
	for i := 0; i < n; i++ {
		win := string(s[i])
		if win == "T" {
			t++
		} else {
			a++
		}
	}
	if t > a {
		fmt.Println("T")
		return
	}
	if a > t {
		fmt.Println("A")
		return
	}

	for i := 0; i < n; i++ {
		win := string(s[i])
		if win == "T" {
			t--
		} else {
			a--
		}
		if t == 0 {
			fmt.Println("T")
			break
		}
		if a == 0 {
			fmt.Println("A")
			break
		}
	}
	defer writer.Flush()
}
