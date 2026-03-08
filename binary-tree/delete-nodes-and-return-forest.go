// 1110. 删点成林
// 给出二叉树的根节点 root，树上每个节点都有一个不同的值。
// 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
// 返回森林中的每棵树。你可以按任意顺序组织答案。

// https://leetcode.cn/problems/delete-nodes-and-return-forest/description/

package Tree

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	result := make([]*TreeNode, 0)
	deletedMap := make(map[int]*TreeNode)
	for _, val := range to_delete {
		deletedMap[val] = nil
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if _, ok := deletedMap[node.Val]; ok {
			deletedMap[node.Val] = node
		}
		if node.Left != nil {
			dfs(node.Left)
			if _, ok := deletedMap[node.Left.Val]; ok {
				deletedMap[node.Left.Val] = node.Left
				node.Left = nil
			}
		}
		if node.Right != nil {
			dfs(node.Right)
			if _, ok := deletedMap[node.Right.Val]; ok {
				deletedMap[node.Right.Val] = node.Right
				node.Right = nil
			}
		}
	}
	dfs(root)
	for _, node := range deletedMap {
		if node != nil {
			if node.Left != nil {
				result = append(result, node.Left)
			}
			if node.Right != nil {
				result = append(result, node.Right)
			}
			node.Left = nil
			node.Right = nil
		}
	}
	if _, ok := deletedMap[root.Val]; !ok {
		result = append(result, root)
	}
	return result
}
