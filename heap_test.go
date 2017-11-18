package dsalgo

import (
	"reflect"
	"testing"
)

func Test_Heap(t *testing.T) {
	h := NewHeap(func(x interface{}, y interface{}) bool {
		return x.(int) < y.(int)
	})

	h.Push(7, 3) // [3, 7]
	if a := h.Peek().(int); a != 3 {
		t.Errorf("Top element should be %d, but got %v", 3, a)
	}

	h.Push(11, 4, 4, 5) // [3, 5, 7, 11]
	if !reflect.DeepEqual(h.Peek(), h.Pop()) {
		t.Errorf("Peek and Pop should return the same thing")
	}

	// Test sorted order
	expected := []int{4, 4, 5, 7, 11}
	for _, e := range expected {
		if a := h.Pop(); !reflect.DeepEqual(a, e) {
			t.Errorf("Heap ordering wrong: expected %d but got %v", e, a)
		}
	}
}
