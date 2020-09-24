package Week_07

/*
并查集
*/

//连通分量个数
// var count int

//节点 x 的父节点是parent[x]
// var parent []int

//让树平衡,这里称为重量
// var size []int

type UnionFind struct {
	count  int   //连通分量个数
	parent []int //节点 x 的父节点是parent[x]
	size   []int //让树平衡,这里称为重量
}

// 一切从这里开始
func (uf *UnionFind) New(n int) *UnionFind {
	_parent := make([]int, n)
	_size := make([]int, n)

	for i := 0; i < n; i++ {
		_parent[i] = i //建立树
		_size[i] = 1   //重量应该初始化 1
	}
	return &UnionFind{
		count:  n, //连通分量是元素个数,因为初始化个元素各自是一个集合
		parent: _parent,
		size:   _size,
	}
}

// p , q 合并
func (uf *UnionFind) Union(p, q int) {
	var P_root int = uf.Find(p)
	var Q_root int = uf.Find(q)
	//同根,不用合并
	if P_root == Q_root {
		return
	}
	//parent[P_root] = Q_root //或者是 parent[Q_root] = P_root

	//小树接到大树下面，较平衡,将树的高度减少从而减少时间复杂度到O(logN)
	if uf.size[P_root] > uf.size[Q_root] {
		uf.parent[Q_root] = P_root
		uf.size[P_root] += uf.size[Q_root]
	} else {
		uf.parent[P_root] = Q_root
		uf.size[Q_root] += uf.size[P_root]
	}
	uf.count-- //合并一次,连通分量减少一个
}

/*
查找某个节点的根节点

Find()依赖树的高度,如果树退化成了链表复杂度会变成O(N),时间复杂度O(longN)~O(N)卡住了性能,所以要对 find() 进行路径压缩优化
*/
func (uf *UnionFind) Find(x int) int {
	//根结点 parent[x]==x
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]] //路径压缩,树高度变为常数,时间复杂度O(1)
		x = uf.parent[x]                       //向上查找
	}
	return x
}

//p , q是否连通
func (uf *UnionFind) Connected(p, q int) bool {
	var P_root int = uf.Find(p)
	var Q_root int = uf.Find(q)
	//根结点相同,是同根,有连通性
	if P_root == Q_root {
		return true
	}
	return false
}

//返回连通分量个数
func (uf *UnionFind) Count() int {
	return uf.count
}
