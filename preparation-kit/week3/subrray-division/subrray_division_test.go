package subrraydivision_test

import (
	"testing"

	"subrraydivision"

	"gotest.tools/assert"
)

func TestPangrams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		inS  []int32
		inD  int32
		inM  int32
		out  int32
	}{
		{
			name: "1",
			inS:  []int32{2, 2, 1, 3, 2},
			inD:  4,
			inM:  2,
			out:  2,
		},
		{
			name: "2",
			inS:  []int32{5, 2, 2, 1, 5, 3, 2},
			inD:  9,
			inM:  3,
			out:  2,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := subrraydivision.Birthday(tt.inS, tt.inD, tt.inM)

			assert.Equal(t, want, got)
		})
	}
}
