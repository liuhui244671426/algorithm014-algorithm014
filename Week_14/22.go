package Week_14

//22. 括号生成
//还能用动态规划做.
func GenerateParenthesis() []string {
	var n int = 3
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	_generate(n, 0, 0, "", &result)
	return result
}

func _generate(n int, left int, right int, path string, result *[]string) {
	if n == left && n == right {
		//fmt.Println(path)
		*result = append(*result, path)
		return
	}
	//剪枝,效果很不错
	if left < right {
		return
	}
	if left < n {
		_generate(n, left+1, right, path+"(", result)
	}
	if right < left {
		_generate(n, left, right+1, path+")", result)
	}
}
