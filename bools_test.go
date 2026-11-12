package genz_test

import (
	"testing"

	"github.com/brunty/genz"
)

func TestTrue(t *testing.T) {
	for name, v := range map[string]bool{
		"NoCap":   genz.NoCap,
		"Bet":     genz.Bet,
		"Fr":      genz.Fr,
		"FrFr":    genz.FrFr,
		"Valid":   genz.Valid,
		"W":       genz.W,
		"Facts":   genz.Facts,
		"Based":   genz.Based,
		"Slay":    genz.Slay,
		"Ate":     genz.Ate,
		"Bussin":  genz.Bussin,
		"Gucci":   genz.Gucci,
		"SayLess": genz.SayLess,
		"Deadass": genz.Deadass,
		"Fire":    genz.Fire,
		"Lit":     genz.Lit,
		"Slaps":   genz.Slaps,
		"Periodt": genz.Periodt,
	} {
		if !v {
			t.Errorf("%s = %v, want true", name, v)
		}
	}
}

func TestFalse(t *testing.T) {
	for name, v := range map[string]bool{
		"Cap":       genz.Cap,
		"L":         genz.L,
		"Delulu":    genz.Delulu,
		"Sus":       genz.Sus,
		"Cringe":    genz.Cringe,
		"Flop":      genz.Flop,
		"NPC":       genz.NPC,
		"Basic":     genz.Basic,
		"Washed":    genz.Washed,
		"Unserious": genz.Unserious,
	} {
		if v {
			t.Errorf("%s = %v, want false", name, v)
		}
	}
}
