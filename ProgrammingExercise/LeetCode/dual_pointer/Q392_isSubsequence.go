package main

// 392. 判断子序列 https://leetcode.cn/problems/is-subsequence/description/?envType=study-plan-v2&envId=top-interview-150

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	} else if len(t) == 0 {
		return false
	}
	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return i == len(s)
}

func isSubsequence1(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(s)
}
