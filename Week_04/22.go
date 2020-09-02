package Week_04

//括号生成

func GenerateParenthesis() []string {
	return generateParenthesis(3)
}

func generateParenthesis(n int) []string {
	var strs []string
	dfs(0, 0, n, "", &strs)
	return strs
}

func dfs(left, right, n int, str string, strs *[]string) {
	if left == right && right == n {
		//fmt.Println(str)
		*strs = append(*strs, str)
		return
	}
	if left < n {
		dfs(left+1, right, n, str+"(", strs)
	}
	if right < left {
		dfs(left, right+1, n, str+")", strs)
	}
}
