package main

// 59. 螺旋矩阵 II https://leetcode.cn/problems/spiral-matrix-ii/description/
// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		for i := top + 1; i <= bottom; i++ {
			result[i][right] = num
			num++
		}
		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				result[bottom][i] = num
				num++
			}
			for i := bottom - 1; i > top; i-- {
				result[i][left] = num
				num++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
