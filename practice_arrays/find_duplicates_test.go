package practice_arrays

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Single duplicate in middle",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "Single duplicate at start",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "Duplicate repeated multiple times",
			nums: []int{2, 2, 2, 2, 2},
			want: 2,
		},
		{
			name: "Duplicate is the largest number",
			nums: []int{1, 4, 6, 3, 2, 5, 6},
			want: 6,
		},
		{
			name: "Duplicate is the smallest number",
			nums: []int{1, 1, 2, 3, 4, 5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("findDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
