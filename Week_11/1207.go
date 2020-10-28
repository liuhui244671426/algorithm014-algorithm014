package Week_11

import "fmt"

//1207. 独一无二的出现次数

func UniqueOccurrences() bool {
	var arr []int
	arr = []int{1, 2, 2, 1, 1, 3}
	arr = []int{1, 2}
	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	return uniqueOccurrences(arr)
}

func uniqueOccurrences(arr []int) bool {
	var set1 map[int]int = make(map[int]int)
	var set2 map[int]int = make(map[int]int)

	for i := 0; i < len(arr); i++ {
		set1[arr[i]]++
	}
	//fmt.Println(set1, set2)
	for _, v := range set1 {
		set2[v]++
		if set2[v] > 1 {
			return false
		}
	}
	return true
}
