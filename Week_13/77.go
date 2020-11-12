package Week_13

import "fmt"

//77. 组合
//给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

func Combine() [][]int {
	var n int = 7
	var k int = 4

	return combine(n, k)
}

func combine(n int, k int) [][]int {
	var result [][]int = make([][]int, 0)
	if k < 0 || n < k {
		return result
	}
	var path []int
	backtrack2(n, k, 1, path, &result)
	return result
}

func backtrack2(n int, k int, start int, path []int, result *[][]int) {
	if len(path) == k {
		dst := make([]int, k)
		copy(dst, path)
		*result = append(*result, dst)
		return
	}
	//i <= n 的话,未剪枝,时间复杂度高
	//下面是剪枝思路,确定搜索的上界
	//1.要选择的元素个数 = k - path.size()
	//2.搜索起点的上界 = n - (k - path.size()) + 1
	var bound int = n - (k - len(path)) + 1
	for i := start; i <= bound; i++ {
		path = append(path, i)
		fmt.Println("befor", path)
		backtrack2(n, k, i+1, path, result)
		path = path[:len(path)-1]
		fmt.Println("after", path)
	}
}
