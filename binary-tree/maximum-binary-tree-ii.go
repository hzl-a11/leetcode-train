// 998. 最大二叉树 II
// 最大树 定义：一棵树，并满足：其中每个节点的值都大于其子树中的任何其他值。
// 给你最大树的根节点 root 和一个整数 val 。
// 就像 之前的问题 那样，给定的树是利用 Construct(a) 例程从列表 a（root = Construct(a)）递归地构建的：

// 如果 a 为空，返回 null 。
// 否则，令 a[i] 作为 a 的最大元素。创建一个值为 a[i] 的根节点 root 。
// root 的左子树将被构建为 Construct([a[0], a[1], ..., a[i - 1]]) 。
// root 的右子树将被构建为 Construct([a[i + 1], a[i + 2], ..., a[a.length - 1]]) 。
// 返回 root 。
// 请注意，题目没有直接给出 a ，只是给出一个根节点 root = Construct(a) 。
// 假设 b 是 a 的副本，并在末尾附加值 val。题目数据保证 b 中的值互不相同。
// 返回 Construct(b) 。
// https://leetcode.cn/problems/maximum-binary-tree-ii/description/

package Tree

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

	// 题目描述的树是由一个列表构建的最大树，列表中的元素互不相同，且新元素被添加到列表末尾，所以新元素要么成为新的根节点，要么成为原树的右子树
	// 如果新元素比原树的根节点值大，那么新元素就成为新的根节点，原树成为新根节点的左子树
	// 否则，新元素就成为原树的右子树，原树的右子树成为新元素的左子树
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
