package main

import "testing"

func TestTimeMap(t *testing.T) {
	m := NewTimeMap()
	m.Set("foo", "bar", 1)
	if got := m.Get("foo", 1); got != "bar" {
		t.Errorf("Get(foo,1) = %q", got)
	}
	if got := m.Get("foo", 3); got != "bar" {
		t.Errorf("Get(foo,3) = %q, want bar (latest ≤ 3)", got)
	}
	m.Set("foo", "bar2", 4)
	if got := m.Get("foo", 4); got != "bar2" {
		t.Errorf("Get(foo,4) = %q", got)
	}
	if got := m.Get("foo", 5); got != "bar2" {
		t.Errorf("Get(foo,5) = %q", got)
	}
	if got := m.Get("foo", 0); got != "" {
		t.Errorf("Get(foo,0) = %q, want empty (before first set)", got)
	}
	if got := m.Get("missing", 9); got != "" {
		t.Errorf("Get(missing) = %q, want empty", got)
	}
}
