package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{name: "1", args: args{
			l1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
			l2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			want: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 1, Next: &ListNode{
						Val: 2, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4, Next: &ListNode{
									Val: 4}}}}}}},
		{name: "2", args: args{l1: nil, l2: nil}, want: nil},
		{name: "3", args: args{l1: nil, l2: &ListNode{Val: 0}}, want: &ListNode{Val: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
