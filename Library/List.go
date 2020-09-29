package Library

import (
	"fmt"
)

//定义元素类型
type Element interface{}

//定义节点
type linkNode struct {
	Data Element   //数据域
	Next *linkNode //指针域，指向下一个节点
}

func NewLinkNode() *linkNode {
	return &linkNode{Data: nil}
}

//函数接口
type LinkNoder interface {
	Add(data Element)               //后面添加
	Delete(index int) Element       //删除指定index位置元素
	Insert(index int, data Element) //在指定index位置插入元素
	Length() int                    //获取元素长度
	Search(data Element) int        //查询元素的位置
	GetData(index int) Element      //获取指定index位置的元素
	GetAll() []Element              //获取所有的元素
}

//添加 头结点，数据
func (head *linkNode) Add(data Element) {
	point := head //临时指针，为最后一个节点
	for point.Next != nil {
		point = point.Next //移位
	}

	newNode := linkNode{Data: data} //需要添加的新节点
	point.Next = &newNode
}

//删除 头结点 index 位置
func (head *linkNode) Delete(index int) Element {
	//判断index合法性
	if index < 0 || index > head.Length() {
		fmt.Println("please check index")
		return fmt.Errorf("check index error")
	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Next //移位
		}
		point.Next = point.Next.Next //赋值
		data := point.Next.Data
		return data
	}
}

//插入 头结点 index位置 data元素
func (head *linkNode) Insert(index int, data Element) {
	//检验index合法性
	if index < 0 || index > head.Length() {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		newNode := &linkNode{Data: data} //初始化新加的元素
		newNode.Next = point.Next        //指定新加元素的下一个元素
		point.Next = newNode
	}
}

//获取长度 头结点
func (head *linkNode) Length() int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	length++
	return length
}

//搜索 头结点 data元素
func (head *linkNode) Search(data Element) int {
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			return index
		} else {
			index++
			point = point.Next
			if index > head.Length()-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
	if point.Data == data { //判断最后一个元素是否匹配
		fmt.Println(data, "exist at", index, "th")
		return index
	}

	return -1
}

//获取data 头结点 index位置
func (head *linkNode) GetData(index int) Element {
	point := head
	if index < 0 || index > head.Length() {
		return fmt.Errorf("check index error")
	} else {
		for i := 0; i < index-1; i++ {
			point = point.Next
		}
		return point.Data
	}
}

//获取所有的data
func (head *linkNode) GetAll() []Element {
	dataList := make([]Element, 0, head.Length())
	point := head
	for point.Next != nil {
		dataList = append(dataList, point.Data)
		point = point.Next
	}
	dataList = append(dataList, point.Data)
	return dataList
}
