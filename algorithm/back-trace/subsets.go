// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
// // https://leetcode.cn/problems/subsets/

package backtrace

// 思路：先找包含num[0]的所有子集=num[0]+剩余的item子集合。包含num[0]的所有子集+剩余的item子集合=结果
func subsets(nums []int) [][]int {
	var backtrace func(idx int)
	n := len(nums)
	r := make([][]int, 0)
	backtrace = func(idx int) {
		if idx == n-1 {
			r = [][]int{{nums[idx]}, nil}
			return
		}
		// 包含num[0]的所有子集=num[0]+剩余的item子集合
		backtrace(idx + 1)
		for i := range r {
			newItem := append(r[i], nums[idx])
			r = append(r, newItem)
		}
	}
	backtrace(0)
	return r
}
