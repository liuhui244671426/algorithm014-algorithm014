package Week_13

//78. 子集

func Subsets() [][]int {
	var nums []int = []int{
		1, 2, 3,
	}
	return subsets(nums)
}

/*
1.递归树
2.结束条件
3.选择列表
4.是否需要剪枝
5.做出选择
6.撤销选择
*/

var res [][]int //全局结果

func subsets(nums []int) [][]int {
	res = make([][]int, 0)  //最终结果
	path := make([]int, 0)  //选择列表
	res = append(res, path) //先选择

	if len(nums) == 0 {
		return res
	}

	backtrack(nums, path, 0)
	return res
}

func backtrack(nums []int, path []int, start int) {
	if start >= len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {

		path = append(path, nums[i])
		depthPath := make([]int, len(path)) //GO深浅拷贝
		copy(depthPath, path)               //GO深浅拷贝
		//fmt.Println(e)
		res = append(res, depthPath)
		backtrack(nums, path, i+1)
		path = path[:len(path)-1]
	}
}
