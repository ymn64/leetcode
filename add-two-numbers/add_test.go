package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	l1   *ListNode
	l2   *ListNode
	want *ListNode
}{
	{"test 1", NewList(2, 4, 3), NewList(5, 6, 4), NewList(7, 0, 8)},
	{"test 2", NewList(0), NewList(0), NewList(0)},
	{"test 3", NewList(9, 9, 9, 9, 9, 9, 9), NewList(9, 9, 9, 9), NewList(8, 9, 9, 9, 0, 0, 0, 1)},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.l1, tt.l2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
