package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Add test cases.
		{
			"case1",
			args{
				[]int{1, 2, 3},
				5,
			},
			[]int{1, 2},
		},
		{
			"case1",
			args{
				[]int{5, 6, 3, 8, 0},
				5,
			},
			[]int{0, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
