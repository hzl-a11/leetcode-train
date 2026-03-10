// 450. 删除二叉搜索树中的节点
// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
// 一般来说，删除节点可分为两个步骤：
// 首先找到需要删除的节点；
// 如果找到了，删除它。
// https://leetcode.cn/problems/delete-node-in-a-bst/description/

package bst_tree

// 利用BTS的特性二分搜索找到要删除的节点，用要删除的节点的左子树最大值，或右子树最小只替换掉当前节点的值，然后删除
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		return _executeDelete(root)
	}
	if root.Left != nil && root.Left.Val == key {
		root.Left = _executeDelete(root.Left)
		return root
	}
	if root.Right != nil && root.Right.Val == key {
		root.Right = _executeDelete(root.Right)
		return root
	}
	return root
}

func _executeDelete(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil && node.Right != nil {
		//找左子树最大的node or 右子树最小的node来接替自己
		node.Val = _findLeftChildMaxValAndDelete(node)
		return node
	} else if node.Left != nil {
		return node.Left
	} else if node.Right != nil {
		return node.Right
	}
	return nil
}

func _findLeftChildMaxValAndDelete(child *TreeNode) int {
	pre := child
	max := child
	for max.Right != nil {
		pre = max
		max = max.Right
	}
	pre.Right = nil
	return max.Val
}
