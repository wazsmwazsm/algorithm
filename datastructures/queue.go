package datastructures

type Queue interface {
	Offer(x int)
	Pull() int
	Peek() int
	Empty() bool
}
