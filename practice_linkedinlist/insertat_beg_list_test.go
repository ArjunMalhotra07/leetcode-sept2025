package practice_linkedinlist

import (
	"reflect"
	"testing"
)

func TestInsertAtBeg(t *testing.T) {
	tests := []struct {
		name         string
		head         *ListNode
		newNodeVal   int
		wantListVals []int
	}{
		{
			name:         "insert into empty list",
			head:         nil,
			newNodeVal:   10,
			wantListVals: []int{10},
		},
		{
			name:         "insert at beginning of single node list",
			head:         &ListNode{Num: 20},
			newNodeVal:   10,
			wantListVals: []int{10, 20},
		},
		{
			name:         "insert at beginning of multi-node list",
			head:         createList([]int{30, 40, 50}),
			newNodeVal:   20,
			wantListVals: []int{20, 30, 40, 50},
		},
		{
			name:         "insert negative number",
			head:         createList([]int{1, 2, 3}),
			newNodeVal:   -5,
			wantListVals: []int{-5, 1, 2, 3},
		},
		{
			name:         "insert zero",
			head:         createList([]int{100, 200}),
			newNodeVal:   0,
			wantListVals: []int{0, 100, 200},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newNode := &ListNode{Num: tt.newNodeVal}
			got := InsertAtBeg(tt.head, newNode)

			gotVals := listToSlice(got)
			if !reflect.DeepEqual(gotVals, tt.wantListVals) {
				t.Errorf("InsertAtBeg() = %v, want %v", gotVals, tt.wantListVals)
			}
		})
	}
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Num: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Num: vals[i]}
		current = current.Next
	}

	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Num)
		current = current.Next
	}

	return result
}

func compareLists(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Num != l2.Num {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
