package main

func main() {
	//var n int
	//fmt.Scan(&n)
	//r := Fibonacci(n)
	//fmt.Println(r)
	//r2 := Fibonacci2(n)
	//fmt.Println(r2)

	//var x, y int
	//fmt.Scan(&x, &y)
	//z := Gcd(x, y)
	//fmt.Printf("x:%d y:%d gcd >>> %d \n", x, y, z)
}

// Fibonacci return length of fibonacci slice
func Fibonacci(length int) []int {
	l := make([]int, length)
	l[0] = 1
	l[1] = 1
	for i := 2; i < length; i++ {
		l[i] = l[i-1] + l[i-2]
	}
	return l
}

// Fibonacci2 return up to n fibonacci slice
func Fibonacci2(n int) []int {
	l := []int{1, 1}
	for i := 2; l[i-1]+l[i-2] < n; i++ {
		l = append(l, l[i-1]+l[i-2])
	}
	return l
}

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	remainder := x % y
	return Gcd(y, remainder)
}
