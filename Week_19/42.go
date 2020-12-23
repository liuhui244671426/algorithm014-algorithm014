package Week_19

//42. 接雨水

func Trap() int {
	var height []int = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	return trap3(height)
}

//nomarl
func trap(height []int) int {
	sum := 0
	//两边不需要考虑
	for i := 1; i < len(height)-1; i++ {
		max_left := 0
		//找出左边最大
		for j := i - 1; j >= 0; j-- {
			if height[j] > max_left {
				max_left = height[j]
			}
		}
		max_right := 0
		//找出右边最大
		for j := i + 1; j < len(height); j++ {
			if height[j] > max_right {
				max_right = height[j]
			}
		}
		//找出最小
		min := min(max_left, max_right)
		//最小高度>当前列, 才能进行计算
		if min > height[i] {
			//formula
			sum = sum + (min - height[i])
		}
	}
	return sum
}

//dp
func trap2(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft[i] = max(height[i-1], maxLeft[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(height[i+1], maxRight[i+1])
	}
	for i := 1; i < len(height)-1; i++ {
		min := min(maxLeft[i], maxRight[i])
		if min > height[i] {
			sum = sum + (min - height[i])
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//stack
func trap3(height []int) int {
	var stack []int = []int{}
	current := 0
	sum := 0
	for current < len(height) {
		//栈非空 且 当前高度 大于 栈顶高度
		for len(stack) > 0 && height[current] > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1] //栈顶元素出栈
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1            //计算距离
			min := min(height[current], height[stack[len(stack)-1]]) //计算
			sum = sum + distance*(min-top)                           //计算
		}
		stack = append(stack, current)
		current++
	}
	return sum
}
