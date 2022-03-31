package dp

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {

	if 3 != climbStairs(3) {
		t.Error("climbStairs err")
	}

	if 5 != climbStairs(4) {
		t.Error("climbStairs err")
	}
}
