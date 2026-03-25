package backtrace

import "testing"

func Test_combine(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		k    int
		want [][]int
	}{
		{n: 4,
			k: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.n, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
