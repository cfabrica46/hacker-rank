package marsexploration_test

import (
	"testing"

	"marsexploration"

	"gotest.tools/assert"
)

func TestMarsExploration(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   string
		out  int32
	}{
		/* {
			name: "3 letters changed",
			in:   "SOSSPSSQSSOR",
			out:  int32(3),
		}, */
		{
			name: "1 letters changed",
			in:   "SOSSOT",
			out:  int32(1),
		},
		/* {
			name: "20 letters changed",
			in:   "OOSDSSOSOSWEWSOSOSOSOSOSOSSSSOSOSOSS",
			out:  int32(20),
		}, */
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := marsexploration.MarsExploration(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
