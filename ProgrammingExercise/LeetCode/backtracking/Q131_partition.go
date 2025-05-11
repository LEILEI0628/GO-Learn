package backtracking

// 131. 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/description/
// 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
func partition(s string) (result [][]string) {
	path := make([]string, 0, len(s))
	var backTracking func(start int)
	backTracking = func(start int) {
		if start == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}
		for i := start + 1; i <= len(s); i++ {
			str := s[start:i]
			if isPalindrome(str) {
				path = append(path, str)
				backTracking(i)
				path = path[:len(path)-1]
			}
		}
	}
	backTracking(0)
	return
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
