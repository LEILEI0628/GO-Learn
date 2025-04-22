package main

// 226. 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
