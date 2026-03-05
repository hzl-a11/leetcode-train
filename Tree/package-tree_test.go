package Tree

import (
	"testing"
)

func TestDelNodes(t *testing.T) {
	// 构造树：[1,2,3,null,null,null,4]
	//     1
	//    / \
	//   2   3
	//        \
	//         4
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	toDelete := []int{2, 1}
	result := delNodes(root, toDelete)

	// 删除节点 2 和 1 后，应该得到两棵树：
	// 树1：3 (节点3及其右子树)
	// 树2：4 (节点4单独成树，作为节点3的右子树)
	if len(result) != 1 {
		t.Errorf("Expected 1 tree, got %d", len(result))
	}

	// 验证结果中是否包含值为 3 的根节点
	values := make([]int, 0)
	for _, tree := range result {
		values = append(values, tree.Val)
	}

	if !contains(values, 3) {
		t.Errorf("Expected tree with root 3, got %v", values)
	}
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
