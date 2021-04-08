package dp

import (
	"fmt"
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
func TestEditDistancePath(t *testing.T) {
	s1 := "horse"
	s2 := "ros"
	p, res := editDistancePath(s1, s2)

	fmt.Printf("\ns1 %s to s2 %s\n", s1, s2)
	for _, detail := range p {
		fmt.Println(detail)
	}
	fmt.Printf("min edit distance %d\n", res)
}
func TestEditDistancePath2(t *testing.T) {
	s1 := "intention"
	s2 := "execution"
	p, res := editDistancePath(s1, s2)

	fmt.Printf("\ns1 %s to s2 %s\n", s1, s2)
	for _, detail := range p {
		fmt.Println(detail)
	}
	fmt.Printf("min edit distance %d\n", res)
}
func TestEditDistancePath3(t *testing.T) {
	s1 := "pion"
	s2 := "pationte"
	p, res := editDistancePath(s1, s2)

	fmt.Printf("\ns1 %s to s2 %s\n", s1, s2)
	for _, detail := range p {
		fmt.Println(detail)
	}
	fmt.Printf("min edit distance %d\n", res)
}
