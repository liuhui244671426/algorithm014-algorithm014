package Week_03

import (
	"fmt"
	"sort"
)

var result3 [][]int

func permuteUnique(nums []int) [][]int {
	var nums2 []int
	var visited = make([]bool, len(nums)) //已访问的点
	result3 = [][]int{}
	sort.Ints(nums)                  //排个序
	backtrack3(nums, nums2, visited) //同包已有同名方法,这里改成 2
	return result3
}

func backtrack3(nums []int, nums2 []int, visited []bool) {
	//终止条件
	if len(nums) == len(nums2) {
		tmp := make([]int, len(nums))
		copy(tmp, nums2)
		fmt.Println("跳出", tmp)
		result3 = append(result3, tmp)
		return
	}
	//当前层逻辑
	for i := 0; i < len(nums); i++ {
		//进入未访问
		if visited[i] {
			continue
		}
		//基于 46 题全排列
		//这里会有重复 nums2
		if i >= 1 && !visited[i-1] && nums[i] == nums[i-1] {
			continue
		}
		//1.做选择
		//2.执行选择
		//3.撤销选择
		visited[i] = true
		nums2 = append(nums2, nums[i])
		fmt.Println("visited ", visited, nums2)
		backtrack3(nums, nums2, visited) //下一层

		nums2 = nums2[:len(nums2)-1]
		visited[i] = false

	}
}

func PermuteUnique(nums []int) [][]int {
	return permuteUnique(nums)
}
