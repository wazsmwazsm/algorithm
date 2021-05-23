package dp

import (
	"testing"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	if 4 != longestPalindromeSubseq("bbbab") {
		t.Error("lps err")
	}
	if 11 != longestPalindromeSubseq("ssjdkkkdjjllddss") {
		t.Error("lps err")
	}
	if 4 != longestPalindromeSubseq("helloh") {
		t.Error("lps err")
	}
	if 1 != longestPalindromeSubseq("abc") {
		t.Error("lps err")
	}
	if 3 != longestPalindromeSubseq("adbcag") {
		t.Error("lps err")
	}
}
