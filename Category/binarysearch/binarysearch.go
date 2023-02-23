package main

import "fmt"

func main() {
	var target int
	fmt.Println("plz input target int between 1 to 100...")
	fmt.Scan(&target)
	l := make([]int, 100)
	c := 0

	for i := 0; i < 100; i++ {
		l[i] = i + 1
	}
	fmt.Println(l)
	r := BinarySearch(target, l, &c)

	fmt.Printf("target: %d search count %d", r, c)
}

func BinarySearch(target int, list []int, count *int) int {
	if len(list) == 1 {
		return -1
	}

	*count++
	next := len(list)/2 - 1
	fmt.Printf("no.%d list val is %d target> %d \n", *count, list[next], target)
	if list[next] == target {
		return list[next]
	}
	if list[next] > target {
		return BinarySearch(target, list[:next], count)
	} else {
		return BinarySearch(target, list[next:], count)
	}
}

func BinarySearch2(target int, list []int, low int, high int, count *int) int {
	if low > high {
		return -1
	}

	*count++
	next := low + high/2
	fmt.Printf("no.%d list val is %d target> %d \n", *count, list[next], target)
	if list[next] == target {
		return list[next]
	}
	if list[next] > target {
		return BinarySearch2(target, list, next+1, high, count)
	} else {
		return BinarySearch2(target, list, low, next-1, count)
	}
}
