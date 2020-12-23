package Week_19

//387. 字符串中的第一个唯一字符

func FirstUniqChar() int {
	var s string = "loveleetcode" //2
	return firstUniqChar2(s)
}

//map
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	m := make(map[int32]int)
	for _, char := range s {
		m[char]++
	}
	//fmt.Println(m)
	for index, char := range s {
		if m[char] == 1 {
			return index
		}
	}

	return -1
}

//array ,array时间复杂度和空间复杂度都优于 map
func firstUniqChar2(s string) int {
	if len(s) == 0 {
		return -1
	}
	var arr [26]int32
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	// fmt.Println(arr)

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
