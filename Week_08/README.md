学习总结

学习知识点
1.位运算
2.排序算法(1.手写排序在 sort.go 2.时间复杂度在 排序算法时间复杂度.md)
3.布隆过滤器(hash table + duoble link)
4.LRU 淘汰

```
#布隆过滤器代码模板 python

import mmh3
from bitarray import bitarray


# zhihu_crawler.bloom_filter

# Implement a simple bloom filter with murmurhash algorithm.
# Bloom filter is used to check wether an element exists in a collection, and it has a good performance in big data situation.
# It may has positive rate depend on hash functions and elements count.



BIT_SIZE = 5000000

class BloomFilter:
    
    def __init__(self):
        # Initialize bloom filter, set size and all bits to 0
        bit_array = bitarray(BIT_SIZE)
        bit_array.setall(0)

        self.bit_array = bit_array
        
    def add(self, url):
        # Add a url, and set points in bitarray to 1 (Points count is equal to hash funcs count.)
        # Here use 7 hash functions.
        point_list = self.get_postions(url)

        for b in point_list:
            self.bit_array[b] = 1

    def contains(self, url):
        # Check if a url is in a collection
        point_list = self.get_postions(url)

        result = True
        for b in point_list:
            result = result and self.bit_array[b]
    
        return result

    def get_postions(self, url):
        # Get points positions in bit vector.
        point1 = mmh3.hash(url, 41) % BIT_SIZE
        point2 = mmh3.hash(url, 42) % BIT_SIZE
        point3 = mmh3.hash(url, 43) % BIT_SIZE
        point4 = mmh3.hash(url, 44) % BIT_SIZE
        point5 = mmh3.hash(url, 45) % BIT_SIZE
        point6 = mmh3.hash(url, 46) % BIT_SIZE
        point7 = mmh3.hash(url, 47) % BIT_SIZE


        return [point1, point2, point3, point4, point5, point6, point7]
```

```
#LRU cache

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
```