package Week_09

//5. 最长回文子串

func LongestPalindrome() string {
	var s string
	s = "caba"
	s = "aacabdkacaa"
	return longestPalindrome(s)
}

//时间复杂度O(n^2),空间复杂度 O(n^2),还可以优化,了解Manacher's算法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	var lenS = len(s)
	var reverseS string = ""
	for i := lenS - 1; i >= 0; i-- {
		reverseS += string(s[i])
	}
	// fmt.Println(s)
	// fmt.Println(reverseS)

	var dp [][]int = make([][]int, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]int, lenS)
	}

	var maxLen int = 0
	var maxEnd int = 0

	for i := 0; i < lenS; i++ {
		for j := 0; j < lenS; j++ {
			//形成一个由正序和逆序组成的二维表格
			//当有相同字符时,若 i 或 j 是 0,则置 1; 否则 当前 ij 与左上角数值加 1
			if s[i] == reverseS[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
			if dp[i][j] > maxLen {
				pre := lenS - j - 1 //倒置前的下标
				if pre+dp[i][j]-1 == i {
					maxLen = dp[i][j] //更新最大长度
					maxEnd = i        //更新最大结尾
				}

			}
		}
	}
	// fmt.Println("dp", dp)
	// fmt.Println(lenS, maxEnd, maxLen, maxEnd-maxLen+1, maxEnd+1)

	return s[maxEnd-maxLen+1 : maxEnd+1]
}
