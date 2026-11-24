# Time Based Key-Value Store

> NeetCode 150 · Binary Search — per-key history slices, sorted for free by the strictly-increasing timestamp contract; Get is one upper-bound search. Go, stdlib only, tested.

## Problem

Design a store where `Set(key, value, timestamp)` appends and `Get(key, timestamp)` returns the value with the *largest timestamp not exceeding* the query — "" if nothing that old exists. Versioned configuration in miniature.

## Approach

The problem's quiet gift is the contract that each key's timestamps arrive strictly increasing — so appending keeps every history slice sorted with no sorting ever performed. Get then reduces to a predecessor query: `sort.Search` finds the first entry *newer* than the query, and the answer is its left neighbor (index 0 meaning "nothing that old"). Using the stdlib's search instead of a hand-rolled loop is the right call in Go — the custom predicate (`ts > timestamp`) expresses upper-bound directly. This history-slice-plus-predecessor-search shape is how MVCC version lookups and TSDB point queries actually work.

**Complexity:** Set amortized O(1); Get O(log h) for that key's history length.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
