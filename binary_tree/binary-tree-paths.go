package binary_tree

import (
	"strconv"
)

// 257. 二叉树的所有路径
// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
// 叶子节点 是指没有子节点的节点。
// https://leetcode.cn/problems/binary-tree-paths/description/
func binaryTreePaths(root *TreeNode) []string {
	intResult := binaryTreePathsHelper(root)
	return convertIntResultToStringSlice(intResult)
}

func convertIntResultToStringSlice(intResult [][]int) []string {
	result := make([]string, 0)
	for _, path := range intResult {
		pathStr := ""
		for i := len(path) - 1; i >= 0; i-- {
			pathStr += string(strconv.Itoa(path[i])) + "->"
		}
		pathStr = pathStr[:len(pathStr)-2]
		result = append(result, pathStr)
	}
	return result
}

func binaryTreePathsHelper(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	leftTreePaths := binaryTreePathsHelper(root.Left)
	rightTreePaths := binaryTreePathsHelper(root.Right)
	result := make([][]int, 0)
	for _, path := range leftTreePaths {
		result = append(result, append(path, root.Val))
	}
	for _, path := range rightTreePaths {
		result = append(result, append(path, root.Val))
	}
	return result
}
