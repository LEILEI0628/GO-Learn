package main

// 48. 旋转图像 https://leetcode.cn/problems/rotate-image/description/?envType=study-plan-v2&envId=top-100-liked
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

func rotate(matrix [][]int) {
	// 数学本质：两次翻转等于一次旋转（转置后翻转）
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ { // 仅遍历主对角线下方元素
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i] // 转置
		}
	}
	for _, row := range matrix {
		for i := 0; i < len(row)/2; i++ {
			row[i], row[len(row)-i-1] = row[len(row)-i-1], row[i] // 水平翻转
		}
	}
}
