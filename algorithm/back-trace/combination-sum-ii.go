// 40. 组合总和 II
// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
// 注意：解集不能包含重复的组合。

// https://leetcode.cn/problems/combination-sum-ii/

package backtrace

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	result := make([][]int, 0)
	var backtrace func(start int, tar int, trace []int)
	backtrace = func(start int, tar int, trace []int) {
		for i := start; i < len(candidates); i++ {
			// 如果当前元素和前一个元素相同，且我们不是在处理当前层级的第一个元素，则跳过
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			value := candidates[i]
			if value > tar {
				break
			} else if value == tar {
				temp := make([]int, len(trace))
				copy(temp, trace)
				temp = append(temp, value)
				result = append(result, temp)
				break
			} else {
				trace = append(trace, value)
				backtrace(i+1, tar-value, trace)
				trace = trace[:len(trace)-1]
			}
		}
	}

	backtrace(0, target, []int{})
	return result
}
