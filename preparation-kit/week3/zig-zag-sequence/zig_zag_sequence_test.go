package zigzagsequence_test

import (
	"testing"

	"zigzagsequence"

	"github.com/stretchr/testify/assert"
)

func TestFindZigZagSequence(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   []int32
		out  []int32
	}{
		{
			name: "1",
			in:   []int32{1, 2, 3, 4, 5, 6, 7},
			out:  []int32{1, 2, 3, 7, 6, 5, 4},
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := zigzagsequence.FindZigZagSequence(tt.in, len(tt.in))

			assert.Equal(t, want, got)
		})
	}
}
