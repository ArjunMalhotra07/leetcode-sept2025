package practice_arrays

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Missing in middle",
			nums: []int{3, 0, 1},
			want: 2,
		},
		{
			name: "Missing at end",
			nums: []int{0, 1, 2, 3},
			want: 4,
		},
		{
			name: "Missing at start",
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			name: "Single element, missing 0",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Single element, missing 1",
			nums: []int{0},
			want: 1,
		},
		{
			name: "Two elements, missing middle",
			nums: []int{2, 0},
			want: 1,
		},
		{
			name: "Empty input",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(append([]int(nil), tt.nums...)) // copy to avoid mutation
			if got != tt.want {
				t.Errorf("missingNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
