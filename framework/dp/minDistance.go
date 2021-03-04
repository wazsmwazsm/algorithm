package dp

/*
  两个字符串变得相同的最少删除次数
  给定两个单词 s1 s2 找到使 s1 s2 相同所需的最小步数
  每步可以删除任意一个字符串中的一个字符

  例如: "sea"  "eat"  结果 2
  第一步 sea 变为 ea，第二部 eat 变为 ea


  int minDistance(String s1, String s2);

  思考题目，两个字符串变得相同, 其实就是找最长字串，最少删除次数，就是两个字符串找到最长字串删了几个字符

*/

// 时间复杂度、空间复杂度 O(kn)
func minDistance(s1, s2 string) int {

	lS1 := len(s1)
	lS2 := len(s2)
	lcs := longestCommonSubsequence3(s1, s2)

	return (lS1 - lcs) + (lS2 - lcs)
}
