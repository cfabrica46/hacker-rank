package betweentwosets_test

import (
	"testing"

	"betweentwosets"

	"gotest.tools/assert"
)

func TestGetTotalX(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		inA  []int32
		inB  []int32
		out  int32
	}{
		/* {
			name: "1",
			inA:  []int32{2, 6},
			inB:  []int32{24, 36},
			out:  2,
		}, */
		{
			name: "2",
			inA:  []int32{2, 4},
			inB:  []int32{16, 32, 96},
			out:  3,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := betweentwosets.GetTotalX(tt.inA, tt.inB)

			assert.Equal(t, want, got)
		})
	}
}
