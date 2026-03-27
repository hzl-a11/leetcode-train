// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次
// 返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。

// https://leetcode.cn/problems/combination-sum-iii/

package backtrace

func combinationSum3(k int, n int) [][]int {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if n > k*nums[len(nums)-1] || n < k*nums[0] {
		return nil
	}
	result := make([][]int, 0)
	trace := make([]int, 0)

	var backtrace func(start, target int)
	backtrace = func(start, target int) {
		if (len(trace) == k && target != 0) || (len(trace) != k && target == 0) {
			return
		}
		if target == 0 && len(trace) == k {
			temp := make([]int, len(trace))
			copy(temp, trace)
			result = append(result, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			value := nums[i]
			if value > target {
				break
			} else {
				trace = append(trace, value)
				backtrace(i+1, target-value)
				trace = trace[:len(trace)-1]
			}
		}
	}
	backtrace(0, n)
	return result
}
