package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const bufferTime = 10 * time.Millisecond // Add buffer to ensure the reap loop has ample time
	cache := NewCache(baseTime)              // Initialize cache with quick reap interval

	// Step 1: Add an entry to the cache
	cache.Add("https://example.com", []byte("testdata"))

	// Step 2: Verify the entry exists immediately after adding
	val, ok := cache.Get("https://example.com")
	if !ok || string(val) != "testdata" {
		t.Errorf("expected key to exist with correct value immediately after adding")
		return
	}

	// Step 3: Wait for more than the reap interval (baseTime + bufferTime) to allow reaping
	time.Sleep(baseTime + bufferTime)

	// Step 4: Verify the entry has been reaped (should no longer exist)
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected key to be reaped and not found in cache after interval")
		return
	}
}
