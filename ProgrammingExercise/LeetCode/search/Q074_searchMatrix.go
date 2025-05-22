package search

// 74. 搜索二维矩阵 https://leetcode.cn/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个满足下述两条属性的 m x n 整数矩阵：
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

func searchMatrix(matrix [][]int, target int) (result bool) {
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	m := len(matrix) - 1
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] > target {
			m = i - 1
			break
		}
	}
	for _, v := range matrix[m] {
		if v == target {
			return true
		}
	}
	return false
}
