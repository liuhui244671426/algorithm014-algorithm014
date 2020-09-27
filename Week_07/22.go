package Week_07

//括号生成

func GenerateParenthesis() []string {
	var n int
	n = 2
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	result := new([]string)
	// dfs(n, 0, 0, "", result)

	backtrack(n, "", 0, 0, result)
	return *result
}

/*
深度优先
*/
func dfs(n int, left int, right int, str string, result *[]string) {
	if left == n && right == n {
		*result = append(*result, str)
		return
	}
	if left < n {
		dfs(n, left+1, right, str+"(", result)
	}
	if right < left {
		dfs(n, left, right+1, str+")", result)
	}
	return
}

/*
回溯法
*/
func backtrack(n int, cur string, left, right int, result *[]string) {
	if len(cur) == 2*n {
		*result = append(*result, cur) //选中
		return
	}
	if left < n {
		cur += "("                               //选择
		backtrack(n, cur, left+1, right, result) //回溯
		cur = cur[:len(cur)-1]                   //撤销
	}
	if right < left {
		cur += ")"
		backtrack(n, cur, left, right+1, result)
		cur = cur[:len(cur)-1]
	}
}
