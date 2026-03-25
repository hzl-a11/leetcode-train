// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 你可以按 任何顺序 返回答案。k<=n

// https://leetcode.cn/problems/combinations/

package backtrace

// 也就是找只有k位的子集，start指开始的位数，idx指trace的下一个空位
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	idx := 0
	var backtrace func(start int, trace []int)
	backtrace = func(start int, trace []int) {
		if idx == k {
			result = append(result, append([]int{}, trace...))
			return
		}
		for i := start; i <= n; i++ {
			trace[idx] = i
			idx++
			backtrace(i+1, trace)
			idx--
		}
	}

	trace := make([]int, k, k)
	backtrace(1, trace)
	return result
}
