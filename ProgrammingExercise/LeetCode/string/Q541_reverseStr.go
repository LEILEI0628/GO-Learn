package string

// 541. 反转字符串 II https://leetcode.cn/problems/reverse-string-ii/description/
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
func reverseStr(s string, k int) string {
	str := []byte(s)
	length := len(s)
	start := 0
	for start+k <= length { // 对最后没有剩余或剩余量大于k的情况（最后k的值大于length，不会进入下面的循环）
		for i := 0; i < k/2; i++ {
			str[start+i], str[start+k-i-1] = str[start+k-i-1], str[start+i]
		}
		start += 2 * k
	}
	if start < length { // 对最后剩余元素小于k的情况单独处理
		for i := 0; i < (length-start)/2; i++ {
			str[start+i], str[length-i-1] = str[length-i-1], str[start+i]
		}
	}
	return string(str)
}
