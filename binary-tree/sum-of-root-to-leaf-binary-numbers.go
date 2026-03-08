// 1022. 从根到叶的二进制数之和
// 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。

// 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
// 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。

// 返回这些数字之和。题目数据保证答案是一个 32 位 整数。
// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/description/

package Tree

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val<<1 | node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}
