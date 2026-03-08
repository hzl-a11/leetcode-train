// 199. 二叉树的右视图
// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
// https://leetcode.cn/problems/binary-tree-right-side-view/description/

package binary_tree
func rightSideView(root *TreeNode) []int {
	//也就是找每层的最后一个节点
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := make([]int, 0)
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		result = append(result, queue[len(queue)-1].Val)
		for _, node := range queue {
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		queue = newQueue
	}
	return result
}
