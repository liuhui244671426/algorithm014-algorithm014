学习笔记

本周自我感觉吃力,题目能看到,但没有思路... 教学视频也看了好多遍,怎么办?进入了迷茫期.套用模板写出来,也是有些地方调试不出来结果,可能是对递归的理解还是比较浅吧.我不知道其他同学是不是像我一样!只能对老师讲的东西先背下来.
递归模板:

```
//终止条件
//处理当前层逻辑
//下个梦境
//处理当前层局部变量(非必须)
```
分治模板:
```
//终止条件
//处理当前层逻辑
//下个梦境
//合并结果
//处理当前层局部变量(非必须)
```
本周做的好几道题目,都是参考题解的思路,翻译过来的.唯一自己思考写出来的题目是 98.<<验证二叉搜索树>> , 附上代码,分析分析

```
func isValidBST(root *TreeNode) bool {
	return helper_isValidBST(root, math.MinInt64, math.MaxInt64)
}

func helper_isValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	//fmt.Println(root)
	if root.Val <= min || root.Val >= max {
		return false
	}
	l := helper_isValidBST(root.Left, min, root.Val)
	r := helper_isValidBST(root.Right, root.Val, max)
	return l && r
}
```