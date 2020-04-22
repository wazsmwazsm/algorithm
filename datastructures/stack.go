package datastructures

type Stack interface {
	Push(x int)
	Pop() int
	Top() int
	Empty() bool
}

// MySliceStack stack by slice 不安全, 想要安全要加锁
type MySliceStack []int

func NewMySliceStack() MySliceStack {
	return MySliceStack([]int{})
}
func (s *MySliceStack) Push(x int) {
	*s = append(*s, x)
}

func (s *MySliceStack) Pop() int {
	n := len(*s)
	if n > 0 {
		top := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return top
	}

	return 0
}

func (s *MySliceStack) Top() int {
	n := len(*s)
	if n > 0 {
		return (*s)[len(*s)-1]
	}
	return 0
}

func (s *MySliceStack) Empty() bool {

	return len(*s) == 0
}
