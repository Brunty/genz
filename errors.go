package genz

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrCooked     = errors.New("it's cooked")
	ErrSus        = errors.New("that's sus")
	ErrIck        = errors.New("the ick")
	ErrTouchGrass = errors.New("touch grass")
	ErrSkillIssue = errors.New("skill issue")
	ErrGhosted    = fmt.Errorf("ghosted: %w", context.DeadlineExceeded)
	ErrNoRizz     = errors.New("no rizz")
	ErrRedFlag    = errors.New("red flag")
	ErrRatio      = errors.New("ratio")
	ErrCaughtIn4K = errors.New("caught in 4K")
	ErrLeftOnRead = errors.New("left on read")
	ErrFumbled    = errors.New("fumbled the bag")
	ErrBigYikes   = errors.New("big yikes")
	ErrNotTheVibe = errors.New("not the vibe")
)

func VibeCheck(passed bool) error {
	if !passed {
		return ErrNotTheVibe
	}
	return nil
}
