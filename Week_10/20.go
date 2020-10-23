package Week_10

//20. 有效的括号

func IsValid() bool {
	var s string
	s = "[{}]()"
	s = "([)]"
	s = "]"
	return isValid(s)
}

func isValid(s string) bool {
	var stack []rune
	if s == "" {
		return true
	}
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != m[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	//fmt.Println(stack)
	return len(stack) == 0
}
