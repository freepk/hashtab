package hashtab

import (
	"sync/atomic"
)

type entry struct {
	key   uint64
	value uint64
}

type HashTab struct {
	mask    uint64
	entries []entry
}

func bitmask(n int) int {
	n |= (n >> 1)
	n |= (n >> 2)
	n |= (n >> 4)
	n |= (n >> 8)
	n |= (n >> 16)
	return n
}

func NewHashTab(num int) *HashTab {
	mask := uint64(bitmask(num - 1))
	entries := make([]entry, (mask + 1))
	return &HashTab{mask: mask, entries: entries}
}

func (h *HashTab) Size() int {
	return len(h.entries)
}

func (h *HashTab) GetOrSet(k, v uint64) (uint64, bool) {
	var i, t uint64
	var e *entry

	i = k
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint64(&e.key)
		if t == k {
			return atomic.LoadUint64(&e.value), true
		}
		if t == 0 {
			if atomic.CompareAndSwapUint64(&e.key, 0, k) {
				atomic.StoreUint64(&e.value, v)
				return v, false
			}
		}
		i++
	}
}

func (h *HashTab) Set(k, v uint64) {
	var i, t uint64
	var e *entry

	i = k
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint64(&e.key)
		if t == k {
			atomic.StoreUint64(&e.value, v)
			return
		}
		if t == 0 {
			if atomic.CompareAndSwapUint64(&e.key, 0, k) {
				atomic.StoreUint64(&e.value, v)
				return
			}
		}
		i++
	}
}

func (h *HashTab) Get(k uint64) (uint64, bool) {
	var i, t uint64
	var e *entry

	i = k
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint64(&e.key)
		if t == k {
			return atomic.LoadUint64(&e.value), true
		}
		if t == 0 {
			return 0, false
		}
		i++
	}
}

func (h *HashTab) Del(k uint64) {
	var i, t uint64
	var e *entry

	i = k
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint64(&e.key)
		if t == k {
			atomic.StoreUint64(&e.value, 0)
			atomic.StoreUint64(&e.key, 0)
			return
		}
		if t == 0 {
			return
		}
		i++
	}
}
