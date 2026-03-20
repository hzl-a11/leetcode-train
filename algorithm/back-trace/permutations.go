package backtrace

// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// https://leetcode.cn/problems/permutations/

// 用分解+交换的想法
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		subChoice := permute(nums[1:len(nums)])
		for t := range subChoice {
			result = append(result, append([]int{nums[0]}, subChoice[t]...))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}

	return result
}
