package Week_12

//389. 找不同

func FindTheDifference() byte {
	var s string
	var t string
	s = "abcd"
	t = "abcde"
	return findTheDifference(s, t)
}

func findTheDifference(s string, t string) byte {
	var diff byte
	var m map[byte]int = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		if _, ok := m[t[i]]; ok {
			m[s[i]]--
		}
	}

	for k, v := range m {
		if v >= 1 {
			diff = k
		}
	}
	return diff
}
