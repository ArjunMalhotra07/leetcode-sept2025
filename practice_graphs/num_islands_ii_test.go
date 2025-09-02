package practice_graphs

import (
	"reflect"
	"testing"
)

func TestNumIslands2(t *testing.T) {
	tests := []struct {
		m, n     int
		queries  [][]int
		expected []int
	}{
		{
			m: 3, n: 3,
			queries: [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}},
			// Explanation:
			// After (0,0) → 1 island
			// After (0,1) → 1 island (merged)
			// After (1,2) → 2 islands
			// After (2,1) → 3 islands
			expected: []int{1, 1, 2, 3},
		},
		{
			m: 4, n: 5,
			queries: [][]int{
				{0, 0}, {0, 0}, {1, 1}, {1, 0}, {0, 1},
				{0, 3}, {1, 3}, {0, 4}, {3, 2}, {2, 2}, {1, 2}, {0, 2},
			},
			expected: []int{1,1,2,1,1,2,2,2,3,3,1,1},
		},
		{
			m: 1, n: 1,
			queries:  [][]int{{0, 0}, {0, 0}, {0, 0}},
			expected: []int{1, 1, 1},
		},
		{
			m: 2, n: 2,
			queries: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 0}},
			// After (0,0) → 1
			// After (1,1) → 2
			// After (0,1) → 1 (merged with (0,0))
			// After (1,0) → 1 (all connected)
			expected: []int{1, 2, 1, 1},
		},
	}

	for i, tt := range tests {
		got := NumIslands2(tt.m, tt.n, tt.queries)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("Test %d failed: got %v, expected %v", i+1, got, tt.expected)
		}
	}
}
