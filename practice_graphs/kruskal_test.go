package practice_graphs

import (
	"testing"
)

func TestKruskalAlgorithm(t *testing.T) {
	tests := []struct {
		name     string
		graph    [][]int
		n        int
		expected int
	}{
		{
			name: "Single Edge",
			graph: [][]int{
				{},
				{0, 0, 5},
				{0, 5, 0},
			},
			n:        2,
			expected: 5,
		},
		{
			name: "Triangle",
			graph: [][]int{
				{},
				{0, 0, 1, 3},
				{0, 1, 0, 2},
				{0, 3, 2, 0},
			},
			n:        3,
			expected: 3, // 1 + 2
		},
		{
			name: "Square (Cycle)",
			graph: [][]int{
				{},
				{0, 0, 1, 0, 1},
				{0, 1, 0, 1, 0},
				{0, 0, 1, 0, 1},
				{0, 1, 0, 1, 0},
			},
			n:        4,
			expected: 3, // 1+1+1
		},
		{
			name: "Star Graph",
			graph: [][]int{
				{},
				{0, 0, 1, 1, 1, 1},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
			},
			n:        5,
			expected: 4, // four edges from center
		},
		{
			name: "Line Graph",
			graph: [][]int{
				{},
				{0, 0, 2, 0, 0, 0},
				{0, 2, 0, 2, 0, 0},
				{0, 0, 2, 0, 2, 0},
				{0, 0, 0, 2, 0, 2},
				{0, 0, 0, 0, 2, 0},
			},
			n:        5,
			expected: 8, // 2+2+2+2
		},
		{
			name: "Disconnected Graph",
			graph: [][]int{
				{},
				{0, 0, 1, 0, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 0, 0, 2},
				{0, 0, 0, 2, 0},
			},
			n:        4,
			expected: 3, // MST of comp1 (1) + comp2 (2)
		},
		{
			name: "Dense Graph",
			graph: [][]int{
				{},
				{0, 0, 1, 2, 3},
				{0, 1, 0, 2, 4},
				{0, 2, 2, 0, 5},
				{0, 3, 4, 5, 0},
			},
			n:        4,
			expected: 6, // 1+2+3
		},
		{
			name: "Kruskal Example",
			graph: [][]int{
				{},
				{0, 0, 2, 0, 1, 4, 0},
				{0, 2, 0, 3, 3, 0, 7},
				{0, 0, 3, 0, 5, 0, 8},
				{0, 1, 3, 5, 0, 9, 0},
				{0, 4, 0, 0, 9, 0, 0},
				{0, 0, 7, 8, 0, 0, 0},
			},
			n:        6,
			expected: 17,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := KruskalAlgorithm(tc.graph, tc.n)
			if got != tc.expected {
				t.Errorf("%s: expected %d, got %d", tc.name, tc.expected, got)
			}
		})
	}
}
