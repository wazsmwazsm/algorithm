package dp

import (
	"testing"
)

func TestMaxEnvelopes(t *testing.T) {
	var e Envelopes = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	if 3 != maxEnvelopes(e) {
		t.Error("max envelopes err")
	}

	e = [][]int{{8, 4}, {6, 3}, {11, 7}, {2, 3}, {1, 1}}
	if 4 != maxEnvelopes(e) {
		t.Error("max envelopes err")
	}
}
func TestMaxEnvelopes2(t *testing.T) {
	var e Envelopes = [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	if 3 != maxEnvelopes2(e) {
		t.Error("max envelopes err")
	}

	e = [][]int{{8, 4}, {6, 3}, {11, 7}, {2, 3}, {1, 1}}
	if 4 != maxEnvelopes2(e) {
		t.Error("max envelopes err")
	}
}
