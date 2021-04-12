package structures

import (
	"reflect"
	"testing"
)

func TestInts2ListNode(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// Add test cases.
		{
			name: "case1",
			args: args{
				ints: []int{1, 2, 3, 4},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ints2ListNode(tt.args.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints2ListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
