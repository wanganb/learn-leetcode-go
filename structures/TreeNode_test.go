package structures

import (
	"reflect"
	"testing"
)

func TestInts2TreeNode(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				ints: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints2TreeNode(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints2TreeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
