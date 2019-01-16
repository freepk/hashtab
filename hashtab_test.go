package hashtab

import (
	"testing"
	"github.com/spaolacci/murmur3"
)

func TestHashTab(t *testing.T) {
	h := NewHashTab(10)
	h.Set(100, 1000)
	h.Set(101, 1001)
	if v, ok := h.Get(100); !ok || v != 1000 {
		t.Fail()
	}
	if v, ok := h.Get(101); !ok || v != 1001 {
		t.Fail()
	}
	if v, ok := h.GetOrSet(100, 1111); !ok || v != 1000 {
		t.Log(ok, v)
	}
	h.Del(105)
	h.Del(100)
	if v, ok := h.Get(100); ok || v != 0 {
		t.Fail()
	}
	if v, ok := h.GetOrSet(100, 1111); ok || v != 1111 {
		t.Log(ok, v)
	}
	if v, ok := h.Get(50); ok || v != 0 {
		t.Fail()
	}
}

func TestGetOrSet(t *testing.T) {
	h := NewHashTab(16)
	{
		hash := murmur3.Sum64([]byte{0,1,2,3})
		t.Log(h.GetOrSet(hash, 100))
	}
	{
		hash := murmur3.Sum64([]byte{0,1,2,3})
		t.Log(h.GetOrSet(hash, 100))
	}
	{
		hash := murmur3.Sum64([]byte{0,1,2,3})
		t.Log(h.GetOrSet(hash, 100))
	}
}