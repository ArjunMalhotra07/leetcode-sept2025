package practice_graphs

import "testing"

func TestCheckGraphCycleUsingStack(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want bool
	}{
		{
			name: "No edges, no cycle",
			mtx: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: false,
		},
		{
			name: "Simple cycle 0->1->2->0",
			mtx: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 0, 0},
			},
			want: true,
		},
		{
			name: "Self loop at node 0",
			mtx: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: true,
		},
		{
			name: "Linear chain, no cycle",
			mtx: [][]int{
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Disconnected graph with a cycle",
			mtx: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 1, 0, 0}, // cycle between 1-2-3
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckGraphCycleUsingStack(tt.mtx)
			if got != tt.want {
				t.Errorf("checkGraphCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCheckGraphCycleUsingQueue(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want bool
	}{
		{
			name: "No edges, no cycle",
			mtx: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: false,
		},
		{
			name: "Simple cycle 0->1->2->0",
			mtx: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 0, 0},
			},
			want: true,
		},
		{
			name: "Self loop at node 0",
			mtx: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			want: true,
		},
		{
			name: "Linear chain, no cycle",
			mtx: [][]int{
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 0, 0, 0},
			},
			want: false,
		},
		{
			name: "Disconnected graph with a cycle",
			mtx: [][]int{
				{0, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
				{0, 1, 0, 0}, // cycle between 1-2-3
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckGraphCycleUsingQueue(tt.mtx)
			if got != tt.want {
				t.Errorf("checkGraphCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
