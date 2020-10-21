package Week_10

import "fmt"

//84. 柱状图中最大的矩形

func LargestRectangleArea() int {
	var heights []int
	heights = []int{2, 1, 5, 6, 2, 3}
	heights = []int{1, 1}
	return largestRectangleArea(heights)
}

/*
参考代码模板
#单调栈
stack<int> st;
for(int i = 0; i < nums.size(); i++)
{
	while(!st.empty() && st.top() > nums[i])
	{
		st.pop();
	}
	st.push(nums[i]);
}

*/
func largestRectangleArea(heights []int) int {
	var ans int = 0
	var stack []int //递增栈,存放下标
	var top int = 0 //栈顶元素
	// var left int = 0  //左边界
	// var right int = 0 //右边界,区间 (right - left + 1)
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1] //pop
			//宽度 = i - stack[end_index] - 1 ==> right-left
			ans = max(ans, (i-stack[len(stack)-1]-1)*heights[top])
		}

		stack = append(stack, i) //push,这里存放下标
	}
	fmt.Println(stack, ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
