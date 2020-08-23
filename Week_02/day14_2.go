package Week_02

//层序遍历
//队列 先进先出
//bfs+queue
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var queue []*Node  //队列
	var result [][]int //结果

	queue = append(queue, root)

	for len(queue) > 0 {
		var layer []int
		len := len(queue)
		for i := 0; i < len; i++ {
			top := queue[0]
			queue = queue[1:]
			for _, child := range top.Children {
				queue = append(queue, child)
			}
			layer = append(layer, top.Val)
		}
		result = append(result, layer)
	}
	return result
}
