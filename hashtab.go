package hashtab

import (
	"sync/atomic"
)

type entry struct {
	key   uint64
	value uint64
}

type HashTab struct {
	size    uint64
	mask    uint64
	entries []entry
}

const MaxPower = 32

func NewHashTab(power uint8) *HashTab {
	if power > MaxPower {
		return nil
	}
	size := uint64(1 << power)
	mask := size - 1
	entries := make([]entry, size)
	return &HashTab{size: size, mask: mask, entries: entries}
}

func (h *HashTab) Size() uint64 {
	return h.size
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
