package Week_13

import "fmt"

//47. 全排列 II
func PermuteUnique() [][]int {
	var nums []int
	nums = []int{
		1, 2, 1,
	}
	return permuteUnique(nums)
}

var ans [][]int

func permuteUnique(nums []int) [][]int {
	ans = make([][]int, 0)
	var path []int
	var visited map[int]bool = make(map[int]bool)
	backstrack4(nums, path, visited)
	return ans
}

func backstrack4(nums []int, path []int, visited map[int]bool) {
	if len(path) == len(nums) {
		fmt.Println("path", path)
		ans = append(ans, path)
		return
	}
	//m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}

		path = append(path, nums[i])
		visited[i] = true
		backstrack4(nums, path, visited)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
