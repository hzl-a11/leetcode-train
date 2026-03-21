// 51. N 皇后
// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

// https://leetcode.cn/problems/n-queens/

package backtrace

// todo check this func,reply me with Chinese
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	row := 0
	board := make([]int, n)
	for i := range n {
		board[i] = -1
	}
	// 棋盘中每行只有1个Q，所以用一个[]int来存储每行Q的位置
	var backtrace func(board []int) bool
	backtrace = func(board []int) bool { //用DFS遍历每个选择

		//DFS,遍历本行的每个选项
		for i := 0; i < n; i++ {
			if !_isValid(board, row, i) {
				//当前格不满足规则，剪枝跳过
				continue
			}
			board[row] = i

			if row == n-1 {
				//说明棋盘每一行符合规则地填好了，记录结果
				result = append(result, convertIntBoard2StringBoard(board))
				return true
			}

			row++
			backtrace(board) //计算剩余行
			//撤销当前选择，以进行下一个选择
			board[row] = -1
			row--
		}
		return false
	}
	backtrace(board)
	return result
}

func _isValid(board []int, r, c int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i] == c {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := r-1, c-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i] == j {
			return false
		}
	}
	// 检查右边上方是否有皇后互相冲突
	for i, j := r-1, c+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i] == j {
			return false
		}
	}
	return true
}
func convertIntBoard2StringBoard(board []int) []string {
	n := len((board))
	byteBoard := make([][]byte, n, n)
	for i := range n {
		byteBoard[i] = make([]byte, n)
		for j := range n {
			byteBoard[i][j] = '.'
		}
	}
	for i := range board {
		byteBoard[i][board[i]] = 'Q'
	}
	stringBoard := make([]string, n, n)
	for i := range byteBoard {
		stringBoard[i] = string(byteBoard[i])
	}
	return stringBoard
}
