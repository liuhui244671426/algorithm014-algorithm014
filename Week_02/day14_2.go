package Week_02

//https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
//层序遍历
//队列 先进先出
//bfs+queue
import (
	. "homework/Mstruct"
)

func levelOrder(root *QueueNode) [][]int {
	if root == nil {
		return nil
	}

	var queue []*QueueNode //队列
	var result [][]int     //结果

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
