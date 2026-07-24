// A concurrent-safe in-memory key-value store with TTL-based expiration
// and a background janitor for active cleanup.
//
// Run:   go run main.go
// Race:  go run -race main.go   (proves no data races under concurrency)
package main

import (
	"fmt"
	"sync"
	"time"
)

// ----------------------------------------------------------------------------
// Data model
// ----------------------------------------------------------------------------

// entry is the internal representation of a stored value.
// We store the ABSOLUTE expiry timestamp, so "is it expired?" is a single
// comparison against time.Now() — no elapsed-time bookkeeping required.
type entry struct {
	value     any
	expiresAt time.Time
	hasTTL    bool // false => this key never expires
}

func (e entry) isExpired(now time.Time) bool {
	return e.hasTTL && now.After(e.expiresAt)
}

// ----------------------------------------------------------------------------
// Store
// ----------------------------------------------------------------------------

type Store struct {
	mu     sync.RWMutex // read-heavy workload => RWMutex, not plain Mutex
	data   map[string]entry
	stopCh chan struct{}  // signals the janitor to exit
	wg     sync.WaitGroup // lets Close() block until the janitor is gone
	once   sync.Once      // makes Close() idempotent (no double-close panic)
}

// New creates a store. If cleanupInterval > 0, it launches the background
// janitor that actively sweeps expired keys on that interval.
func New(cleanupInterval time.Duration) *Store {
	s := &Store{
		data:   make(map[string]entry),
		stopCh: make(chan struct{}),
	}
	if cleanupInterval > 0 {
		s.wg.Add(1)
		go s.janitor(cleanupInterval)
	}
	return s
}

// Set stores a value. ttl <= 0 means the key never expires.
func (s *Store) Set(key string, value any, ttl time.Duration) {
	// Build the entry BEFORE locking — no reason to hold the write lock
	// while doing arithmetic. Keep critical sections short.
	e := entry{value: value}
	if ttl > 0 {
		e.expiresAt = time.Now().Add(ttl)
		e.hasTTL = true
	}

	s.mu.Lock()
	s.data[key] = e
	s.mu.Unlock()
}

// Get returns the value and whether it was present-and-live.
// This is where LAZY (passive) expiration happens: an expired key is
// treated as absent and deleted on the way out.
func (s *Store) Get(key string) (any, bool) {
	now := time.Now()

	s.mu.RLock()
	e, ok := s.data[key]
	s.mu.RUnlock()

	if !ok {
		return nil, false
	}
	if !e.isExpired(now) {
		return e.value, true
	}

	// Expired: delete it lazily. We held a READ lock above and now need a
	// WRITE lock, but Go's RWMutex can't upgrade atomically. So we release,
	// re-acquire, and RE-CHECK expiry under the write lock — because between
	// the two locks another goroutine may have re-Set this key with a fresh
	// value. Deleting blindly would destroy valid data.
	s.mu.Lock()
	if cur, ok := s.data[key]; ok && cur.isExpired(time.Now()) {
		delete(s.data, key)
	}
	s.mu.Unlock()

	return nil, false
}

// Delete removes a key unconditionally.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
}

// Len reports the current number of entries (including any not-yet-swept
// expired ones). Handy for tests/observability.
func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// ----------------------------------------------------------------------------
// Active expiration (background janitor)
// ----------------------------------------------------------------------------

// janitor runs until stopCh is closed. On each tick it sweeps expired keys.
// This is what bounds memory for keys that are set and never read again —
// lazy expiration alone would leak those forever.
func (s *Store) janitor(interval time.Duration) {
	defer s.wg.Done()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.cleanupExpired()
		case <-s.stopCh:
			return // clean exit; wg.Done fires via defer
		}
	}
}

// cleanupExpired does one full sweep. NOTE: O(n) scan under the write lock —
// fine for thousands of keys, a bottleneck for millions. Sharding and/or a
// min-heap of expiry times are the standard fixes at scale.
func (s *Store) cleanupExpired() {
	now := time.Now()

	s.mu.Lock()
	defer s.mu.Unlock()
	for k, e := range s.data {
		if e.isExpired(now) {
			delete(s.data, k)
		}
	}
}

// ----------------------------------------------------------------------------
// Graceful shutdown
// ----------------------------------------------------------------------------

// Close stops the janitor and blocks until it has actually exited.
// No leaked goroutines. Safe to call more than once.
func (s *Store) Close() {
	s.once.Do(func() {
		close(s.stopCh)
	})
	s.wg.Wait()
}

// ----------------------------------------------------------------------------
// Demo
// ----------------------------------------------------------------------------

func main() {
	store := New(500 * time.Millisecond) // janitor sweeps twice a second
	defer store.Close()

	store.Set("name", "alice", 0)                  // never expires
	store.Set("session", "xyz-123", 1*time.Second) // expires in 1s

	if v, ok := store.Get("session"); ok {
		fmt.Println("session:", v) // -> session: xyz-123
	}

	fmt.Println("keys before expiry:", store.Len()) // -> 2

	time.Sleep(1500 * time.Millisecond) // let it expire AND get swept

	if _, ok := store.Get("session"); !ok {
		fmt.Println("session expired and was cleaned up")
	}
	if v, ok := store.Get("name"); ok {
		fmt.Println("name still here:", v) // -> alice
	}

	fmt.Println("keys after sweep:", store.Len()) // -> 1 (janitor reclaimed it)

	// Quick concurrency smoke test — run with -race to check for data races.
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("k-%d", n)
			store.Set(key, n, 200*time.Millisecond)
			store.Get(key)
		}(i)
	}
	wg.Wait()
	fmt.Println("concurrent writes/reads completed without a crash")
}
