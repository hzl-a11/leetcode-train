// 80. 删除有序数组中的重复项 II
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
package slice

func removeDuplicates(nums []int) int {
	r, w := 0, 0 //r是读指针，用于遍历数组，w是写指针，用于删除元素
	for r < len(nums) {
		//每次遍历进来，r指向某个数字的第一次出现。
		nums[w] = nums[r]
		w++
		//若有重复，则写入第2次
		isDuplicate := false
		//通过遍历，找到这个数字的最后一次出现位置
		for r+1 < len(nums) && nums[r] == nums[r+1] {
			isDuplicate = true
			r++
		}
		if isDuplicate {
			nums[w] = nums[r]
			w++
		}
		r++ //移动读指针，使之指向他的最后一个元素
	}
	return w
}
