package genz_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/brunty/genz"
)

func TestErrors(t *testing.T) {
	for name, tc := range map[string]struct {
		err  error
		want string
	}{
		"ErrCooked":     {genz.ErrCooked, "it's cooked"},
		"ErrSus":        {genz.ErrSus, "that's sus"},
		"ErrIck":        {genz.ErrIck, "the ick"},
		"ErrTouchGrass": {genz.ErrTouchGrass, "touch grass"},
		"ErrSkillIssue": {genz.ErrSkillIssue, "skill issue"},
		"ErrGhosted":    {genz.ErrGhosted, "ghosted: context deadline exceeded"},
		"ErrNoRizz":     {genz.ErrNoRizz, "no rizz"},
		"ErrRedFlag":    {genz.ErrRedFlag, "red flag"},
		"ErrRatio":      {genz.ErrRatio, "ratio"},
		"ErrCaughtIn4K": {genz.ErrCaughtIn4K, "caught in 4K"},
		"ErrLeftOnRead": {genz.ErrLeftOnRead, "left on read"},
		"ErrFumbled":    {genz.ErrFumbled, "fumbled the bag"},
		"ErrBigYikes":   {genz.ErrBigYikes, "big yikes"},
		"ErrNotTheVibe": {genz.ErrNotTheVibe, "not the vibe"},
	} {
		if tc.err.Error() != tc.want {
			t.Errorf("%s = %q, want %q", name, tc.err.Error(), tc.want)
		}
	}
}

func TestErrorsIs(t *testing.T) {
	wrapped := errors.Join(errors.New("deploy failed"), genz.ErrCooked)
	if !errors.Is(wrapped, genz.ErrCooked) {
		t.Error("errors.Is(wrapped, ErrCooked) = false, want true")
	}
	if errors.Is(wrapped, genz.ErrSus) {
		t.Error("errors.Is(wrapped, ErrSus) = true, want false")
	}
}

func TestErrGhostedIsTimeout(t *testing.T) {
	if !errors.Is(genz.ErrGhosted, context.DeadlineExceeded) {
		t.Error("errors.Is(ErrGhosted, context.DeadlineExceeded) = false, want true")
	}
	if !errors.Is(fmt.Errorf("fetch: %w", genz.ErrGhosted), context.DeadlineExceeded) {
		t.Error("errors.Is(wrapped ErrGhosted, context.DeadlineExceeded) = false, want true")
	}
}

func TestVibeCheck(t *testing.T) {
	if err := genz.VibeCheck(true); err != nil {
		t.Errorf("VibeCheck(true) = %v, want nil", err)
	}
	if err := genz.VibeCheck(false); !errors.Is(err, genz.ErrNotTheVibe) {
		t.Errorf("VibeCheck(false) = %v, want ErrNotTheVibe", err)
	}
}
