package Week_10

import (
	"fmt"
)

//151. 翻转字符串里的单词
//todo  时间复杂度较低,需要优化

func ReverseWords() string {
	var s string
	// 输入："the sky is blue"
	// 输出："blue is sky the"
	s = "     the sky is blue    "
	s = "a good   example"
	return reverseWords(s)
}

func reverseWords(s string) string {

	s = trim(s)
	var j int = len(s)
	var i int = len(s) - 1
	var res []string
	for i >= 0 {
		//fmt.Println("i, len", j, i, len(s), string(s[i]))
		if s[i] == 32 { //空格 => byte->32
			res = append(res, trim(s[i:j]))
			j = i
		}

		if i == 0 {
			res = append(res, trim(s[:j]))
			j = i
		}
		i--
	}
	fmt.Println("res: ", res)
	var ss string = ""
	for k := 0; k < len(res); k++ {
		if res[k] != " " {
			ss = ss + res[k] + " "
		}
	}
	//var ss string = strings.Join(res, " ")
	return trim(ss)
}

func trim(s string) string {
	//fmt.Println("origin s: ", s)

	j := 0
	for j < len(s) {
		if s[j] != 32 {
			s = s[j:]
			break
		}
		j++
	}
	//fmt.Println(s)
	//fmt.Println(len(s))
	i := len(s) - 1

	for i > 0 {
		if s[i] != 32 {
			//fmt.Println(string(s[i]), i)
			s = s[:i+1]
			break
		}
		i--
	}

	//fmt.Println("trim s: ", s)
	return s
}
