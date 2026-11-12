package genz_test

import (
	"math"
	"testing"

	"github.com/brunty/genz"
)

func TestInts(t *testing.T) {
	for name, tc := range map[string]struct{ got, want int }{
		"HighKey":      {genz.HighKey, 1},
		"LowKey":       {genz.LowKey, 0},
		"Goat":         {genz.Goat, math.MaxInt},
		"SixSeven":     {genz.SixSeven, 67},
		"Extra":        {genz.Extra, 2},
		"Aura":         {genz.Aura, 1000},
		"NegativeAura": {genz.NegativeAura, -1000},
	} {
		if tc.got != tc.want {
			t.Errorf("%s = %d, want %d", name, tc.got, tc.want)
		}
	}
}

func TestMid(t *testing.T) {
	if genz.Mid != 0.5 {
		t.Errorf("Mid = %v, want 0.5", genz.Mid)
	}
	if !(genz.LowKey < genz.Mid && genz.Mid < genz.HighKey) {
		t.Errorf("Mid = %v, want it between LowKey (%d) and HighKey (%d)", genz.Mid, genz.LowKey, genz.HighKey)
	}
}
