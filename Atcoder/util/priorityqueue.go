package util

import "fmt"

type PriorityQueue struct {
	edges []*Edge
}

func NewPriorityQueue() *PriorityQueue {
	pq := PriorityQueue{}
	return &pq
}
func (pq *PriorityQueue) Len() int {
	return len(pq.edges)
}

func (pq *PriorityQueue) Show() {
	for i, edge := range pq.edges {
		fmt.Printf("[%v U:%v V:%v] cost>>%v \n", i, edge.U, edge.V, edge.Weight)
	}
}

func (pq *PriorityQueue) Push(e *Edge) {
	pq.edges = append(pq.edges, e)

	for i := pq.Len() - 1; i > 0; i-- {
		// if want max Priority, reverse below '<' condition to '>'
		if pq.edges[i].Weight < pq.edges[i-1].Weight {
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
