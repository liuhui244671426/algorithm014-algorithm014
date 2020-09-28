package Week_07

//单词接龙
//warning 提交超时,时间复杂度待优化
func LadderLength() int {
	var beginWord string
	var endWord string
	var wordList []string

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot", "dot", "dog", "lot", "log", "cog"}
	wordList = []string{"hot", "dot", "dog", "lot", "log"}

	return ladderLength(beginWord, endWord, wordList)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var count int = 0
	var node string
	var visited map[string]bool = make(map[string]bool)
	var queue []string = []string{beginWord}

	for len(queue) > 0 {
		size := len(queue)
		count += 1
		for i := 0; i < size; i++ {

			node = queue[0]
			queue = queue[1:]
			// fmt.Println("queue node", queue, node)
			for _, w := range wordList {
				if visited[w] == true {
					continue
				}
				if isCan(node, w) > 1 {
					continue
				}
				// fmt.Println("node endWord", node, endWord)
				if node == endWord || isCan(node, endWord) == 1 {
					return count + 1
				}

				visited[node] = true
				visited[w] = true
				queue = append(queue, w)
			}

		}
	}

	return 0
}

func isCan(a, b string) int {
	if len(a) != len(b) {
		return 0
	}
	if a == b {
		return 0
	}
	var count int = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
			if count > 1 {
				return count
			}

		}
	}
	return 1
}
