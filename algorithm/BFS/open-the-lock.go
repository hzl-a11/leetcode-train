// 752. 打开转盘锁
// 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
// 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
// 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
// 字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。
// https://leetcode.cn/problems/open-the-lock/description/

package bfs

func openLock(deadends []string, target string) int {
	deadMap := make(map[string]struct{})
	for i := range deadends {
		deadMap[deadends[i]] = struct{}{}
	}
	vistedMap := make(map[string]struct{})
	queue := []string{"0000"}
	step := 0
	for len(queue) > 0 {

		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			if _, ok := vistedMap[cur]; ok {
				continue
			}
			if _, ok := deadMap[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			if _, ok := vistedMap[cur]; ok {
				continue
			}
			if _, ok := deadMap[cur]; ok {
				continue
			}
			vistedMap[cur] = struct{}{}
			//拨动一位，4个中的任意一个，上拨或者下拨动
			for i := 0; i < 4; i++ {
				queue = append(queue, plusOne(cur, i))
				queue = append(queue, minusOne(cur, i))
			}

		}
		step++
	}

	return -1
}

// 将 s[j] 向上拨动一次
func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j]--
	}
	return string(ch)
}
