package Week_04

// 分发饼干
import (
	"fmt"
	"sort"
)

func FindContentChildren() int {
	var g []int = []int{
		1, 2, 3,
	}
	var s []int = []int{
		1, 1,
	}
	return findContentChildren(g, s)
}

//贪心算法,1.排序 2.贪心查找
func findContentChildren(g []int, s []int) int {
	var count int = 0
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		fmt.Println(i, j)
		if g[i] <= s[j] {
			count++
			i++
		}
		j++
	}
	return count
}
