package main

import (
	"container/heap"
	"fmt"
)

//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

//思路：优先队列

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{key: key, count: count})
	}
	var res []int
	for len(res) < k {
		item := heap.Pop(&q).(*Item)
		res = append(res, item.key)
	}
	return res
}

type Item struct {
	key   int
	count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	//注意：因为golang中的heap是按最小堆组织的，所以count越大，Less()越小，越靠近堆顶.
	return pq[i].count > pq[j].count

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	frequent := topKFrequent(nums, k)
	fmt.Println(frequent)

}
