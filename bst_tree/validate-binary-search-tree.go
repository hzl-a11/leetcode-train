// 98. 验证二叉搜索树
// 有效 二叉搜索树定义如下：
// 节点的左子树只包含 严格小于 当前节点的数。
// 节点的右子树只包含 严格大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
// https://leetcode.cn/problems/validate-binary-search-tree/

package bst_tree

func isValidBST(root *TreeNode) bool {
	// 利用BST的特性，中序遍历是升序的
	var prev *TreeNode
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if prev != nil && prev.Val >= node.Val {
			return false
		}
		prev = node
		return dfs(node.Right)
	}
	return dfs(root)
}
