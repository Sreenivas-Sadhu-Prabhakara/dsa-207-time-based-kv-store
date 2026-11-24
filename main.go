package main

import (
	"fmt"
	"sort"
)

type entry struct {
	ts  int
	val string
}

// TimeMap stores (key, value, timestamp) triples where timestamps per key
// arrive strictly increasing (so each key's history is sorted by
// construction), and answers "the latest value at or before t".
type TimeMap struct {
	data map[string][]entry
}

// NewTimeMap returns an empty store.
func NewTimeMap() *TimeMap {
	return &TimeMap{data: map[string][]entry{}}
}

// Set records value for key at the given timestamp.
func (m *TimeMap) Set(key, value string, timestamp int) {
	m.data[key] = append(m.data[key], entry{timestamp, value})
}

// Get returns the value whose timestamp is the largest one ≤ timestamp,
// or "" when the key has no value that old.
func (m *TimeMap) Get(key string, timestamp int) string {
	history := m.data[key]
	// First index with ts > timestamp; the answer sits just before it.
	i := sort.Search(len(history), func(i int) bool { return history[i].ts > timestamp })
	if i == 0 {
		return ""
	}
	return history[i-1].val
}

func main() {
	m := NewTimeMap()
	m.Set("foo", "bar", 1)
	fmt.Println(m.Get("foo", 1), m.Get("foo", 3))
	m.Set("foo", "bar2", 4)
	fmt.Println(m.Get("foo", 4), m.Get("foo", 5))
}
