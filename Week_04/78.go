package Week_04

func Subsets(nums []int) [][]int {
	return subsets(nums)
}

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	var result [][]int //全局变量

	backtrack(0, nums, path, &result)
	return result
}

func backtrack(cur int, nums []int, tmp []int, result *[][]int) {
	item := make([]int, len(tmp))
	copy(item, tmp) //复制全新一份
	*result = append(*result, item)

	for j := cur; j < len(nums); j++ {
		tmp = append(tmp, nums[j])        //做选择
		backtrack(j+1, nums, tmp, result) //下一个梦境
		tmp = tmp[:len(tmp)-1]            //撤销选择
	}
}
