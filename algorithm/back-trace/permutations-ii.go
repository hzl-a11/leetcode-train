// 47. 全排列 II
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

// https://leetcode.cn/problems/permutations-ii/

package backtrace

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	trace := make([]int, 0)
	var backtrace func(start int)
	backtrace = func(start int) {
		traceMap := make(map[int]struct{})
		if start == len(nums) {
			result = append(result, append([]int{}, trace...))
		}
		for i := start; i < len(nums); i++ {
			if _, ok := traceMap[nums[i]]; ok {
				continue
			}
			traceMap[nums[i]] = struct{}{}
			trace = append(trace, nums[i])
			nums[i], nums[start] = nums[start], nums[i]
			backtrace(start + 1)
			trace = trace[:len(trace)-1]
			nums[i], nums[start] = nums[start], nums[i]

		}
	}
	backtrace(0)
	return result
}
