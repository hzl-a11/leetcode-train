// 1302. 层数最深叶子节点的和
// 给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
// https://leetcode.cn/problems/deepest-leaves-sum/description/

package binary_tree
func deepestLeavesSum(root *TreeNode) int {
	maxLevel := 0
	levelSumMap := make(map[int]int) //key:层数 value:该层的节点值总和
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > maxLevel {
			maxLevel = depth
		}
		levelSumMap[depth] += node.Val
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	return levelSumMap[maxLevel]
}
