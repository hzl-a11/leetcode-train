// 224. 基本计算器
// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

// https://leetcode.cn/problems/basic-calculator/

//遍历，如果是数字，就累加，num = 10*num + int(c-'0')
//当遇到非数字时，说明一个数字已经算好了，可以加到pop中
//利用战来进行运算
// +和-：用一个sign，来存储数字前面的符号，如果是+，直接存数值，如果是-，存相反值，这样直接累加栈中的所有元素即可
// *和/：用当前num与stack中的前一个num进行计算
// （）：其实就是迭代进去了，用一个map来存储左右括号关系，方便直接计算出括号中的子计算的位置

package algorithm

import (
	"errors"
	"strings"
	"sync"
)

type Stack struct {
	mu      sync.Mutex
	items   []int
	lastIdx int
}

func (s *Stack) Pop() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.lastIdx < 0 {
		return 0, errors.ErrUnsupported
	}
	val := s.items[s.lastIdx]
	s.lastIdx--
	return val, nil
}
func (s *Stack) Push(val int) {
	if s.lastIdx+1 >= len(s.items) {
		s.items = append(s.items, val)
	} else {
		s.items[s.lastIdx+1] = val
	}
	s.lastIdx++
}

func NewStack() *Stack {
	return &Stack{
		items:   make([]int, 0),
		lastIdx: -1,
	}
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	bucketMap := make(map[int]int) //key：左括号的idx，右括号的idx
	stack := NewStack()
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(':
			stack.Push(i)
		case ')':
			leftIdx, _ := stack.Pop()
			bucketMap[leftIdx] = i
		}
	}
	return _calculate(s, 0, len(s)-1, bucketMap)
}

func _calculate(s string, left, right int, bracketMap map[int]int) int {
	if left < 0 || right >= len(s) {
		return 0
	}
	stack := NewStack()
	sign := 1
	num := 0
	for i := left; i <= right; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			// 如果是数字字符，进行累加，从而得到真实的树枝
			num = 10*num + int(c-'0')
		}
		if c == '+' || c == '-' || i == right {
			stack.Push(num * sign)
			if c == '+' {
				sign = 1
			} else {
				sign = -1
			}
		} else if c == '*' {
			preNum, _ := stack.Pop()
			stack.Push(num * preNum)
		} else if c == '/' {
			preNum, _ := stack.Pop()
			stack.Push(num / preNum)
		} else if c == '(' {
			// 左括号，对应的右括号的位置：
			rightIdx := bracketMap[i]
			val := _calculate(s, i+1, rightIdx-1, bracketMap)
			stack.Push(val * sign)
			i = rightIdx
		}
		if !(c >= '0' && c <= '9') {
			num = 0
		}
	}
	res := 0
	for {
		if val, err := stack.Pop(); err == nil {
			res += val
		} else {
			break
		}
	}
	return res
}
