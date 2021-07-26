package dp

import (
	"testing"
)

func TestSearch(t *testing.T) {

	if 4 != search("aaabaaac", "aaac") {
		t.Error("Search err")
	}

	if 2 != search("hello", "ll") {
		t.Error("Search err")
	}
}
func TestSearchKMP(t *testing.T) {

	kmp := NewKMP("aaac")

	if 4 != kmp.Search("aaabaaac") {
		t.Error("Search err")
	}

	kmp = NewKMP("ll")
	if 2 != kmp.Search("hello") {
		t.Error("Search err")
	}
}
