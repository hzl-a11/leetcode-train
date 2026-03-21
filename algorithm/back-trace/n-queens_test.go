package backtrace

import "testing"

func Test_solveNQueens(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want [][]string
	}{
		{
			name: "",
			n:    4,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
