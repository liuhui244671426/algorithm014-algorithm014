package Mstruct

type CircularDeque struct {
	Val    int
	Next   *CircularDeque
	Priv   *CircularDeque
	Length int
}
