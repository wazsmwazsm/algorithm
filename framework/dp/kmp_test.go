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
