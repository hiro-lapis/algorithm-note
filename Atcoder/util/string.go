package util

import (
	"regexp"
	"strconv"
	"strings"
)

// SumAscii 文字列の数値として合算値を取得
func SumAscii(s string) int {
	var sum int32
	// rune
	for _, v := range s {
		sum += v
	}
	return int(sum)
}

// StrSearch check string from slice of string
func StrSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}

// Reverse 文字列の左右反転
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsDigit check if s is numeric string
func IsDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
	// another variation...
	integerRegexp := regexp.MustCompile(`^-?\d+$`)
	return integerRegexp.MatchString(s)
}

// Split 文字分割
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// SearchIndex needleが出現するindexを出力 0:1文字目 -1:無し
func SearchIndex(target, needle string) int {
	return strings.Index(target, needle)
}

// CountTarget count str appearance count
func CountTarget(str string, t string) int {
	count := strings.Count(str, t)
	return count
}

// IsLowerCase 小文字のみか判定
func isLowerCase(str string) bool {
	for _, char := range str {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}

// IsPalindrome 回文か判定
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {

			return false
		}
		start++
		end--
	}
	return true
}

// StrJoin 文字列の結合
func StrJoin(s ...string) string {
	var r string
	for _, v := range s {
		r += v
	}
	return r
}

// GetLength 文字数の取得
func GetLength(s ...string) int {
	var r int
	for _, v := range s {
		r += len(v)
	}
	return r
}
