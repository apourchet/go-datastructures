package ringbuffer

type RingBuffer struct {
	array []int
	size  int
	head  int
	tail  int
}

func Construct(size int) *RingBuffer {
	if size == 0 {
		size = 16
	}
	r := RingBuffer{}
	r.array = make([]int, size)
	r.size = size
	r.head = 0
	r.tail = 0
	return &r
}

func (r *RingBuffer) Push(objectPointer int) bool {
	r.array[r.tail] = objectPointer
	r.tail++
	if r.tail == r.size {
		r.tail = 0
	}
	return true
}

func (r *RingBuffer) Poll() (int, bool) {
	if r.tail == r.head {
		return 0, false
	}
	ret := r.array[r.head]
	r.head++
	if r.head == r.size {
		r.head = 0
	}
	return ret, true
}

func (r *RingBuffer) Peek() (int, bool) {
	if r.tail == r.head {
		return 0, false
	}
	ret := r.array[r.head]
	return ret, true
}
