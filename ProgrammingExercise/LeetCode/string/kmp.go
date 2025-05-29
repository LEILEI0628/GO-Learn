package string

func kmpSearch(text, pattern string) int {
	next := getNext(pattern)
	i, j := 0, 0
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				return i - j
			}
		} else {
			if j == 0 {
				i++ // 如果j为0，说明第一个字符就不匹配，i需要向后移动一位
			} else {
				j = next[j-1] // j回到next[j-1]的位置
			}
		}

	}
	return -1
}

func getNext(pattern string) []int {
	next := make([]int, 1, len(pattern))
	preLen := 0                     // 前缀长度
	for i := 1; i < len(pattern); { // 从第二个字符开始计算next数组
		if pattern[preLen] == pattern[i] {
			preLen += 1
			next = append(next, preLen) // 如果匹配成功，前缀长度加1
			i++
		} else {
			if preLen == 0 {
				next = append(next, 0) // 如果前缀长度为0，说明没有匹配成功，next[i]为0
				i++                    // i向后移动一位
			} else {
				preLen = next[preLen-1] // 如果前缀长度不为0，回到上一个前缀的长度（例：ABACABAB的最后一位）
			}
		}
	}
	return next
	// 计算模式串的next数组
	// next[i]表示模式串pattern[0:i]的最长前缀和后缀的长度
	// 例如：abcab -> next[5] = 2
	// 例如：abcabc -> next[6] = 3
	// 例如：a -> next[1] = 0
	// 例如：aa -> next[2] = 1
	// 例如：aaa -> next[3] = 2
}
