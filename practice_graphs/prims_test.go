package practice_graphs

import (
	"reflect"
	"sort"
	"testing"
)

// ---------- MinHeap Unit Tests ----------
func sortTriples(a [][]int) {
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] != a[j][0] {
			return a[i][0] < a[j][0] // weight
		}
		if a[i][1] != a[j][1] {
			return a[i][1] < a[j][1] // node
		}
		return a[i][2] < a[j][2] // parent
	})
}

func TestMinHeap_AddAndExtract(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name:  "random order",
			input: [][]int{{3, 1, 0}, {1, 2, 0}, {5, 3, 1}, {2, 4, 2}},
			want:  [][]int{{1, 2, 0}, {2, 4, 2}, {3, 1, 0}, {5, 3, 1}},
		},
		{
			name:  "already sorted ascending",
			input: [][]int{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}},
			want:  [][]int{{1, 1, 0}, {2, 2, 0}, {3, 3, 0}},
		},
		{
			name:  "descending input",
			input: [][]int{{5, 1, 0}, {4, 2, 0}, {3, 3, 0}, {1, 4, 0}},
			want:  [][]int{{1, 4, 0}, {3, 3, 0}, {4, 2, 0}, {5, 1, 0}},
		},
		{
			name:  "duplicates",
			input: [][]int{{2, 1, 0}, {2, 2, 0}, {2, 3, 0}},
			want:  [][]int{{2, 1, 0}, {2, 2, 0}, {2, 3, 0}},
		},
		{
			name:  "single element",
			input: [][]int{{42, 1, 0}},
			want:  [][]int{{42, 1, 0}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			h := &MinHeap{}
			for _, item := range tc.input {
				h.add(item)
			}

			var got [][]int
			for len(h.Nums) > 0 {
				got = append(got, h.extractRoot())
			}

			// sort both slices before comparing
			sortTriples(got)
			sortTriples(tc.want)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestExtractRootOnEmptyHeap(t *testing.T) {
	h := &MinHeap{}
	got := h.extractRoot()
	want := []int{-1, -1, -1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractRoot() on empty heap = %v, want %v", got, want)
	}
}

// ---------- MST Integration Test ----------

func TestCheckForHeap_PrimsMST(t *testing.T) {
	graph := [][]int{
		{0, 2, 1, 0, 0}, // node 0
		{2, 0, 1, 0, 0}, // node 1
		{1, 1, 0, 2, 2}, // node 2
		{0, 0, 2, 0, 1}, // node 3
		{0, 0, 2, 1, 0}, // node 4
	}

	// Expected MST edges and weight
	wantEdges := [][]int{
		{1, 2, 0}, // 0 -- 2 (w=1)
		{1, 1, 2}, // 2 -- 1 (w=1)
		{2, 4, 2}, // 2 -- 4 (w=2)
		{1, 3, 4}, // 4 -- 3 (w=1)
	}
	wantWeight := 5

	h := &MinHeap{}
	h.add([]int{0, 0, -1})
	visited := make([]bool, 5)
	var ans [][]int
	dist := 0

	for len(h.Nums) != 0 {
		root := h.extractRoot()
		wt, node, parent := root[0], root[1], root[2]
		if !visited[node] && parent != -1 {
			ans = append(ans, root)
			dist += wt
		}
		if !visited[node] {
			visited[node] = true
			for i, weight := range graph[node] {
				if weight > 0 && !visited[i] {
					h.add([]int{weight, i, node})
				}
			}
		}
	}

	if !reflect.DeepEqual(ans, wantEdges) {
		t.Errorf("MST edges got %v, want %v", ans, wantEdges)
	}
	if dist != wantWeight {
		t.Errorf("MST weight got %d, want %d", dist, wantWeight)
	}
}
