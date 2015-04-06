package hashtab

import (
	"sync/atomic"
)

type entry struct {
	key   uint32
	value uint32
}

type HashTab struct {
	mask    uint32
	entries []entry
}

func hash(v uint32) uint32 {
	v ^= v >> 16
	v *= 0x85ebca6b
	v ^= v >> 13
	v *= 0xc2b2ae35
	v ^= v >> 16
	return v
}

func NewHashTab(power uint32) *HashTab {
	return &HashTab{
		mask:    (1 << power) - 1,
		entries: make([]entry, (1 << power))}
}

func (h *HashTab) Set(k, v uint32) {
	var i, t uint32
	var e *entry

	i = hash(k)
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint32(&e.key)
		if t == k {
			atomic.StoreUint32(&e.value, v)
			return
		}
		if t == 0 {
			if atomic.CompareAndSwapUint32(&e.key, 0, k) {
				atomic.StoreUint32(&e.value, v)
				return
			}
		}
		i++
	}
}

func (h *HashTab) Get(k uint32) uint32 {
	var i, t uint32
	var e *entry

	i = hash(k)
	for {
		i &= h.mask
		e = &h.entries[i]
		t = atomic.LoadUint32(&e.key)
		if t == k {
			return atomic.LoadUint32(&e.value)
		}
		if t == 0 {
			return 0
		}
		i++
	}
}
