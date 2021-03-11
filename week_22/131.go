package week_22

func Partition() [][]string {
	var s string = "aab"
	return partition3(s)
}

func partition(s string) [][]string {
	res := [][]string{}
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
	}
	dfs1([]string{}, 0, &res, s, memo)
	return res
}

func dfs1(temp []string, start int, res *[][]string, s string, memo [][]int) {
	if start >= len(s) {
		t := make([]string, len(temp))
		copy(t, temp)
		*res = append(*res, t)
		return
	}
	for i := start; i < len(s); i++ {
		if memo[start][i] == 2 {
			continue
		}
		if memo[start][i] == 1 || ishuili(s, start, i, memo) {
			temp = append(temp, s[start:i+1])
			dfs1(temp, i+1, res, s, memo)
			temp = temp[:len(temp)-1]
		}
	}
}

func ishuili(s string, left, right int, memo [][]int) bool {
	for left < right {
		if s[left] != s[right] {
			memo[left][right] = 2
			return false
		}
		left++
		right--
	}
	memo[left][right] = 1
	return true
}
