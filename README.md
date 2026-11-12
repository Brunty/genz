# Gen Z Go

> Silly things.

A Go port of [genz-php](https://github.com/brunty/genz-php).

# Installation

```bash
go get github.com/brunty/genz
```

# Documentation

Telling the truth?
```go
return genz.NoCap
```

Lying?
```go
return genz.Cap
```

Agreeing?
```go
return genz.Bet // or genz.Fr, genz.FrFr, genz.Valid, genz.Facts, genz.Based,
// genz.SayLess, genz.Deadass, genz.Periodt
```

Winning or losing?
```go
return genz.W // or genz.L
```

Nailed it?
```go
return genz.Slay // or genz.Ate, genz.Bussin, genz.Gucci, genz.Fire, genz.Lit, genz.Slaps
```

Not buying it?
```go
return genz.Sus // or genz.Delulu
```

Didn't land?
```go
return genz.Cringe // or genz.Flop, genz.NPC, genz.Basic, genz.Washed, genz.Unserious
```

Turning it up or down?
```go
var volume float64 = genz.Extra // 2
volume = genz.HighKey           // 1
volume = genz.Mid               // 0.5
volume = genz.LowKey            // 0
```

The best, or just a meme?
```go
score := genz.Goat        // math.MaxInt
score = genz.SixSeven     // 67
score = genz.Aura         // 1000
score = genz.NegativeAura // -1000
```

Something went wrong?
```go
return genz.ErrCooked // or genz.ErrSus, genz.ErrIck, genz.ErrTouchGrass,
// genz.ErrSkillIssue, genz.ErrGhosted, genz.ErrNoRizz, genz.ErrRedFlag, genz.ErrRatio,
// genz.ErrCaughtIn4K, genz.ErrLeftOnRead, genz.ErrFumbled, genz.ErrBigYikes, genz.ErrNotTheVibe

if errors.Is(err, genz.ErrCooked) {
	// it's over
}

// ErrGhosted wraps context.DeadlineExceeded, so timeout checks catch it too
if errors.Is(genz.ErrGhosted, context.DeadlineExceeded) {
	// true
}
```

Describing the vibe?
```go
genz.ItsGiving("main character") // "it's giving main character"
genz.Era("gym")                  // "in my gym era"
genz.RentFree("that song")       // "that song is living rent free in my head"
```

Checking the vibe?
```go
if err := genz.VibeCheck(passed); err != nil {
	return err // genz.ErrNotTheVibe
}
```
