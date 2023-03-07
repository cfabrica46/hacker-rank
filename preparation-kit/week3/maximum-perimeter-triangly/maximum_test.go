package maximumperimetertriangly_test

import (
	"testing"

	"maximumperimetertriangly"

	"github.com/stretchr/testify/assert"
)

func TestMaximumPerimeterTriangle(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   []int32
		out  []int32
	}{
		{
			name: "1",
			in:   []int32{1, 2, 3, 4, 5, 10},
			out:  []int32{3, 4, 5},
		},
		{
			name: "2",
			in:   []int32{1, 2, 3},
			out:  []int32{-1},
		},
		{
			name: "3",
			in:   []int32{3, 9, 2, 15, 3},
			out:  []int32{2, 3, 3},
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := maximumperimetertriangly.MaximumPerimeterTriangle(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
