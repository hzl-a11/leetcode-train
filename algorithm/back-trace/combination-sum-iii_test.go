package backtrace

import "testing"

func Test_combinationSum3(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		k    int
		n    int
		want [][]int
	}{
		{k: 9, n: 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum3(tt.k, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
