package practiceheaps

import (
	"reflect"
	"testing"
)

func TestMinHeap_AddAndExtract(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int // expected ascending order after extracting all
	}{
		{
			name:     "random order",
			input:    []int{10, 3, 4, 5, 2, 1, 100, 56},
			expected: []int{1, 2, 3, 4, 5, 10, 56, 100},
		},
		{
			name:     "already sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "descending input",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "duplicates",
			input:    []int{7, 7, 7, 7},
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &MinHeap{}

			// insert all numbers
			for _, num := range tc.input {
				h.add(num)
			}

			// extract all in sorted order
			var result []int
			for range tc.input {
				result = append(result, h.extractRoot())
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestExtractRootOnEmptyHeap(t *testing.T) {
	h := &MinHeap{}
	val := h.extractRoot()
	if val != -1 {
		t.Errorf("expected -1 for empty heap, got %d", val)
	}
}
