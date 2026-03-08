// 1161. 最大层内元素和
// 给你一个二叉树的根节点 root。设根节点位于二叉树的第 1 层，而根节点的子节点位于第 2 层，依此类推。
// 返回总和 最大 的那一层的层号 x。如果有多层的总和一样大，返回其中 最小 的层号 x。
// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/description/

package binary_tree
func maxLevelSum(root *TreeNode) int {
	var dfs func(node *TreeNode, depth int)
	levelSumMap := make(map[int]int) //key:层数 value:该层的节点值总和
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		levelSumMap[depth] += node.Val
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	maxSum := levelSumMap[1]
	result := 1
	for i := 2; i <= len(levelSumMap); i++ {
		if levelSumMap[i] > maxSum {
			maxSum = levelSumMap[i]
			result = i
		}
	}
	return result
}
