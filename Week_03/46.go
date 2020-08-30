package Week_03

import "fmt"

var result [][]int

func permute(nums []int) [][]int {
	var nums2 []int
	var visited = make([]bool, len(nums)) //已访问的点
	result = [][]int{}
	backtrack2(nums, nums2, visited) //同包已有同名方法,这里改成 2
	return result
}

func backtrack2(nums []int, nums2 []int, visited []bool) {
	//终止条件
	if len(nums) == len(nums2) {
		tmp := make([]int, len(nums))
		copy(tmp, nums2)
		fmt.Println("跳出", tmp)
		result = append(result, tmp)
		return
	}
	//当前层逻辑
	for i := 0; i < len(nums); i++ {
		//进入未访问
		if !visited[i] {
			//1.做选择
			//2.执行选择
			//3.撤销选择
			visited[i] = true
			nums2 = append(nums2, nums[i]) 
			backtrack2(nums, nums2, visited) //下一层
			
			nums2 = nums2[:len(nums2)-1]
			visited[i] = false
		}
	}
}

func Permute(nums []int) [][]int {
	return permute(nums)
}
