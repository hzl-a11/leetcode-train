// 515. 在每个树行中找最大值
// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/description/

package Tree

func largestValues(root *TreeNode) []int {
	depthMaxValueMap := make(map[int]int) //key:层数 value:该层的最大值
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depthMaxVal, ok := depthMaxValueMap[depth]
		if !ok {
			depthMaxValueMap[depth] = node.Val
		}
		if node.Val > depthMaxVal {
			depthMaxValueMap[depth] = node.Val
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	result := make([]int, len(depthMaxValueMap))
	for i := 1; i <= len(depthMaxValueMap); i++ {
		result[i-1] = depthMaxValueMap[i]
	}
	return result
}
