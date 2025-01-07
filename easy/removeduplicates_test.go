package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type want struct {
		value int
		slice []int
	}
	tests := []struct {
		name string
		args args
		want
	}{
		{
			name: "case [1,1,2]",
			args: args{nums: []int{1, 1, 2}},
			want: want{value: 2, slice: []int{1, 2}},
		},
		{
			name: "case [0,0,1,1,1,2,2,3,3,4]",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: want{value: 5, slice: []int{0, 1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(&tt.args.nums); got != tt.want.value {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			assert.Equal(t, tt.want.slice, tt.args.nums)
		})
	}
}
