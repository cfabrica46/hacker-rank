package permutingtwoarrays_test

import (
	"testing"

	"permutingtwoarrays"

	"gotest.tools/assert"
)

func TestPangrams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		inK  int32
		inA  []int32
		inB  []int32
		out  string
	}{
		{
			name: "YES",
			inK:  10,
			inA:  []int32{2, 1, 3},
			inB:  []int32{7, 8, 9},
			out:  "YES",
		},
		{
			name: "NO",
			inK:  5,
			inA:  []int32{1, 2, 2, 1},
			inB:  []int32{3, 3, 3, 4},
			out:  "NO",
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := permutingtwoarrays.TwoArrays(tt.inK, tt.inA, tt.inB)

			assert.Equal(t, want, got)
		})
	}
}
