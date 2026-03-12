package algorithm

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "",
			s:    " 2-1 + 2 ",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculate(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
