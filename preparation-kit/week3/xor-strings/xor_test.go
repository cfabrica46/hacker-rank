package xor_test

import (
	"testing"

	"xor"

	"gotest.tools/assert"
)

func TestPangrams(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		inA  string
		inB  string
		out  string
	}{
		{
			name: "1",
			inA:  "10101",
			inB:  "00101",
			out:  "10000",
		},
	} {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			want := tt.out

			got := xor.XOR(tt.inA, tt.inB)

			assert.Equal(t, want, got)
		})
	}
}
