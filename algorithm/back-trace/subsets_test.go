package backtrace

import "testing"

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want [][]int
	}{
		{nums: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
