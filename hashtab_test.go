package hashtab

import (
	"testing"
)

func TestHashTab(t *testing.T) {
	h := NewHashTab(10)
	if h == nil {
		t.Fail()
	}
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
