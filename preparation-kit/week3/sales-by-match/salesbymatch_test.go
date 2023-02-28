package salesbymatch_test

import (
	"testing"

	"salesbymatch"

	"gotest.tools/assert"
)

func TestSalesByMatch(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name  string
		inN   int32
		inArr []int32
		out   int32
	}{
		{
			name:  "1",
			inN:   9,
			inArr: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			out:   3,
		},
		/* {
			name:  "2",
			inN:   10,
			inArr: []int32{2, 1, 3},
			out:   10,
		}, */
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := salesbymatch.SockMerchant(tt.inN, tt.inArr)

			assert.Equal(t, want, got)
		})
	}
}
