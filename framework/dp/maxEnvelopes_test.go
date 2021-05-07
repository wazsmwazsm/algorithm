package dp

import (
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	var e Envelopes = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	if 3 != maxEnvelopes(e) {
		t.Error("max envelopes err")
	}
}
