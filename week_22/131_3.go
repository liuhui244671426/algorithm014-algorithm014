package week_22

func partition3(s string) [][]string {
	ans := make([][]string, 0)
	//求所有可能性,回溯
	dfs3([]string{}, s, 0, &ans)
	return ans
}

func dfs3(temp []string, s string, start int, ans *[][]string) {
	//termina
	if start == len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*ans = append(*ans, t)
		return
	}

	for i := start; i < len(s); i++ {
		if huili3(s, start, i) {
			temp = append(temp, s[start:i+1])
			dfs3(temp, s, i+1, ans)
			temp = temp[:len(temp)-1]
		}
	}
}

func huili3(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
