package main

import (
	"competitive-programming/Atcoder/util"
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if util.IsPrime(x) {
		fmt.Printf("%v is Prime \n", x)
	} else {
		fmt.Printf("%v is 'NOT' Prime \n", x)
	}

	pl := util.GetPrimes(1, x)
	fmt.Printf("range of 1~%v, : %v\n", x, pl)

	r := util.GetPrimeList(x)
	fmt.Printf("prime num list %v\n", r)
}
