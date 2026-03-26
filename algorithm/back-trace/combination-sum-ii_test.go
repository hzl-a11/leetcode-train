package backtrace

import "testing"

func Test_combinationSum2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		candidates []int
		target     int
		want       [][]int
	}{
		{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
