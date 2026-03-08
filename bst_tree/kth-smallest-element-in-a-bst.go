// 230. 二叉搜索树中第 K 小的元素
// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（k 从 1 开始计数）。
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/description/

package bst_tree

// 利用BST特性，中序遍历是升序的
func kthSmallest(root *TreeNode, k int) int {
	counter := k
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		if node.Left != nil {
			left := dfs(node.Left)
			if left != -1 && counter <= 0 {
				return left
			}
		}
		counter--
		if counter == 0 {
			return node.Val
		}
		return dfs(node.Right)
	}
	return dfs(root)
}
