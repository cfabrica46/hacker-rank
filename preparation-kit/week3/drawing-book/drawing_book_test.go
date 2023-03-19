package drawingbook_test

import (
	"testing"

	"drawingbook"

	"github.com/stretchr/testify/assert"
)

func TestPageCount(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		inN  int32
		inP  int32
		out  int32
	}{
		{
			name: "",
			inN:  6,
			inP:  5,
			out:  1,
		},
		{
			name: "Equal",
			inN:  4,
			inP:  4,
			out:  0,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := drawingbook.PageCount(tt.inN, tt.inP)

			assert.Equal(t, want, got)
		})
	}
}
