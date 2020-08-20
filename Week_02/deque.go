package Week_02

import "container/list"

type Deque struct {
	deque *list.List
}

func NewDeque() *Deque {
	l := list.New()
	return &Deque{l}
}

func (dq *Deque) PushFront(el interface{}) {
	dq.deque.PushFront(el)
	return
}

func (dq *Deque) PushBack(el interface{}) {
	dq.deque.PushBack(el)
	return
}

func (dq *Deque) Front() interface{} {
	if dq.deque.Front() == nil {
		return nil
	}
	return dq.deque.Front().Value
}

func (dq *Deque) Back() interface{} {
	if dq.deque.Back() == nil {
		return nil
	}
	return dq.deque.Back().Value
}

func (dq *Deque) PopFront() {
	front := dq.deque.Front()
	dq.deque.Remove(front)
}

func (dq *Deque) PopBack() {
	back := dq.deque.Back()
	dq.deque.Remove(back)
}

func (dq *Deque) Len() int {
	return dq.deque.Len()
}

func (dq *Deque) IsEmpty() bool {
	return dq.deque.Len() == 0
}
