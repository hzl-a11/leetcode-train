// 107. 二叉树的层序遍历 II
// 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/description/

package binary_tree

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]TreeNode, 0)
	queue = append(queue, *root)
	for len(queue) > 0 {
		newQueue := make([]TreeNode, 0)
		slice := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			slice = append(slice, queue[i].Val)
			if queue[i].Left != nil {
				newQueue = append(newQueue, *queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, *queue[i].Right)
			}
		}
		result = append(result, slice)
		queue = newQueue
	}
	// 反转result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
