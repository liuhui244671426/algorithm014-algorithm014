package week_22

func partition2(s string) [][]string {

	ans := make([][]string, 0)
	dfs2([]string{}, s, 0, &ans)
	return ans
}

func dfs2(temp []string, s string, start int, ans *[][]string) {
	if start >= len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*ans = append(*ans, t)
		return
	}

	for i := start; i < len(s); i++ {
		if ishuili2(s, start, i) {
			temp = append(temp, s[start:i+1])
			dfs2(temp, s, i+1, ans)
			temp = temp[:len(temp)-1]
		}
		//选择
		//进入
		//撤销
	}
}

func ishuili2(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
