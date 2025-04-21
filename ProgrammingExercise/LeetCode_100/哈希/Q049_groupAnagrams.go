package main

import "sort"

// 49. 字母异位词分组 https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 方法一：排序
// 由于互为字母异位词的两个字符串包含的字母相同，因此对两个字符串分别进行排序之后得到的字符串一定是相同的，故可以将排序之后的字符串作为哈希表的键。
func groupAnagrams(strs []string) [][]string {
	strMap := map[string][]string{} // key是排序后的str，value是异位词分组
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { // 对str进行排序
			return s[i] < s[j]
		})
		strSort := string(s)
		strMap[strSort] = append(strMap[strSort], str) // 排序后的str作为key

	}
	result := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		result = append(result, v)
	}
	return result
}

// 方法二：计数
// 由于互为字母异位词的两个字符串包含的字母相同，因此两个字符串中的相同字母出现的次数一定是相同的，故可以将每个字母出现的次数使用字符串表示，作为哈希表的键。
// 由于字符串只包含小写字母，因此对于每个字符串，可以使用长度为 26 的数组记录每个字母出现的次数。需要注意的是，在使用数组作为哈希表的键时，不同语言的支持程度不同，因此不同语言的实现方式也不同。
func groupAnagrams1(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
