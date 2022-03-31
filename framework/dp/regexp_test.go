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
func TestIsMath4(t *testing.T) {

	if !isMath4("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMath4("abc", ".*") {
		t.Error("is match err")
	}

	if !isMath4("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMath4("ccc", "a.*") {
		t.Error("is match err")
	}
}

func TestIsMath5(t *testing.T) {

	if !isMath5("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMath5("abb", ".a*b*") {
		t.Error("is match err")
	}

	if !isMath5("hello", "h.*o") {
		t.Error("is match err")
	}

	if isMath5("hello", "hs*o") {
		t.Error("is match err")
	}

	if !isMath5("hello", "a*llo") {
		t.Error("is match err")
	}

	if !isMath5("hello", "a*") {
		t.Error("is match err")
	}

	if !isMath5("abc", ".*") {
		t.Error("is match err")
	}

	if !isMath5("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMath5("ccc", "a.*") {
		t.Error("is match err")
	}

	// 错误格式
	if isMath5("hello", "*b*") {
		t.Error("is match err")
	}
	if isMath5("hello", "a**") {
		t.Error("is match err")
	}
}
