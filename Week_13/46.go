package Week_13

//46. 全排列

func Permute() [][]int {
	var nums []int
	nums = []int{
		1, 2, 3,
	}
	return permute(nums)
}

var result [][]int

func permute(nums []int) [][]int {
	var used map[int]bool = make(map[int]bool) //是否使用过
	var path []int
	result = make([][]int, 0)
	backstrack3(nums, path, used)
	return result
}

func backstrack3(nums []int, path []int, used map[int]bool) {
	if len(path) == len(nums) {
		//fmt.Println(path)
		dst := make([]int, len(path))
		copy(dst, path)
		result = append(result, dst)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == true {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backstrack3(nums, path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
