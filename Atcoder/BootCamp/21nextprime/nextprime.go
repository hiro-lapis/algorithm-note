package main

import "fmt"

func main() {
	var x int64
	fmt.Scan(&x)
	tmp := x
	prime(x * x)
	for {
		if !IsPrime(tmp) {
			tmp++
		}
		break
	}
	fmt.Println(tmp)
}

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

func prime(x int64) bool {
	var i int64
	for i = 2; i*i <= x; i++ {
		fmt.Printf("i is %v \n", i)
		if x%i == 0 {
			return false
		}
	}
	return true
}
