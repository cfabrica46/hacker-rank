package final_test

import (
	"testing"

	"final/final"

	"gotest.tools/assert"
)

func TestPangrams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   [][]int32
		out  int32
	}{
		{
			name: "Revers 2x2",
			in: [][]int32{
				{1, 2},
				{3, 4},
			},
			out: 4,
		},
		{
			name: "Revers 4x4",
			in: [][]int32{
				{1, 2, 6, 7},
				{3, 4, 9, 11},
				{1, 2, 6, 7},
				{3, 4, 9, 11},
			},
			out: 40,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := final.FlippingMatrix(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
