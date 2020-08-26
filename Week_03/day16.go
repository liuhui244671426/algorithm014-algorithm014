package Week_03

import "fmt"

func replaceSpace(s string) string {
	var res string
	//j := len(s) - 1
	for i := 0; i < len(s); i++ {
		if fmt.Sprintf("%c", s[i]) == " " {
			res = res + "%20"
		} else {
			res = res + fmt.Sprintf("%c", s[i])
		}
	}

	return res
}

func ReplaceSpace(s string) string {
	return replaceSpace(s)
}
