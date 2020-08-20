package Week_02

//import "deque"

type MoQueue struct {
	data *Deque
	//data *deque.Deque
}

func NewMoQueue() MoQueue {
	return MoQueue{
		data: NewDeque(),
	}
}

func (mo *MoQueue) Push(x int) {
	for mo.data.IsEmpty() == false && x > mo.data.Back().(int) {
		mo.data.PopBack()
	}
	mo.data.PushBack(x)
}

// 单调队列的最大值
func (Mo *MoQueue) Max() int {
	return Mo.data.Front().(int)
}

// 单调队列pop一个元素
func (Mo *MoQueue) Pop(x int) {
	if Mo.data.IsEmpty() == false && Mo.data.Front() == x {
		Mo.data.PopFront()
	}
}
