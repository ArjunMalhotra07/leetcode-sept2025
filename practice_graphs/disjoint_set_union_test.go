package practice_graphs

import "testing"

func TestUnionFind(t *testing.T) {
	tests := []struct {
		name     string
		vertices int
		unions   [][2]int
		queries  [][2]int
		expected []bool
	}{
		{
			name:     "single component",
			vertices: 5,
			unions:   [][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			queries:  [][2]int{{1, 5}, {2, 4}, {1, 3}},
			expected: []bool{true, true, true},
		},
		{
			name:     "disconnected sets",
			vertices: 6,
			unions:   [][2]int{{1, 2}, {3, 4}},
			queries:  [][2]int{{1, 2}, {1, 3}, {4, 3}, {5, 6}},
			expected: []bool{true, false, true, false},
		},
		{
			name:     "merge two sets",
			vertices: 6,
			unions:   [][2]int{{1, 2}, {3, 4}, {2, 3}},
			queries:  [][2]int{{1, 4}, {2, 3}, {5, 6}},
			expected: []bool{true, true, false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uf := NewUnionBySize(tc.vertices)

			for _, e := range tc.unions {
				uf.AddEdge(e[0], e[1])
			}

			for i, q := range tc.queries {
				got := uf.IsConnected(q[0], q[1])
				want := tc.expected[i]
				if got != want {
					t.Errorf("query %v-%v: got %v, want %v", q[0], q[1], got, want)
				}
			}
		})
	}
}
