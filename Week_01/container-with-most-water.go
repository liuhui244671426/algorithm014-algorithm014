//https://leetcode-cn.com/problems/container-with-most-water/submissions/
//双指针收紧法
//盛最多水的容器
func maxArea(height []int) int {
	//fmt.Println(height, len(height))
	i, j, max_area := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			max_area = max(max_area, height[i]*(j-i))
			i++
		} else {
			max_area = max(max_area, height[j]*(j-i))
			j--
		}
	}
	return max_area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}