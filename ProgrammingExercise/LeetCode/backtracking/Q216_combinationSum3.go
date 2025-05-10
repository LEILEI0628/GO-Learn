package backtracking

// 216. 组合总和 III https://leetcode.cn/problems/combination-sum-iii/description/
// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。
func combinationSum3(k int, n int) (result [][]int) {
	path := make([]int, 0, k)
	var backTracking func(start int)
	backTracking = func(start int) {
		if len(path) == k {
			sum := 0
			for _, num := range path {
				sum += num
			}
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				result = append(result, tmp)
			}
			return
		}
		for i := start; i <= 9-(k-len(path))+1; i++ {
			path = append(path, i)
			backTracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracking(1)
	return
}

func combinationSum3V1(k int, n int) (result [][]int) {
	path := make([]int, 0, k)
	var backTracking func(start int, sum int)
	backTracking = func(start int, sum int) {
		if len(path) == k {
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				result = append(result, tmp)
			}
			return
		}
		for i := start; i <= 9-(k-len(path))+1; i++ { // 剪枝
			if sum+i > n { // 剪枝
				return
			}
			path = append(path, i)
			backTracking(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	backTracking(1, 0)
	return
}
