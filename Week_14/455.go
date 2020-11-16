package Week_14

import "sort"

//455. 分发饼干

func FindContentChildren() int {
	var g []int //胃口值
	var s []int //尺寸
	g = []int{1, 2}
	s = []int{1, 2, 3}
	return findContentChildren(g, s)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	p1, p2 := 0, 0
	count := 0
	for p1 < len(g) && p2 < len(s) {
		if g[p1] <= s[p2] {
			count++
			p1++
		}
		p2++
	}
	
	return count
}
