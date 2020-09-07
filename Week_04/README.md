学习笔记

本周学习要点:

1  二分查找
2  贪心
3  深度优先搜索
4  广度优先搜索

学习感悟:

1  二分查找,将字符串拆成 2 段,进行比较大小,以达到加快查询速度的算法
2  贪心算法,是当下环境最好选择而非全局最好选择,重点在于"贪婪"二字.
3  深度优先搜索,放在二叉树的遍历顺序是前序遍历 preorder, 将这一种可能计算到底.
4  广度优先搜索,更像是二叉树的层级顺序的遍历方式, 适合使用队列辅助遍历顺序.


代码模板:
```
#dfs 模板-栈

def dfs(self, tree):
    if tree.root is None:
        return []
    visited, stack = [], [tree.root]
    while stack:
        node = stack.pop()
        visited.app(node)
        #process
        nodes = generate_relate_nodes(node)
        stack.push(nodes)
        #other process works
        ...


#dfs 模板-递归
visited = set()
def dfs2(node, visited):
    visited.add(node)
    #process current node here
    ...
    for next_node in node.children():
        if not next_node in visited:
            dfs2(next_node, visited)

#bfs 模板-队列
def bfs(self, graph, start, end):
    quque = []
    visited = set()
    quque.append([start])
    visited.add(start)

    #process
    nodes = generate_relate_nodes(node)
    quque.push(nodes)
    #other process works
    ...

#二分查找
left, right = 0, len(array) - 1
while left <= right:
    mid = (left+right)/2 # left+((right-left)/2)
    if array[mid] == target:
        break or return result
    else if array[mid] < target:
        left = mid + 1
    else:
        right = mid - 1

```
