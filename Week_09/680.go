package Week_09

//680. 验证回文字符串 Ⅱ

func ValidPalindrome() bool {
	var s string = "abca"
	return validPalindrome(s)
}

func validPalindrome(s string) bool {
	return palindromeHelper(s, true)
}
func palindromeHelper(s string, isDelete bool) bool {
	var left, right int = 0, len(s) - 1
	for left < right {
		if s[left] != s[right] {
			if isDelete {
				return palindromeHelper(s[left:right], false) || palindromeHelper(s[left+1:right+1], false)
			} else {
				return false
			}
		}
		left++
		right--
	}
	return true
}
