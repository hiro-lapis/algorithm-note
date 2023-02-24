package main

import "fmt"

/**
 * https://atcoder.jp/contests/abc121/tasks/abc121_b
 */
func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		b[i] = tmp
	}
	var ans int
	for i := 0; i < n; i++ {
		var sum, tmp int
		for j := 0; j < m; j++ {
			fmt.Scan(&tmp)
			sum += tmp * b[j]
		}
		if sum+c > 0 {
			ans += 1
		}
	}
	fmt.Println(ans)
}
