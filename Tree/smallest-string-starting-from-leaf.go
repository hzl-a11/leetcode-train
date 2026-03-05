// 988. 从叶结点开始的最小字符串
// 给定一颗根结点为 root 的二叉树，树中的每一个结点都有一个 [0, 25] 范围内的值，分别代表字母 'a' 到 'z'。
// 返回 按字典序最小 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束。
// 注：字符串中任何较短的前缀在 字典序上 都是 较小 的：
// 例如，在字典序上 "ab" 比 "aba" 要小。叶结点是指没有子结点的结点。
// 节点的叶节点是没有子节点的节点。

// https://leetcode.cn/problems/smallest-string-starting-from-leaf/description/

package Tree

// 用边组装边比较的方式，来优化空间复杂度，避免存储所有路径
func smallestFromLeaf(root *TreeNode) string {

	result := ""
	var dfs func(*TreeNode, string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}
		curPath := string(rune(root.Val+'a')) + path
		if root.Left == nil && root.Right == nil {
			result = smallerString(result, curPath)
		}
		dfs(root.Left, curPath)
		dfs(root.Right, curPath)
	}
	dfs(root, "")
	return result
}

func smallerString(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	if s2 == "" {
		return s1
	}
	if s1 < s2 {
		return s1
	}
	return s2
}
