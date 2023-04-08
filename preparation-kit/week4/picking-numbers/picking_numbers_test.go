package pickingnumbers_test

import (
	"testing"

	"pickingnumbers"

	"github.com/stretchr/testify/assert"
)

func TestPickingNumbers(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   []int32
		out  int32
	}{
		/* {
			name: "Test",
			in:   []int32{1, 1, 2, 2, 4, 4, 5, 5, 5},
			out:  5,
		}, */
		{
			name: "Test",
			in:   []int32{4, 6, 5, 3, 3, 1},
			out:  3,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := pickingnumbers.PickingNumbers(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
