package string

func getNext(pattern string) []int {
	next := make([]int, len(pattern)+1)
	// 计算模式串的next数组
	// next[i]表示模式串pattern[0:i]的最长前缀和后缀的长度
	// 例如：abcab -> next[5] = 2
	// 例如：abcabc -> next[6] = 3
	// 例如：a -> next[1] = 0
	// 例如：aa -> next[2] = 1
	// 例如：aaa -> next[3] = 2
	n := len(pattern)
	next[0] = 0
	j := 0

	for i := 1; i < n; i++ {
		for j != -1 && pattern[i] != pattern[j+1] {
			j = next[j]
		}
		if pattern[i] == pattern[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
