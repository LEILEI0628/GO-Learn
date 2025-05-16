package hash

// 242. 有效的字母异位词 https://leetcode.cn/problems/valid-anagram/description/
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。

func isAnagram(s string, t string) bool { // 统计字符出现频率（数组方式）
	if len(s) != len(t) {
		return false
	}
	hash := make([]int, 26)
	for _, c := range s {
		hash[byte(c)-'a']++
	}
	for _, c := range t {
		hash[byte(c)-'a']--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramV1(s string, t string) bool { // 统计字符出现频率（存储字符方式）
	if len(s) != len(t) {
		return false
	}
	charMap := make(map[byte]int, len(s))
	for _, c := range s {
		charMap[byte(c)] += 1
	}
	for _, c := range t {
		count := charMap[byte(c)]
		if count == 0 {
			return false
		} else if count == 1 {
			delete(charMap, byte(c))
		} else {
			charMap[byte(c)] -= 1
		}
	}
	return len(charMap) == 0
}
