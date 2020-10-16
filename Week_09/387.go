package Week_09

//387. 字符串中的第一个唯一字符

func FirstUniqChar() int {
	var s string
	s = "loveleetcode"
	return firstUniqChar(s)
}

func firstUniqChar(s string) int {
	var m map[byte]int = make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
