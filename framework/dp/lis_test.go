package dp

import (
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {

	if 4 != longestIncreasingSubsequence([]int{10, 9, 2, 5, 3, 7, 101, 18}) {
		t.Error("lis err")
	}
	if 5 != longestIncreasingSubsequence([]int{10, 9, 2, 5, 3, 7, 101, 18, 20}) {
		t.Error("lis err")
	}
	if 5 != longestIncreasingSubsequence([]int{17, 8, 4, 3, 6, 11, 44, 15, 2, 79}) {
		t.Error("lis err")
	}
	if 5 != longestIncreasingSubsequence([]int{10, 2, 4, 6, 7, 8, 1, 5, 3}) {
		t.Error("lis err")
	}
}

// func TestLongestIncreasingSubsequence2(t *testing.T) {

// 	if 4 != longestIncreasingSubsequence2([]int{10, 9, 2, 5, 3, 7, 101, 18}) {
// 		t.Error("lis err")
// 	}
// 	if 5 != longestIncreasingSubsequence2([]int{10, 9, 2, 5, 3, 7, 101, 18, 20}) {
// 		t.Error("lis err")
// 	}

// 	if 5 != longestIncreasingSubsequence2([]int{17, 8, 4, 3, 6, 11, 44, 15, 2, 79}) {
// 		t.Error("lis err")
// 	}

// 	if 5 != longestIncreasingSubsequence2([]int{10, 2, 4, 6, 7, 8, 1, 5, 3}) {
// 		t.Error("lis err")
// 	}
// }
