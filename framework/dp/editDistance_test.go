package dp

import (
	"testing"
)

func TestEditDistance(t *testing.T) {

	if 3 != editDistance("horse", "ros") {
		t.Error("edit distance err")
	}
	if 5 != editDistance("intention", "execution") {
		t.Error("edit distance err")
	}
}
func TestEditDistance2(t *testing.T) {

	if 3 != editDistance2("horse", "ros") {
		t.Error("edit distance err")
	}
	if 5 != editDistance2("intention", "execution") {
		t.Error("edit distance err")
	}
}
func TestEditDistance3(t *testing.T) {

	if 3 != editDistance3("horse", "ros") {
		t.Error("edit distance err")
	}
	if 5 != editDistance3("intention", "execution") {
		t.Error("edit distance err")
	}
}
