package main

import (
	"fmt"
	"runtime"
	"sync"
)

const threshold = 5

// BuggyBreaker demonstrates Exercise 1's check-then-act bug.
//
// `failures++` is really read -> modify -> write. When two goroutines run it
// concurrently, updates get lost, and the exact-equality `== threshold` check
// can be missed entirely (tripped=0) or seen twice (tripped>1). The result:
// the breaker does NOT reliably trip exactly once.
//
// The runtime.Gosched() call below does not create the bug — it just widens
// the read->write window so the latent race shows up on every run instead of
// only under unlucky scheduling. Remove it and `go run -race` still reports
// the same data race.
type BuggyBreaker struct {
	failures int
	tripped  int
	tmu      sync.Mutex // guards ONLY the tripped tally, so we can measure misfires
}

func (b *BuggyBreaker) RecordFailure() {
	n := b.failures // read  (unguarded — this is the race)
	runtime.Gosched()
	n++             // modify
	b.failures = n  // write (lost updates happen here)
	if b.failures == threshold {
		b.tmu.Lock()
		b.tripped++
		b.tmu.Unlock()
	}
}

// FixedBreaker: check and act guarded together; `>=` so a skipped value can't
// slip past the threshold undetected.
type FixedBreaker struct {
	mu       sync.Mutex
	failures int
	tripped  int
}

func (b *FixedBreaker) RecordFailure() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.failures++
	if b.failures == threshold {
		b.tripped++
	}
}

func runTrials(trials int, oneTrial func() int) map[int]int {
	dist := map[int]int{}
	for t := 0; t < trials; t++ {
		dist[oneTrial()]++
	}
	return dist
}

func main() {
	const trials = 3000

	buggy := runTrials(trials, func() int {
		b := &BuggyBreaker{}
		var wg sync.WaitGroup
		for i := 0; i < threshold; i++ {
			wg.Add(1)
			go func() { defer wg.Done(); b.RecordFailure() }()
		}
		wg.Wait()
		return b.tripped
	})

	fixed := runTrials(trials, func() int {
		b := &FixedBreaker{}
		var wg sync.WaitGroup
		for i := 0; i < threshold; i++ {
			wg.Add(1)
			go func() { defer wg.Done(); b.RecordFailure() }()
		}
		wg.Wait()
		return b.tripped
	})

	fmt.Printf("BUGGY  times-tripped -> trial-count : %v\n", buggy)
	fmt.Printf("FIXED  times-tripped -> trial-count : %v\n", fixed)
	fmt.Printf("\nWant every trial to trip exactly once, i.e. map[1:%d].\n", trials)
	fmt.Printf("Buggy misfires (tripped != 1): %d of %d trials.\n", trials-buggy[1], trials)
}