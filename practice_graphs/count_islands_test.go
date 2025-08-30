package practice_graphs

import "testing"

func TestCountIslands(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want int
	}{
		{
			name: "Empty graph",
			mtx:  [][]int{},
			want: 0,
		},
		{
			name: "Single node, no edges",
			mtx: [][]int{
				{0},
			},
			want: 1,
		},
		{
			name: "Two disconnected nodes",
			mtx: [][]int{
				{0, 0},
				{0, 0},
			},
			want: 2,
		},
		{
			name: "Two connected nodes",
			mtx: [][]int{
				{0, 1},
				{1, 0},
			},
			want: 1,
		},
		{
			name: "Three nodes in a line (0-1-2)",
			mtx: [][]int{
				{0, 1, 0},
				{1, 0, 1},
				{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "Three disconnected nodes",
			mtx: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 3,
		},
		{
			name: "Two components: (0-1) and (2-3)",
			mtx: [][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 1},
				{0, 0, 1, 0},
			},
			want: 2,
		},
		{
			name: "Fully connected graph of 4 nodes",
			mtx: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 1},
				{1, 1, 0, 1},
				{1, 1, 1, 0},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countIslands(tt.mtx)
			if got != tt.want {
				t.Errorf("countIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
