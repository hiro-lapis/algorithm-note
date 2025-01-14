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

type mp struct {
	i int
	j int
}
type mps []mp

func (m mps) Len() int      { return len(m) }
func (m mps) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m mps) Less(i, j int) bool {
	if m[i].i == m[j].i {
		return m[i].j > m[j].j
	}
	return m[i].i > m[j].i
}

// https://atcoder.jp/contests/abc325/tasks/abc325_c
func main() {
	fp := os.Stdin
	wfp := os.Stdout
	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)

	// receive input
	h := getNextInt(scanner)
	w := getNextInt(scanner)
	ss := make(mps, 0)
	for i := 0; i < h; i++ {
		tmp := getNextString(scanner)
		for j := 0; j < w; j++ {
			str := string(tmp[j])
			if str == "#" {
				ss = append(ss, mp{i + 1, j + 1})
			}
		}
	}

	vs := make([]mps, 0)
	for i2 := 0; i2 < len(ss); i2++ {
		m := ss[i2]
		c, adjI := check(vs, m)
		if c {
			continue
		}
		if adjI == -1 {
			mm := make(mps, 0)
			mm = append(mm, m)
			vs = append(vs, mm)
		} else {
			vs[adjI] = append(vs[adjI], m)
		}
	}
	for k := 0; k < len(vs); k++ {
		vv := vs[k]
		index := k + 1
		if k == len(vs)-1 {
			index = k - 1
		}
		vv2 := vs[index]
		if check2(vv, vv2) {
			vs[k] = append(vs[k], vs[index]...)
			vs[index] = []mp{}
		}
	}
	ans := 0
	for _, v := range vs {
		if len(v) > 0 {
			ans++
		}
	}

	fmt.Println(ans)
	defer writer.Flush()
}

func check(vs []mps, m mp) (bool, int) {
	adjI := -1
	c := false
	for i, mm := range vs {
		for _, mp := range mm {
			if mp.i == m.i && mp.j == m.j {
				c = true
				adjI = -1
				break
			} else if 2 > Abs(mp.i-m.i) && 2 > Abs(mp.j-m.j) {
				adjI = i
			}
		}
		if c || adjI != -1 {
			break
		}
	}
	return c, adjI
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func check2(v, v2 []mp) bool {
	r := false
	for i := 0; i < len(v); i++ {
		vv := v[i]
		for j := 0; j < len(v2); j++ {
			vv2 := v2[j]
			if 2 > Abs(vv.i-vv2.i) && 2 > Abs(vv.j-vv2.j) {
				r = true
				break
			}
		}
		if r {
			break
		}
	}
	return r
}
