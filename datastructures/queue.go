package datastructures

type Queue interface {
	Push(x int)
	Pop() int
	Peek() int
	Empty() bool
}
