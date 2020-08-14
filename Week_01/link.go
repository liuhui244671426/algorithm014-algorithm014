package main

import "fmt"

/*
 使用golang实现单链表
 add 操作:在链表首部添加
 append 操作:在链表尾部添加,从头部找到最后一个元素,把最后一个元素指针指向需要添加的元素
 insert 操作:在指定位置添加,往中间插入元素,需要调整插入该位置的元素之前的指针,还有该元素的指针需要指向下一位
 delete 操作:删除指定的元素,删除任意位置的元素,需要调整指针
 */

type LinkList struct {
	Data int32 // 链表上的数据
	Next *LinkList // 指针指向下一个数据
}
// 头部节点
type LinkNode struct {
	HeadNode *LinkList
}

// 判断链表是否为空,只需要判断第一个元素是否是 nil
func (self *LinkNode) IsEmpty() bool  {
	if self.HeadNode == nil {
		return true
	} else {
		return false
	}
}
// 获取链表的长度,从头部循环链表,直到元素是 nil 即可
func (self *LinkNode) Length() int  {
	count := 0
	node := self.HeadNode
	for node != nil {
		count += 1
		node = node.Next
	}
	return count
}
// 在头部添加数据,把原来的头部替换掉,还需要把现在的元素指向原来的头部
func (self *LinkNode) Add(value int32)  {
	node := &LinkList{Data:value}
	node.Next = self.HeadNode
	self.HeadNode = node
}
// 在尾部添加数据,需要从头部开始遍历,直到nil
func (self *LinkNode)Append(value int32)  {
	newNode := &LinkList{Data:value}
	node := self.HeadNode
	// 首部是空
	if node == nil {
		self.HeadNode = newNode
	} else {
		for node.Next != nil {
			node = node.Next
		}
		// 已经得到最后的元素
		node.Next = newNode
	}
}
// 在指定位置insert
func (self *LinkNode) Insert(index int, value int32)  {
	newNode := &LinkList{Data:value}
	node := self.HeadNode
	if index < 0 {
		// index 小于 0 就放在首部吧
		self.Add(value)
	} else if index > self.Length() {
		// index 大于 长度 就放在尾部部吧
		self.Append(value)
	} else {
		count := 0
		// 找到 index 之前的元素
		for count < (index - 1) {
			node = node.Next
			count += 1
		}
		// 已经找到index之前的元素
		newNode.Next = node.Next
		node.Next = newNode
	}
}
// 删除指定元素,从首部遍历,找到该元素删除,并且需要维护指针
func (self *LinkNode) Delete(value int32)  {
	node := self.HeadNode
	// 如果是首部
	if node.Data == value {
		self.HeadNode = node.Next
	} else {
		for node.Next != nil {
			// 找到了,改指针
			if node.Next.Data == value {
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
	}
}
// 循环链表输出元素
func (self *LinkNode)forEachLink()  {
	node := self.HeadNode
	for node != nil {
		str := fmt.Sprintf("当前元素 :%v", node.Data)
		fmt.Println(str)
		node = node.Next
	}
}

func main()  {
	linkList := LinkNode{}
	fmt.Println("=========Insert/Append后========")
	linkList.Add(0)
	linkList.Add(1)
	linkList.Append(2)
	linkList.Append(3)
	linkList.Append(5)
	linkList.Append(4)
	linkList.Append(6)
	linkList.forEachLink()
	linkList.Delete(4)
	fmt.Println("=============Delete后============")
	linkList.forEachLink()
	fmt.Println("=============Insert后============")
	linkList.Insert(5, 4)
	linkList.forEachLink()
	fmt.Println("=============链表长度============")
	fmt.Println(linkList.Length())

}