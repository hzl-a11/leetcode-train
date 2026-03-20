package backtrace

import "testing"

func Test_permute(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want [][]int
	}{
		{
			name: "",
			nums: []int{1, 2, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
