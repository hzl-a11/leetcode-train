// 331. 验证二叉树的前序序列化
// 序列化二叉树的一种方法是使用 前序遍历 。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
// https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/description/
package binary_tree

import "strings"

func isValidSerialization(preorder string) bool {
	// 用出入度规则，nil node消耗一个入度，非nil node消耗一个入度但提供两个出度
	valSlice := strings.Split(preorder, ",")
	solt := 1
	for _, val := range valSlice {
		solt--
		if solt < 0 {
			break
		}
		if !isNullNode(val) {
			solt += 2
		}
	}

	return solt == 0
}

func isNullNode(val string) bool {
	return val == "#"
}
