package bst_tree

// 二叉搜索树（Binary Search Tree，BST
// 1、对于 BST 的每一个节点 node，左子树节点的值都比 node 的值要小，右子树节点的值都比 node 的值大。
// 2、对于 BST 的每一个节点 node，它的左侧子树和右侧子树都是 BST。
// 一个重要的性质：BST 的中序遍历结果是有序的（升序）



type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
