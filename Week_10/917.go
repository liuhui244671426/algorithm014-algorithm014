package Week_10

import (
	"fmt"
	"unicode"
)

//917. 仅仅反转字母

func ReverseOnlyLetters() string {
	var S string
	S = "a-bC-dEf-ghIj"
	//输入："a-bC-dEf-ghIj"
	//输出："j-Ih-gfE-dCba"
	return reverseOnlyLetters(S)
}

func reverseOnlyLetters(S string) string {
	//入栈法
	var stack []string
	var i = len(S) - 1
	for i >= 0 {
		if unicode.IsLetter(rune(S[i])) {
			stack = append(stack, string(S[i]))
		}
		i--
	}
	fmt.Println(S, stack)
	var newS string = ""
	for i := 0; i < len(S); i++ {
		//用IsLetter判断是否是字母
		if unicode.IsLetter(rune(S[i])) == false {
			newS += string(S[i])
		} else {
			fmt.Println(i, stack)
			if len(stack) > 0 {
				node := stack[0]
				stack = stack[1:]
				newS += node
			}
		}
	}
	return newS
}
