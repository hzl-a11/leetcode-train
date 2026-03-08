// 1457. 二叉树中的伪回文路径
// 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
// 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
// https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree/description/

package binary_tree
func pseudoPalindromicPaths(root *TreeNode) int {
	var dfs func(node *TreeNode, count [10]int) int
	dfs = func(node *TreeNode, count [10]int) int {
		if node == nil {
			return 0
		}
		count[node.Val]++
		if node.Left == nil && node.Right == nil {
			oddCount := 0
			for i := 1; i <= 9; i++ {
				if count[i]%2 != 0 {
					oddCount++
				}
			}
			if oddCount <= 1 {
				return 1
			}
			return 0
		}
		return dfs(node.Left, count) + dfs(node.Right, count)
	}

	var count [10]int
	return dfs(root, count)
}

// 一下写法是不行的，因为 map 是引用类型，在递归过程中会被修改，导致结果不正确
// func pseudoPalindromicPaths(root *TreeNode) int {
// 	var dfs func(node *TreeNode, count []int) int
// 	dfs = func(node *TreeNode, countMap map[int]struct{}) int {
// 		if node == nil {
// 			return 0
// 		}

// 		if _, ok := countMap[node.Val]; ok {
// 			delete(countMap, node.Val)
// 		} else {
// 			countMap[node.Val] = struct{}{}
// 		}
// 		if node.Left == nil && node.Right == nil {
// 			if len(countMap) <= 1 {
// 				return 1
// 			}
// 			return 0
// 		}
// 		return dfs(node.Left, countMap) + dfs(node.Right, countMap)
// 	}
// 	return dfs(root, make(map[int]struct{}))

// }
