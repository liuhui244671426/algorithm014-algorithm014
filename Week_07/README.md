学习笔记

学习要点:
1.并查集
2.trie 树
3.双向 bfs
4.avl 树 红黑树
5.A*搜索

代码模板:
```
# Python A*搜索
def AstarSearch(graph, start, end):
	pq = collections.priority_queue() # 优先级 —> 估价函数
	pq.append([start]) 
	visited.add(start)
	while pq: 
		node = pq.pop() # can we add more intelligence here ?
		visited.add(node)
		process(node) 
		nodes = generate_related_nodes(node) 
   unvisited = [node for node in nodes if node not in visited]
		pq.push(unvisited)


# C++ 双向 BFS 搜索
void BFS_bothsides()//双向BFS 
{
    if(s1.state==s2.state)//起点终点相同时要特判
    {
           //do something
           found=true;
           return;
    }
    bool found=false;
    memset(visited,0,sizeof(visited));  // 判重数组
    while(!Q1.empty())  Q1.pop();   // 正向队列
    while(!Q2.empty())  Q2.pop();  // 反向队列
    //======正向扩展的状态标记为1，反向扩展标记为2
    visited[s1.state]=1;   // 初始状态标记为1
    visited[s2.state]=2;   // 结束状态标记为2
    Q1.push(s1);  // 初始状态入正向队列
    Q2.push(s2);  // 结束状态入反向队列
    while(!Q1.empty() || !Q2.empty())
    {
           if(!Q1.empty())
                  BFS_expand(Q1,true);  // 在正向队列中搜索
           if(found)  // 搜索结束 
                  return ;
          if(!Q2.empty())
                  BFS_expand(Q2,false);  // 在反向队列中搜索
           if(found) // 搜索结束
                  return ;
    }
}

```