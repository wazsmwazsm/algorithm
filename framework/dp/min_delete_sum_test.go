package dp

import (
	"testing"
)

func TestMinimumDeleteSum(t *testing.T) {
	if 231 != minimumDeleteSum("eat", "sea") { // remove s t
		t.Error("minimumDeleteSum err")
	}

	if 620 != minimumDeleteSum("comebaby", "obty") { // remove t c m e a b
		t.Error("minimumDeleteSum err")
	}
}
func TestMinimumDeleteSum2(t *testing.T) {
	if 231 != minimumDeleteSum2("eat", "sea") { // remove s t
		t.Error("minimumDeleteSum err")
	}

	if 620 != minimumDeleteSum2("comebaby", "obty") { // remove t c m e a b
		t.Error("minimumDeleteSum err")
	}
}
