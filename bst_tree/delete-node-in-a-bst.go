// 450. 删除二叉搜索树中的节点
// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
// 一般来说，删除节点可分为两个步骤：
// 首先找到需要删除的节点；
// 如果找到了，删除它。
// https://leetcode.cn/problems/delete-node-in-a-bst/description/

package bst_tree

// 利用BTS的特性二分搜索找到要删除的节点，用要删除的节点的左子树最大值，或右子树最小只替换掉当前节点的值
func deleteNode(root *TreeNode, key int) *TreeNode {
	// todo code review
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root
		}
		leftChildMax := _findLeftChildMaxVal(root)
		root.Left = deleteNode(root.Left, leftChildMax)
		root.Val = leftChildMax
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root.Left
	}
	return root
}

func _findLeftChildMaxVal(node *TreeNode) int {
	max := node.Left
	for max.Right != nil {
		max = max.Right
	}
	return max.Val
}
