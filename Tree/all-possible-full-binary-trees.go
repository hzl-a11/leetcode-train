// 894. 所有可能的真二叉树
// 给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
// 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
// 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点
// // https://leetcode.cn/problems/all-possible-full-binary-trees/description/

package Tree

func allPossibleFBT(n int) []*TreeNode {
	//用分解的思路，将它拆成左子树，和右子树，左子树和右子树的节点数分别为i和n-1-i，i从1到n-1，每次递增2，因为真二叉树的节点数只能是奇数
	// 可以写一个构造器，开存储结果的map，key是节点数，value是所有可能的树的列表，这样就可以避免重复计算
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	if n%2 == 0 {
		// 题目描述的满二叉树不可能是偶数个节点
		return []*TreeNode{}
	}
	result := make([]*TreeNode, 0)
	for i := 1; i < n; i += 2 {
		leftTrees := allPossibleFBT(i)
		rightTrees := allPossibleFBT(n - 1 - i)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				result = append(result, &TreeNode{Val: 0, Left: leftTree, Right: rightTree})
			}
		}
	}
	return result
}
