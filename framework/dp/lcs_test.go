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

func TestLongestCommonSubsequence2(t *testing.T) {

	if 3 != longestCommonSubsequence2("abdhjsjk", "bjk") {
		t.Error("lcs err")
	}
	if 2 != longestCommonSubsequence2("hellomyfriend", "hibaby") {
		t.Error("lcs err")
	}

	if 6 != longestCommonSubsequence2("ok namers", "oossklatterss") {
		t.Error("lcs err")
	}
}

func TestLongestCommonSubsequence3(t *testing.T) {

	if 3 != longestCommonSubsequence3("abdhjsjk", "bjk") {
		t.Error("lcs err")
	}
	if 2 != longestCommonSubsequence3("hellomyfriend", "hibaby") {
		t.Error("lcs err")
	}

	if 6 != longestCommonSubsequence3("ok namers", "oossklatterss") {
		t.Error("lcs err")
	}
}
