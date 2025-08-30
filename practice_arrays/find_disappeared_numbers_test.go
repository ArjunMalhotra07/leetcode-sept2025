package practice_arrays

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Missing numbers in middle",
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			name: "All numbers present",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{},
		},
		{
			name: "All numbers missing except one",
			nums: []int{2, 2, 2, 2},
			want: []int{1, 3, 4},
		},
		{
			name: "Duplicate at start, missing number later",
			nums: []int{1, 1},
			want: []int{2},
		},
		{
			name: "Unordered input with missing",
			nums: []int{2, 3, 2, 1, 5},
			want: []int{4},
		},
		{
			name: "Empty input",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDisappearedNumbers(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
