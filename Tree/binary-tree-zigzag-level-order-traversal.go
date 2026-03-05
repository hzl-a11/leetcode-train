// 103. 二叉树的锯齿形层序遍历
// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/description/

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		slice := make([]int, 0)
		for i := 0; i < len(queue); i++ {
			slice = append(slice, queue[i].Val)
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		result = append(result, slice)
		queue = newQueue
	}
	for i := 1; i < len(result); i += 2 {
		for j, k := 0, len(result[i])-1; j < k; j, k = j+1, k-1 {
			result[i][j], result[i][k] = result[i][k], result[i][j]
		}
	}
	return result
}