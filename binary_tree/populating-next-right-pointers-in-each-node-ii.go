// 117. 填充每个节点的下一个右侧节点指针 II
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。
// 初始状态下，所有 next 指针都被设置为 NULL 。

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/description/

package binary_tree
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := make([]*Node, 0)
		for i := 0; i < len(queue); i++ {
			curNode := queue[i]
			if i < len(queue)-1 {
				curNode.Next = queue[i+1]
			}
			if curNode.Left != nil {
				newQueue = append(newQueue, curNode.Left)
			}
			if curNode.Right != nil {
				newQueue = append(newQueue, curNode.Right)
			}
		}
		queue = newQueue
	}
	return root
}
