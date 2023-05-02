package util

import (
	"math"
)

/**
 * INT value range
 * int    :
 * int32  :
 * int64  :
 * uint   :
 * uint32 :
 * uint64 : 18446744073709551615 > 10^19
 *
 *
 *
 */
/* Four arithmetic operations */

func Min(num ...int) int {
	min := math.MaxInt32
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(num ...int) int {
	max := math.MinInt32
	for _, v := range num {
		if v < max {
			max = v
		}
	}
	return max
}

func Sum(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func IsEven(num int) bool {
	return num%2 == 0
}

func QuotientAndRemainder(dividend, divider int) []int {
	r := make([]int, 2)
	r[0] = dividend / divider
	r[1] = dividend % divider
	return r
}

// abs  5 -> 5, -5 -> 5
func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Factorial 階乗
// 0: 1
// 1: 1
// 2: 2
// 3: 6
// 4: 24
// 5: 120
func Factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Factorial(n-1)
	}
}

// FactorialSlice return factorial slice up to number of n
// 10: [1, 1, 2, 6]
// 100: [1, 1, 2, 6, 24]
// 1000: [1, 1, 2, 6, 24, 120, 720]
func FactorialSlice(n int) (l []int) {
	// 桁数ごとの組み合わせ値slice
	if n == 0 {
		return
	}
	l[0] = 1
	l[1] = 1
	for i := 2; i < n; i++ {
		l[i] = l[i-1] * i
	}
	return
}

// Pow 累乗
func Pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

// Gcd 最大公約数
func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

// IsPrime O√N素数判定 return default:false
func IsPrime(n int) bool {
	if n == 1 || n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	// √N まで判定
	limit := int(math.Sqrt(float64(n)))
	r := true
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			r = false
			break
		}
	}
	return r
}

// GetPrimes x, y のレンジの素数配列を作成
func GetPrimes(x, y int) []int {
	l := make([]int, 0)
	if x == y {
		return l
	}
	// adjust start,end
	var s, e int
	if x > y {
		s = y
		e = x
	} else {
		s = x
		e = y
	}
	for s <= e {
		if IsPrime(s) {
			l = append(l, s)
		}
		s++
	}
	return l
}

// GetPrimeList indexに対応した整数値が素数かどうかの審議議をもつ配列を作成
func GetPrimeList(n int) []bool {
	// n以下の数の素数判定を行うためのスライスを用意
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	// エラトステネスのふるいを実行
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

// GetPrimeFactorials 素因数分解配列 key=素因数 val:指数を作成
// 2020 => [(2,2),(5,1),(101,1)]
func GetPrimeFactorials(x int) map[int]int {
	ll := map[int]int{}

	// 2~√N までの値で割り算し、割り切れる数値は配列へ追加
	end := int(math.Sqrt(float64(x)))
	current := 2
	for x != 1 || current <= end {
		if x%current == 0 {
			ll[current]++
			x = x / current
			end = int(math.Sqrt(float64(x)))
			continue
		}
		current++
	}
	return ll
}

// getDivisors s~targetまでの範囲で約数配列を作成
func getDivisors(s, target int) []int {
	l := make([]int, 0)
	for s < target {
		if target%s == 0 {
			l = append(l, s)
		}
		s++
	}
	return l
}
