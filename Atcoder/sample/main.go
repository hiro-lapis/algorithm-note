package main

import "fmt"

type PriorityQueue struct {
	edges []*Edge
}

func (pq *PriorityQueue) Len() int {
	return len(pq.edges)
}

func (pq *PriorityQueue) Show() {
	for i, edge := range pq.edges {
		fmt.Printf("%v: cost>>%v \n", i, edge.weight)
	}
}

type Edge struct {
	to     string
	weight int
}

func (pq *PriorityQueue) Push(e *Edge) {
	fmt.Println(pq)
	pq.edges = append(pq.edges, e)

	for i := pq.Len() - 1; i > 0; i-- {
		if pq.edges[i].weight < pq.edges[i-1].weight {
			tmp := pq.edges[i-1]
			pq.edges[i-1] = pq.edges[i]
			pq.edges[i] = tmp
		}
	}
}

func (pq *PriorityQueue) Pop() Edge {
	popped := pq.edges[0]
	pq.edges = pq.edges[1:]
	return *popped
}

func main() {
	arr := []int{1, 100, 222}
	fmt.Println(arr[2:])

	newV := 50
	arr = append(arr, newV)
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			tmp := arr[i-1]
			arr[i-1] = arr[i]
			arr[i] = tmp
		}
	}
	fmt.Println(arr[:])

	e1 := Edge{"A", 100}
	e2 := Edge{"C", 24}
	e3 := Edge{"B", 16}
	pq := PriorityQueue{}
	pq.Push(&e1)
	pq.Push(&e2)
	pq.Push(&e3)
	pq.Show()
	fmt.Println("poppedした値")
	fmt.Println(pq.Pop().weight)
}
