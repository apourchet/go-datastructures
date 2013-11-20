package binarytree

type node struct {
	value       int
	left, right *node
}

type CompareFunction func(o1, o2 int) int

func (bt *node) insert(object int, compare CompareFunction) bool {
	c := compare(object, bt.value)
	if c < 0 {
		if bt.left == nil {
			bt.left = &node{object, nil, nil}
		} else {
			bt.left.insert(object, compare)
		}
	} else if c > 0 {
		if bt.right == nil {
			bt.right = &node{object, nil, nil}
		} else {
			bt.right.insert(object, compare)
		}
	}
	return true
}

func (bt *node) remove(object int, compare CompareFunction) *node {
	c := compare(object, bt.value)
	if c == 0 {
		if bt.right == nil {
			return bt.left
		}
		min := bt.right.min()
		bt.right = bt.right.removeMin()
		bt.value = min
	} else if c < 0 {
		if bt.left == nil {
			// Not found
			return bt
		}
		bt.left = bt.left.remove(object, compare)
	} else if c > 0 {
		if bt.right == nil {
			// Not found
			return bt
		}
		bt.right = bt.right.remove(object, compare)
	}
	return bt
}

func (bt *node) contains(object int, compare CompareFunction) bool {
	c := compare(object, bt.value)
	if c == 0 {
		return true
	} else if c < 0 {
		return bt.left != nil && bt.left.contains(object, compare)
	}
	return bt.right != nil && bt.right.contains(object, compare)

}

func (bt *node) max() int {
	if bt.right == nil {
		return bt.value
	}
	return bt.right.max()
}

func (bt *node) min() int {
	if bt.left == nil {
		return bt.value
	}
	return bt.left.min()
}

func (bt *node) removeMax() *node {
	if bt.right == nil {
		return bt.left
	}
	bt.right = bt.right.removeMax()
	return bt
}

func (bt *node) removeMin() *node {
	if bt.left == nil {
		return bt.right
	}
	bt.left = bt.left.removeMin()
	return bt
}

// Public functions
type BinaryTree struct {
	head    *node
	count   int
	Compare CompareFunction
}

func Construct(compare CompareFunction) *BinaryTree {
	bt := BinaryTree{}
	bt.count = 0
	bt.head = nil
	bt.Compare = compare
	return &bt
}

func (bt *BinaryTree) Insert(object int) {
	if bt.head == nil {
		bt.head = &node{object, nil, nil}
	} else {
		bt.head.insert(object, bt.Compare)
	}
	bt.count++
}

func (bt *BinaryTree) Remove(object int) bool {
	if bt.head == nil {
		return false
	}
	bt.head.remove(object, bt.Compare)
	bt.count--
	return true
}

func (bt *BinaryTree) Contains(object int) bool {
	if bt.head == nil {
		return false
	}
	return bt.head.contains(object, bt.Compare)
}

func (bt *BinaryTree) Max() (int, bool) {
	if bt.head == nil {
		return 0, false
	}
	return bt.head.max(), true
}

func (bt *BinaryTree) Min() (int, bool) {
	if bt.head == nil {
		return 0, false
	}
	return bt.head.min(), true
}
