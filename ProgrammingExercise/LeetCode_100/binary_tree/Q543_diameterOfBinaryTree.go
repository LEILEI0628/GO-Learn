package main

// 543. 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一棵二叉树的根节点，返回该树的 直径 。
// 二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
// 两节点之间路径的 长度 由它们之间边数表示。

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	ans = 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		left := dfs(node.Left) + 1
		right := dfs(node.Right) + 1
		ans = max(ans, left+right)
		return max(left, right)
	}
	dfs(root)
	return
}
