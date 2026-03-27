package backtrace

import "testing"

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		candidates []int
		target     int
		want       [][]int
	}{
		{candidates: []int{2, 3, 6, 7}, target: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
