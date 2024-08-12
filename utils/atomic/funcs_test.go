package atomic

import (
	"flag"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestUint64Sub(t *testing.T) {
	t.Run("sanity", func(t *testing.T) {
		exp := uint64(500000)
		got := atomic.Uint64{}
		got.Add(exp)
		sub := uint64(250000)
		exp -= sub
		Uint64Sub(&got, sub)
		require.Equal(t, exp, got.Load())
	})
	t.Run("rapid", func(t *testing.T) {
		_ = flag.Set("rapid.checks", "50")
		_ = flag.Set("rapid.steps", "50")

		rapid.Check(t, func(r *rapid.T) {
			exp := uint64(0)
			got := atomic.Uint64{}
			r.Repeat(map[string]func(r *rapid.T){
				"": func(r *rapid.T) {
					require.Equal(t, exp, got.Load())
				},
				"add": func(r *rapid.T) {
					d := rapid.Uint64().Draw(r, "d")
					exp += d
					got.Add(d)
				},
				"sub": func(r *rapid.T) {
					d := rapid.Uint64().Draw(r, "d")
					exp -= d
					Uint64Sub(&got, d)
				},
			})
		})
	})
}
