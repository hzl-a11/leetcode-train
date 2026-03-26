// 90. 子集 II
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
// https://leetcode.cn/problems/subsets-ii/

package backtrace

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	trace := make([]int, 0)
	n := len(nums)

	//这里的排序很重要，避免顺序不一致导致的问题（即[1,4] [4,1]是一样的）
	// 假设，nums[4,1,4]
	//先确定第一位为4，会加入【4，1】这个结果集
	//然后确定第一位为1，此时仅需考虑idx>=1的位数就可以了，因为idx=0的元素，在前面的选择已经包含到
	//但由于数组支持重复，idx>=1的item的值会和idx=0的值重复，导致出现顺序不一致重复的问题
	//通过排序，可以保证不会出现这种情况，从而保证唯一性
	sort.Ints(nums)
	var backtrace func(idx int)
	backtrace = func(idx int) {
		result = append(result, append([]int(nil), trace...))
		traceMap := make(map[int]struct{}, 0)

		for i := idx; i < n; i++ {
			value := nums[i]
			if _, ok := traceMap[value]; ok {
				continue
			}
			traceMap[value] = struct{}{}
			trace = append(trace, value)
			backtrace(i + 1)
			trace = trace[:len(trace)-1]

		}
	}
	backtrace(0)
	return result
}
