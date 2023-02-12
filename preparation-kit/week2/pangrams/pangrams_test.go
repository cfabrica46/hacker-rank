package pangrams_test

import (
	"testing"

	"pangrams"

	"github.com/stretchr/testify/assert"
)

func TestPangrams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "Is Pangram",
			in:   "We promptly judged antique ivory buckles for the next prize",
			out:  "pangram",
		},
		{
			name: "Isn't Pangram",
			in:   "We promptly judged antique ivory buckles for the prize",
			out:  "not pangram",
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := pangrams.Pangrams(tt.in)

			assert.Equal(t, want, got)
		})
	}
}
