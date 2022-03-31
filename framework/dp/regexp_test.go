package dp

import (
	"regexp"
	"testing"
)

func TestIsMatch(t *testing.T) {

	if !isMatch("abc", "abc") {
		t.Error("is match err")
	}

	if isMatch("abc", "a.c") {
		t.Error("is match err")
	}
}
func TestIsMatch2(t *testing.T) {

	if !isMatch2("abc", "abc") {
		t.Error("is match err")
	}

	if !isMatch2("abc", "a.c") {
		t.Error("is match err")
	}

	if isMatch2("abc", "ac") {
		t.Error("is match err")
	}
}
func TestIsMatch3(t *testing.T) {

	if !isMatch3("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMatch3("abc", ".*") {
		t.Error("is match err")
	}

	if !isMatch3("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMatch3("ccc", "a.*") {
		t.Error("is match err")
	}
}
func TestIsMatch4(t *testing.T) {

	if !isMatch4("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMatch4("abc", ".*") {
		t.Error("is match err")
	}

	if !isMatch4("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMatch4("ccc", "a.*") {
		t.Error("is match err")
	}
}

func TestIsMatch5(t *testing.T) {

	if !isMatch5("abb", "c*a*b*") {
		t.Error("is match err")
	}

	if !isMatch5("abb", ".a*b*") {
		t.Error("is match err")
	}

	if !isMatch5("hello", "h.*o") {
		t.Error("is match err")
	}

	if isMatch5("hello", "hs*o") {
		t.Error("is match err")
	}

	if !isMatch5("hello", "a*llo") {
		t.Error("is match err")
	}

	if !isMatch5("hello", "a*") {
		t.Error("is match err")
	}

	if !isMatch5("abc", ".*") {
		t.Error("is match err")
	}

	if !isMatch5("a88kksjshh", "a.*") {
		t.Error("is match err")
	}

	if isMatch5("ccc", "a.*") {
		t.Error("is match err")
	}

	// 错误格式
	if isMatch5("hello", "*b*") {
		t.Error("is match err")
	}
	if isMatch5("hello", "a**") {
		t.Error("is match err")
	}
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatch5("hello", "a*llo")
	}
}
func BenchmarkRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.Match("a*llo", []byte("hello"))
	}
}
