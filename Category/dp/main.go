package main

import "fmt"

func main() {
	const scale = 1_000_000_000 + 7
	const t = "chokudai"
	var s string
	fmt.Scan(&s)
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j <= len(t); j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if s[i-1] != t[j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % scale
			}
		}
	}
	fmt.Println(dp[len(dp)-1][8])
}
