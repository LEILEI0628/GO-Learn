package main

// 54. 螺旋矩阵 https://leetcode.cn/problems/spiral-matrix/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func spiralOrder(matrix [][]int) []int {
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
