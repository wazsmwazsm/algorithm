package dp

import (
	"testing"
)

func TestMinDistance(t *testing.T) {

	if 5 != minDistance("abdhjsjk", "bjk") {
		t.Error("min distance err")
	}
	if 2 != minDistance("sea", "eat") {
		t.Error("min distance err")
	}
	if 15 != minDistance("hellomyfriend", "hibaby") {
		t.Error("min distance err")
	}
}
