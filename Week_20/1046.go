package Week_20

import (
	"container/heap"
	"math"
)

//1046. 最后一块石头的重量
/*
思路:大顶堆时间复杂度O(nlogn)
golang大顶堆使用container/heap

*/
type IntHeap []int

func LastStoneWeight() int {
	var stones []int = []int{2, 7, 4, 1, 8, 1}
	return lastStoneWeight(stones)
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range stones {
		heap.Push(h, v)//将数组堆化
	}
	for len(*h) > 0 {
		x := heap.Pop(h)
		if x == nil {
			return 0
		}
		y := heap.Pop(h)
		if y == nil {
			return x.(int)
		}
		// fmt.Println("pop", x, y)
		z := x.(int) - y.(int)
		if z != 0 {
			heap.Push(h, int(math.Abs(float64(z))))
		}
	}

	return 0
}

//实现一个大顶堆
func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j] //大顶堆,小顶堆h[i]<h[j]
}
func (h IntHeap) Swap(i, j int) {
	//fmt.Println("swap i,j", i, j)
	if i >= 0 && j >= 0 {
		h[i], h[j] = h[j], h[i]
	}

}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	var x int = 0
	if len(*h) == 0 {
		return nil
	} else {
		x = old[len(*h)-1]
		*h = old[0 : len(*h)-1]
	}
	return x
}
