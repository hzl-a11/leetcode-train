// 129. 求根节点到叶节点数字之和
// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：

// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/description/
package binary_tree
func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

// 递归函数，参数为当前节点和父节点的路径和
func sumNumbersHelper(root *TreeNode, parentSum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return parentSum*10 + root.Val
	}
	leftSum := sumNumbersHelper(root.Left, parentSum*10+root.Val)
	rightSum := sumNumbersHelper(root.Right, parentSum*10+root.Val)
	return leftSum + rightSum
}
