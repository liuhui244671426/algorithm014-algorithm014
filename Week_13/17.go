package Week_13

//17. 电话号码的字母组合

func LetterCombinations() []string {
	var digits string = "23"

	return letterCombinations(digits)
}

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}
/*
1.回溯法
2.队列思路
*/ 
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}
	result := make([]string, 0)
	backtrack1(digits, "", 0, &result)
	return result
}

func backtrack1(digits string, prefix string, start int, result *[]string) {
	if start >= len(digits) {
		*result = append(*result, prefix)
		return
	}

	letter := m[digits[start]]
	for i := 0; i < len(letter); i++ {
		backtrack1(digits, prefix+letter[i], start+1, result)
	}
}
