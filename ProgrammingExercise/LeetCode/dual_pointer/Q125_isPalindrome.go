package main

import (
	"strings"
)

// 125. 验证回文串 https://leetcode.cn/problems/valid-palindrome/description/?envType=study-plan-v2&envId=top-interview-150

func isPalindrome(s string) bool {
	var newStr string
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			newStr += string(c)
		}
	}

	newStr = strings.ToLower(newStr)

	for i := 0; i < len(newStr)/2; i++ {
		if newStr[i] != newStr[len(newStr)-1-i] {
			return false
		}
	}
	return true
}

func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for !isResult(s[left]) {
			if left >= len(s)-1 {
				return true
			}
			left++
		}
		for !isResult(s[right]) {
			if right <= 0 {
				return true
			}
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isResult(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
