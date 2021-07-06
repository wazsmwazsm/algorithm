package dp

import (
	"testing"
)

func TestIsMath(t *testing.T) {

	if !isMath("abc", "abc") {
		t.Error("is match err")
	}

	if isMath("abc", "a.c") {
		t.Error("is match err")
	}
}
func TestIsMath2(t *testing.T) {

	if !isMath2("abc", "abc") {
		t.Error("is match err")
	}

	if !isMath2("abc", "a.c") {
		t.Error("is match err")
	}

	if isMath2("abc", "ac") {
		t.Error("is match err")
	}
}
func TestIsMath3(t *testing.T) {

	if !isMath3("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMath3("abc", ".*") {
		t.Error("is match err")
	}

	if !isMath3("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMath3("ccc", "a.*") {
		t.Error("is match err")
	}
}
