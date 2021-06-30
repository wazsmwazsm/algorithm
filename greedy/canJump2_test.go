package greedy

import (
	"testing"
)

func TestCanJump2(t *testing.T) {

	if 2 != canJump2([]int{2, 3, 1, 1, 4}) {
		t.Error("can jump err")
	}

	// 跳跃问题 2 里假设一定能到达末尾，这个用例可以不测
	// if 0 != canJump2([]int{3, 2, 1, 0, 4}) {
	// 	t.Error("can jump err")
	// }

	if 2 != canJump2([]int{3, 3, 1, 0, 4}) {
		t.Error("can jump err")
	}

	if 3 != canJump2([]int{2, 3, 1, 2, 4, 2, 3}) {
		t.Error("can jump err")
	}
}
func TestCanJump3(t *testing.T) {

	if 2 != canJump3([]int{2, 3, 1, 1, 4}) {
		t.Error("can jump err")
	}

	if 2 != canJump3([]int{3, 3, 1, 0, 4}) {
		t.Error("can jump err")
	}

	if 3 != canJump3([]int{2, 3, 1, 2, 4, 2, 3}) {
		t.Error("can jump err")
	}
}
