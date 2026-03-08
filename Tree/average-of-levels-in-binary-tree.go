// 637. 二叉树的层平均值
// 给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受
// https://leetcode.cn/problems/average-of-levels-in-binary-tree/description/

func averageOfLevels(root *TreeNode) []float64 {
	levelAvgMap := make(map[int]int) //key:层数 value:该层的节点值列表
	levelCountMap := make(map[int]int) //key:层数 value:该层的节点数量
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		levelCountMap[depth]++
		levelAvgMap[depth] += node.Val
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	result := make([]float64, len(levelAvgMap))
	for i := 1; i <= len(levelAvgMap); i++ {
		result[i-1] = float64(levelAvgMap[i]) / (float64(levelCountMap[i]))
	}
	return result
}