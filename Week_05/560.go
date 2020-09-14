package Week_05

import "fmt"

//

func SubarraySum() int {
	var nums []int = []int{-1, -1, 1}
	var k int = 0
	return subarraySum(nums, k)
}

//前缀和
//重点:[pre-k]
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1 //对于下标为 0 的元素，前缀和为 0，个数为 1
	fmt.Println("pre, count, pre-k, m")
	for i := 0; i < len(nums); i++ {
		pre += nums[i]

		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
		fmt.Println(pre, count, pre-k, m)
	}
	return count
}
