package Week_18

//91. 解码方法

func NumDecodings() int {
	var s string
	s = "226"     //1 => 2 2 6 | 2 => 22 6 | 3 => 2 26
	s = "1201234" // 3
	return numDecodings(s)
}

/*
首先注意以0开头的直接返回0，在这一点上浪费了很多时间！
设dp[i]记录s[0~i]的解码方法， 从第二位起分以下四种情况讨论：

当前位为0，如果前一位为1或2，则只能是当前位和前一位一起解码，dp[i] = dp[i-2]，否则，直接返回0；
前一位为1，且当前为不为0，则有当前位单独解码和与前一位一起解码两种情况，即dp[i] = dp[i-1] + dp[i-2];
前一位为2，且当前为在1~6之间，则有当前位单独解码和与前一位一起解码两种情况，即dp[i] = dp[i-1] + dp[i-2];
其他情况，只能是单独解码，dp[i] = dp[i-1]。
*/
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	prepre, pre, cur := 1, 1, 1 //dp[i-2] dp[i-1] dp[i]
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = prepre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur += prepre
		} else {
			cur = pre
		}
		prepre = pre
		pre = cur
	}
	return cur
}
