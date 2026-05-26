// 3120. 统计特殊字母的数量 I
// 给你一个字符串 word。如果 word 中同时存在某个字母的小写形式和大写形式，则称这个字母为 特殊字母 。
// 返回 word 中 特殊字母 的数量。

package dailyrandom

func numberOfSpecialChars(word string) int {
	chSlice := make([]int, 52) //记录每个字母出现的次数，优化成位运算，可以节省存储空间
	result := 0
	for i := range len(word) {
		c := word[i]
		switch {
		case 'a' <= c && c <= 'z':
			chSlice[c-'a']++
		case 'A' <= c && c <= 'Z':
			chSlice[c-'A'+26]++
		}
	}
	for i := range 26 { //判断并统计特殊字母
		if chSlice[i] != 0 && chSlice[i+26] != 0 {
			result++
		}
	}
	return result
}
