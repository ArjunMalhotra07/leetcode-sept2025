package practice_arrays

import (
	"reflect"
	"testing"
)

func TestCycleSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Already sorted",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse order",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Random order",
			arr:  []int{3, 5, 2, 1, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Single element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "Two elements swapped",
			arr:  []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "Large array",
			arr:  []int{7, 1, 5, 3, 2, 6, 4},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CycleSort(append([]int(nil), tt.arr...)) // copy to avoid mutating original
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CycleSort(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
