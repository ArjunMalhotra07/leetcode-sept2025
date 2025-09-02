package practice_graphs

import "testing"

func TestLargestIsland(t *testing.T) {
	tests := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			name:   "Single cell land",
			grid:   [][]int{{1}},
			expect: 1,
		},
		{
			name:   "Single cell water",
			grid:   [][]int{{0}},
			expect: 1, // flipping the only 0
		},
		{
			name: "All land",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			expect: 4,
		},
		{
			name: "All water",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			expect: 1, // flip any one 0
		},
		{
			name: "Small mixed grid",
			grid: [][]int{
				{1, 0},
				{0, 1},
			},
			expect: 3, // flip (0,1) or (1,0)
		},
		{
			name: "Bigger grid with one 0 inside",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expect: 9, // flip middle 0
		},
		{
			name: "Disconnected islands",
			grid: [][]int{
				{1, 0, 0, 1},
				{0, 0, 0, 0},
				{1, 1, 0, 1},
			},
			expect: 4, // best flip joins bottom-left 2 with new 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestIsland(tt.grid)
			if got != tt.expect {
				t.Errorf("largestIsland() = %v, want %v", got, tt.expect)
			}
		})
	}
}
