package backtracking

// 17. 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
// 回溯写法
func letterCombinations(digits string) (result []string) {
	digitsCount := len(digits)
	if digitsCount == 0 {
		return
	}
	chars := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path := make([]byte, 0, digitsCount)
	var backTracking func(digitIndex int)
	backTracking = func(digitIndex int) {
		if len(path) == digitsCount {
			result = append(result, string(path))
			return
		}
		digit := digits[digitIndex] - '0'
		for _, char := range chars[digit] {
			path = append(path, byte(char))
			backTracking(digitIndex + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return
}

// 递归（隐式回溯）写法
func letterCombinationsV1(digits string) (result []string) {
	digitsCount := len(digits)
	if digitsCount == 0 {
		return
	}
	chars := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var backTracking func(digitIndex int, str string)
	backTracking = func(digitIndex int, str string) {
		if len(str) == digitsCount {
			result = append(result, str)
			return
		}
		digit := digits[digitIndex] - '0'
		for _, char := range chars[digit] {
			backTracking(digitIndex+1, str+string(char))
		}
	}
	backTracking(0, "")
	return
}

func letterCombinationsV2(digits string) (result []string) {
	digitsCount := len([]byte(digits))
	if digitsCount == 0 {
		return
	}
	chars := [][]byte{{}, {}, {'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'},
		{'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	path := make([]byte, 0, digitsCount)
	var backTracking func(digitIndex int)
	backTracking = func(digitIndex int) {
		if len(path) == digitsCount {
			result = append(result, string(path))
			return
		}
		digit := digits[digitIndex] - '0'
		for i := 0; i < len(chars[digit]); i++ {
			path = append(path, chars[digit][i])
			backTracking(digitIndex + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(0)
	return
}
