package greedy

import (
	"testing"
)

func TestCanJump(t *testing.T) {

	if true != canJump([]int{2, 3, 1, 1, 4}) {
		t.Error("can jump err")
	}

	if false != canJump([]int{3, 2, 1, 0, 4}) {
		t.Error("can jump err")
	}
}
