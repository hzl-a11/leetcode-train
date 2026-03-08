// 1609. 奇偶树
// 二叉树根节点所在层下标为 0 ，根的子节点所在层下标为 1 ，根的孙节点所在层下标为 2 ，依此类推。
// 偶数下标 层上的所有节点的值都是 奇 整数，从左到右按顺序 严格递增
// 奇数下标 层上的所有节点的值都是 偶 整数，从左到右按顺序 严格递减
// 给你二叉树的根节点，如果二叉树为 奇偶树 ，则返回 true ，否则返回 false 。
// https://leetcode.cn/problems/even-odd-tree/description/

package Tree

func isEvenOddTree(root *TreeNode) bool {
	depthPreValMap := make(map[int]int) //key:层数 value:该层前一个节点的值
	var dfs func(node *TreeNode, depth int) bool
	dfs = func(node *TreeNode, depth int) bool {
		if node == nil {
			return true
		}
		// 判断是否满足，偶数层的值为奇数
		if node.Val%2 == depth%2 {
			return false
		}
		// 判断是否严格递增或者递减
		if preVal, ok := depthPreValMap[depth]; ok && ((depth%2 == 0 && preVal >= node.Val) || (depth%2 == 1 && preVal <= node.Val)) {
			return false
		}
		depthPreValMap[depth] = node.Val
		if !dfs(node.Left, depth+1) {
			return false
		}
		return dfs(node.Right, depth+1)
	}
	return dfs(root, 0)
}
