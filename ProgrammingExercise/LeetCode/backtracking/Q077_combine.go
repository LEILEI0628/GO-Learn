package backtracking

// 77. 组合 https://leetcode.cn/problems/combinations/description/
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。
func combine(n int, k int) (result [][]int) {
	path := make([]int, 0, 2)
	var backTracking func(n int, k int, start int)
	backTracking = func(n int, k int, start int) {
		if len(path) == k { // path中有k个元素了，满足要求放入result中并返回
			tmp := make([]int, k)
			copy(tmp, path) // 指针最后指向的同一个（最终的）path问题
			result = append(result, tmp)
			return
		}
		//for i := start; i <= n; i++ {
		for i := start; i <= n-(k-len(path))+1; i++ { // 剪枝
			path = append(path, i)    // 选择元素
			backTracking(n, k, i+1)   // 递归
			path = path[:len(path)-1] // 回溯
		}
	}
	backTracking(n, k, 1)
	return
}
