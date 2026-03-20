// 37. 解数独
// 编写一个程序，通过填充空格来解决数独问题。
// 数独的解法需 遵循如下规则：

// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
// 数独部分空格内已填入了数字，空白格用 '.' 表示。

// https://leetcode.cn/problems/sudoku-solver/

// DFS爆破+剪枝+一维数组转化
package backtrace

func solveSudoku(board [][]byte) {
	m, n := len(board), len(board[0]) //m是行数，n是列数
	// 从上到下遍历整个board，从左到右遍历每行，可以将这个二维数组转为一维度数组。假设某个元素在一维数组的下标是i，那么它在二维数组所在的位置是(i/n,i%n)
	var backtrace func(idx int) bool
	backtrace = func(idx int) bool {
		if idx == m*n {
			//说明遍历完了，可以直接返回
			return true
		}

		i, j := idx/n, idx%n
		if board[i][j] != '.' {
			//是已经预填的空，直接填下一位
			return backtrace(idx + 1)
		}
		for ch := byte('1'); ch <= '9'; ch++ {
			//分支去遍历
			if !isValid(ch, board, i, j) {
				//剪枝
				continue
			}
			board[i][j] = ch
			if backtrace(idx + 1) { //填下一位
				//如果找到解，直接返回
				return true
			} else {
				//如果没有，撤销选择
				board[i][j] = '.' //没找到合适选项，撤销选择

			}
		}
		return false
	}
	backtrace(0)
}

// 判断是否可以在 (r, c) 位置放置数字 num
func isValid(num byte, board [][]byte, r, c int) bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[r][i] == num {
			return false
		}
		// 判断列是否存在重复
		if board[i][c] == num {
			return false
		}
		// 判断 3 x 3 方框是否存在重复
		if board[(r/3)*3+i/3][(c/3)*3+i%3] == num {
			return false
		}
	}
	return true
}
