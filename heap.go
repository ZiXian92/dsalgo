package dsalgo

// Heap defines the min-heap data structure.
type Heap interface {
	Push(x ...interface{})
	Peek() interface{}
	Pop() interface{}
}

type heap struct {
	arr  []interface{}
	less func(interface{}, interface{}) bool
}

// NewHeap creates and initializes a new Heap container.
// To not make lessFn cause panic, it is up to the programmer to
// make sure that only the intended types are passed to the heap
// throughout its lifetime.
func NewHeap(lessFn func(interface{}, interface{}) bool) Heap {
	return &heap{arr: make([]interface{}, 0, 1), less: lessFn}
}

// Push adds the new elements to the heap
// and shuffles to maintain min-heap property.
func (h *heap) Push(x ...interface{}) {
	oldLen := len(h.arr)
	h.arr = append(h.arr, x...)
	newLen := len(h.arr)
	for i := oldLen; i < newLen; i++ {
		h.up(i)
	}
}

// Peek returns but does not remove the smallest element in the heap.
func (h *heap) Peek() interface{} {
	if len(h.arr) == 0 {
		return nil
	}
	return h.arr[0]
}

// Pop removes and returns the smallest element in the heap.
// Returns nil if the heap is empty.
func (h *heap) Pop() interface{} {
	numElem := len(h.arr)
	if numElem == 0 {
		return nil
	}
	x := h.arr[0]               // Take out and hold the return element
	h.arr[0] = h.arr[numElem-1] // Use the last element to replace the top element
	h.arr = h.arr[:numElem-1]   // Update slice length
	h.down(0)                   // Bubble down the new top element if needed
	return x
}

func (h *heap) up(i int) {
	for i > 0 { // While not at the top
		j := (i - 1) >> 1                // Parent index
		if !h.less(h.arr[i], h.arr[j]) { // Not less than parent, done
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i] // Swap
		i = j                                   // Update index of the element to shift up
	}
}

func (h *heap) down(i int) {
	numElem := len(h.arr)
	maxLimit := numElem >> 1
	for i < maxLimit { // Ensure that there is at least a left child
		lj, rj := (i<<1)+1, (i<<1)+2 // Get child indices

		// Pick the smaller of the 2 children to compare with current element
		// so that we can handle many cases with only 1 version of code
		if rj < numElem && h.less(h.arr[rj], h.arr[lj]) {
			lj, rj = rj, lj
		}

		// Cannot even swap with the smaller of the 2 children, done
		if !h.less(h.arr[lj], h.arr[i]) {
			break
		}

		h.arr[i], h.arr[lj] = h.arr[lj], h.arr[i] // Swap
		i = lj                                    // Update the new index of the element
	}
}
