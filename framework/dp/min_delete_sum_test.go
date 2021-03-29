package dp

import (
	"testing"
)

func TestMinimumDeleteSum(t *testing.T) {
	if 231 != minimumDeleteSum("sea", "eat") { // remove s t
		t.Error("minimumDeleteSum err")
	}

	if 620 != minimumDeleteSum("comebaby", "obty") { // remove t c m e a b
		t.Error("minimumDeleteSum err")
	}
}
