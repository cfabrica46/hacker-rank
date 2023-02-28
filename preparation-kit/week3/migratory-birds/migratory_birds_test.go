package migratorybirds_test

import (
	"testing"

	"migratorybirds"

	"gotest.tools/assert"
)

func TestMigratoryBirds(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   []int32
		out  int32
	}{
		{
			name: "1",
			in:   []int32{1, 1, 2, 2, 3},
			out:  1,
		},
		{
			name: "2",
			in:   []int32{1, 4, 4, 4, 5, 3},
			out:  4,
		},
		{
			name: "1",
			in:   []int32{5, 1, 5},
			out:  5,
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := migratorybirds.MigratoryBirds(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
