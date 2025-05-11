package backtracking

// 93. 复原 IP 地址 https://leetcode.cn/problems/restore-ip-addresses/description/
// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
func restoreIpAddresses(s string) (result []string) {
	path := make([]string, 0, 4)
	var backTracking func(start int)
	backTracking = func(start int) {
		if len(path) == 4 {
			if start != len(s) { // IP四段集齐但没有用完所有字符
				return
			}
			result = append(result, path[0]+"."+path[1]+"."+path[2]+"."+path[3])
		}
		for i := start + 1; i <= len(s) && i <= start+3; i++ {
			numStr := s[start:i]
			// 判断是否合法
			if toInt(numStr) <= 255 && (numStr == "0" || numStr[0] != '0') {
				path = append(path, numStr)
				backTracking(i)
				path = path[:len(path)-1]
			}
		}
	}
	backTracking(0)
	return
}

func toInt(s string) (num int) {
	for i := 0; i < len(s); i++ {
		num = num*10 + int(s[i]-'0')
	}
	return
}
