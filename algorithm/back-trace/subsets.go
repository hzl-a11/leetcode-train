// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// // https://leetcode.cn/problems/subsets/

package backtrace

// 思路：先找包含num[0]的所有子集=num[0]+剩余的item子集合。包含num[0]的所有子集+剩余的item子集合=结果
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	trace := make([]int, 0)
	var backtrace func(start int)
	backtrace = func(start int) {
		res = append(res, append([]int{}, trace...))
		for i := start; i < len(nums); i++ {
			trace = append(trace, nums[i])
			backtrace(i + 1)
			trace = trace[:len(trace)-1]
		}
	}
	backtrace(0)
	return res
}
