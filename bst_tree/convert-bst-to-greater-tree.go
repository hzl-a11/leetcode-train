// 538. 把二叉搜索树转换为累加树
// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
// 提醒一下，二叉搜索树满足下列约束条件：
// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。
// 注意：本题和 1038: https://leetcode.cn/problems/binary-search-tree-to-greater-sum-tree/ 相同
// https://leetcode.cn/problems/convert-bst-to-greater-tree/description/

package bst_tree

func convertBST(root *TreeNode) *TreeNode {
	// 利用BST的特性，中序遍历的倒叙（右-根-左），是一个升序的遍历
	// 我们将这个遍历结果简化成一个slice，会发现，每个节点的新值=前面遍历过的所有节点的和+当前节点的值
	// 因此，我们可以在遍历的过程中，维护一个sum变量，记录前面遍历过的节点的和，每次访问一个节点时，将当前节点的值加到sum上，并将当前节点的值更新为sum
	var sum int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}
