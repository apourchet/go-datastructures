package stack

type StackItem struct {
	value int
	next  *StackItem
}

type Stack struct {
	head *StackItem
	size int
}

func Construct() *Stack {
	stack := Stack{}
	stack.size = 0
	stack.head = nil
	return &stack
}

func (s *Stack) Push(objectPointer int) bool {
	if s.size == 0 {
		s.head = &StackItem{objectPointer, nil}
	} else {
		snd := *s.head
		s.head = &StackItem{objectPointer, &snd}
	}
	s.size++
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	i := s.head
	s.head = s.head.next
	s.size--
	return i.value, true
}

func (s *Stack) Peek() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	return s.head.value, true
}
