package main

// 3. 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

func lengthOfLongestSubstring(s string) int { // 效率较低
	str := []byte(s)
	i, maxLen := 0, 0
	for i+maxLen < len(str) {
		// 哈希集合，记录每个字符是否出现过
		m := map[byte]bool{}
		for _, v := range str[i : i+maxLen] {
			m[v] = true
		}
		if len(m) != maxLen || m[str[i+maxLen]] == true { // 窗口中有重复字符或新字符在窗口中
			i++ // 窗口向后滑动
		} else {
			m[str[i+maxLen]] = true // 将元素添加到窗口中
			maxLen++
		}
	}
	return maxLen
}

func main() {
	println(lengthOfLongestSubstring(
		"abcabcbb"))
}

func lengthOfLongestSubstring1(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长无重复字符子串
		if ans < rk-i+1 {
			ans = rk - i + 1
		}
	}
	return ans
}
