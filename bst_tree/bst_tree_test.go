package bst_tree

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	n := func(v int, l, r *TreeNode) *TreeNode {
		return &TreeNode{Val: v, Left: l, Right: r}
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		root *TreeNode
		key  int
		want *TreeNode
	}{
		{
			name: "delete root from [5,3,6,2,4,null,7], key=5",
			root: n(5, n(3, n(2, nil, nil), n(4, nil, nil)), n(6, nil, n(7, nil, nil))),
			key:  5,
			want: n(4, n(3, n(2, nil, nil), nil), n(6, nil, n(7, nil, nil))),
		},
		{
			name: "delete leaf from [5,3,6,2,4,null,7], key=7",
			root: n(5, n(3, n(2, nil, nil), n(4, nil, nil)), n(6, nil, n(7, nil, nil))),
			key:  7,
			want: n(5, n(3, n(2, nil, nil), n(4, nil, nil)), n(6, nil, nil)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteNode(tt.root, tt.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
