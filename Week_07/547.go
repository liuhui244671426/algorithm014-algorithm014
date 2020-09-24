package Week_07

//547. 朋友圈
//连通图->并查集
func FindCircleNum() int {
	var M [][]int
	M = [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	return findCircleNum(M)
}
func findCircleNum(M [][]int) int {
	n := len(M)

	var uf UnionFind
	uf_new := uf.New(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				uf_new.Union(i, j)
			}
		}
	}
	return uf_new.Count()
}
