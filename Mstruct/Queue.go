package Mstruct

type QueueNode struct {
	Val      int
	Children []*QueueNode
}