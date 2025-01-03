package medium

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l  *listNode
		ll *listNode
	}
	tests := []struct {
		name string
		args args
		want *listNode
	}{
		{
			name: "case 243 564",
			args: args{
				l: &listNode{
					Val: 2,
					Next: &listNode{
						Val: 4,
						Next: &listNode{
							Val: 3,
						},
					},
				},
				ll: &listNode{
					Val: 5,
					Next: &listNode{
						Val: 6,
						Next: &listNode{
							Val: 4,
						},
					},
				},
			},
			want: &listNode{
				Val: 7,
				Next: &listNode{
					Val: 0,
					Next: &listNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "case 0",
			args: args{
				l: &listNode{
					Val: 0,
				},
				ll: &listNode{
					Val: 0,
				},
			},
			want: &listNode{
				Val: 0,
			},
		},
		{
			name: "case 9999999 9999",
			args: args{
				l:  fillList([]int{9, 9, 9, 9, 9, 9, 9}),
				ll: fillList([]int{9, 9, 9, 9}),
			},
			want: fillList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name: "case 5 5",
			args: args{
				l:  fillList([]int{5}),
				ll: fillList([]int{5}),
			},
			want: fillList([]int{0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l, tt.args.ll); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fillList(nums []int) *listNode {
	result := new(listNode)
	var parent *listNode
	for k, i := range nums {
		if k == 0 {
			result.Val = i
			parent = result
		} else {
			el := &listNode{
				Val: i,
			}
			parent.Next = el
			parent = el
		}
	}
	return result
}
