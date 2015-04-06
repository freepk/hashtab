package hashtab

import (
	"testing"
)

func TestHashTab(t *testing.T) {
	h := NewHashTab(10)
	h.Set(100, 1000)
	h.Set(101, 1001)
	if h.Get(100) != 1000 {
		t.Fail()
	}
	if h.Get(101) != 1001 {
		t.Fail()
	}
	if h.Get(50) != 0 {
		t.Fail()
	}
}
