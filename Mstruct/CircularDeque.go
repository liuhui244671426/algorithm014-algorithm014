package Mstruct

type MyCircularDeque struct {
	Val    int
	Next   *MyCircularDeque
	Priv   *MyCircularDeque
	Length int
}
