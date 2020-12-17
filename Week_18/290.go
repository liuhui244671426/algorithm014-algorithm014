package Week_18

import (
	"fmt"
	"strings"
)

//290. 单词规律

func WordPattern() bool {
	var pattern = "abba"
	var str = "dog cat cat dog"
	str = "dog dog dog dog"
	return wordPattern(pattern, str)
}

func wordPattern(pattern string, s string) bool {
	h := make(map[byte]string)
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		//h[p]
		if val, ok := h[p]; ok {
			if val != ss[i] {
				return false
			}
		} else {
			//!h[p]
			for _, v := range h {
				if v == ss[i] {
					return false
				}
			}
			h[p] = ss[i]
		}
	}

	// fmt.Println(h)
	return true
}

func mySplit(s string) []string {
	var word string = ""
	start := 0
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' {
			word = s[start:j]
			start = j + 1
		} else if j == len(s)-1 {
			word = s[start : j+1]
		}

		fmt.Println(word)
		// if m[string(ch)] != word {
		// 	return false
		// }
	}
	return []string{}
}
