package main

import "fmt"

func main() {
	const t = "chokudai"
	const mod = 1_000_000_000 + 7
	var s string
	fmt.Scan(&s)
	n := len(s)

	// 二次元配列を作成(縦横それぞれ1マス多めに作る)
	// 縦:s 横:t ＊どちらが縦横かで漸化式適用時の参照マスが変わるので注意!
	dp := make([][]int, len(s)+1) // 縦の長さs
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, len(t)+1)
		// dpの作法として左列を1埋め
		tmp[0] = 1
		dp[i] = tmp
	}

	// マス目はt+1で作ったがfor文での探索は0~t-1でOK
	// マス目はn+1で作ったがfor文での探索は0~nでOK
	for y := 0; y < n; y++ {
		for x := 0; x < len(t); x++ {
			// 縦:s 横:tなのでそれぞれのindexで比較
			if s[y] != t[x] {
				// 漸化式(文字が合わない場合)
				//　1埋めによる文字のindex(1文字目:0)と配列上のindex(1つ目:1)の対応関係に注意して値を入れる
				dp[y+1][x+1] = dp[y][x+1]
			} else {
				// 漸化式(文字が合う場合)
				dp[y+1][x+1] = (dp[y][x+1] + dp[y][x]) % mod
			}
		}
	}
	// 最大値は右下にあるので出力
	//    chokudai
	// 0:100000000
	// 1:110000000 c
	// 2:111000000 ch
	// 3:121000000 chc (真上の1を入れて1にする)
	// 4:123000000 chch (真上の1+左上の2=3)
	// 5:123000000 chcho
	// 6:123300000 chchok
	// 7:123333000 chchoku
	// 8:123333300 chchokud
	// 9:123333330 chchokuda
	//10:123333333 chchokudai
	fmt.Println(dp[n][8])
}

//func main() {
//	const scale = 1_000_000_000 + 7
//	const t = "chokudai"
//	var s string
//	fmt.Scan(&s)
//	dp := make([][]int, len(s)+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, len(t)+1)
//	}
//
//	for i := 0; i < len(dp); i++ {
//		for j := 0; j <= len(t); j++ {
//			if j == 0 {
//				dp[i][j] = 1
//			} else if i == 0 {
//				dp[i][j] = 0
//			} else if s[i-1] != t[j-1] {
//				dp[i][j] = dp[i-1][j]
//			} else {
//				dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % scale
//			}
//		}
//	}
//	fmt.Println(dp[len(dp)-1][8])
//}
