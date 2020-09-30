package Week_08

//LRU cache

/*
hash table + duoble link

*/
type LinkNode struct {
	key, val  int       //hash table 必备 k/v
	pre, next *LinkNode //duoble link 必备pre ,next
}

type LRUCache struct {
	hashTable  map[int]*LinkNode
	cap        int //hash table 的容量,超过这个容量会出发移除旧元素
	head, tail *LinkNode
}

//初始化
func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil} //头尾,均初始化未空
	tail := &LinkNode{0, 0, nil, nil} //头尾,均初始化未空
	//初始化,需要把头和尾互相指向对方 head.next->tail tail.pre->head
	head.next = tail //头->尾
	tail.pre = head  //尾->头
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

//取
func (this *LRUCache) Get(key int) int {
	cache := this.hashTable
	if v, exist := cache[key]; exist {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

//移除node结点
func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//添加node节点
func (this *LRUCache) AddNode(node *LinkNode) {
	head := this.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

//移动到头部
func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

//放
func (this *LRUCache) Put(key int, value int) {
	tail := this.tail
	cache := this.hashTable

	if v, exist := cache[key]; exist {
		//存在于 hash table
		//将元素移到头部即可
		v.val = value
		this.MoveToHead(v)
	} else {
		//不存在于 hash table
		//步骤:
		//1 . 如果超出 hash table 长度则 A 删除尾部元素; B 删除 hash table 的 k/v;
		//2 . 添加新元素到头部;	添加hash table 的k/v
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == this.cap {
			delete(cache, tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		this.AddNode(v)
		cache[key] = v
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
