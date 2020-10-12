package Week_08

//242. 有效的字母异位词

func IsAnagram() bool {
	var s string
	var t string

	s = "anagram"
	t = "nagaras"
	return isAnagram(s, t)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m map[byte]int = make(map[byte]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
