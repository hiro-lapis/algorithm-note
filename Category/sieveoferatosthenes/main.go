package main

import "fmt"

func main() {
	var x int64
	fmt.Scan(&x)
	tmp := x
	Next(x * x)
	for {
		if !IsPrime(tmp) {
			tmp++
		}
		break
	}
	fmt.Println(tmp)
}

// IsPrime
func IsPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	r := true
	for i := n / 2; i > 2; i-- {
		if n%i == 0 {
			r = false
			break
		}
	}

	return r
}

func Next(x int64) bool {
	var i int64
	for i = 2; i*i <= x; i++ {
		fmt.Printf("i is %v \n", i)
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Eratosthenes(n int64) {
	prime := make([]bool, n+1)
	prime[0] = true
	prime[1] = true
	for i := 2; i <= int(n); i++ {
		if !prime[i] {
			continue
		}
		for j := 0; j < int(n); j += i {
			prime[j] = false
		}
	}
	DisplayPrime(int(n), prime)
}

func DisplayPrime(n int, prime []bool) {
	fmt.Printf("Primes less than %d : ", n)
	for i := 2; i <= n; i++ {
		if prime[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
