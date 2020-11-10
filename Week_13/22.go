package Week_13

//22. 括号生成

func GenerateParenthesis() []string {
	var n int
	n = 3
	return generateParenthesis(n)
}

func generateParenthesis(n int) []string {
	result := new([]string)
	_generate(0, 0, n, "", result)
	return *result
}

func _generate(left int, right int, n int, str string, result *[]string) {
	//terminator
	if left == right && right == n {
		*result = append(*result, str)
		return
	}
	//process current level
	//dill down
	if left < n {
		_generate(left+1, right, n, str+"(", result)
	}
	if right < left {
		_generate(left, right+1, n, str+")", result)
	}

	//reverse
}
