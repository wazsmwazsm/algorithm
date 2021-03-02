package dp

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {

	if 3 != longestCommonSubsequence("abdhjsjk", "bjk") {
		t.Error("lcs err")
	}
	if 2 != longestCommonSubsequence("hellomyfriend", "hibaby") {
		t.Error("lcs err")
	}

	if 6 != longestCommonSubsequence("ok namers", "oossklatterss") {
		t.Error("lcs err")
	}
}
