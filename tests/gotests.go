package main

import (
	"../ArrayList"
	"../BinaryTree"
	"../BloomFilter"
	"../RingBuffer"
	"../Stack"
	"../Trie"
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	fmt.Println("Compiled fine")
	// arraylistTest()
	// stackTest()
	// ringbufferTest()
	// bloomfilterTest()
	// binarytreeTest()
	trieTest()
}

func arraylistTest() {
	// ArrayList testing
	fmt.Println("*******ARRAYLIST TEST*******")
	a := arraylist.Construct(0)
	fmt.Println(a.Get(0))
	fmt.Println(a.Insert(2))
	fmt.Println(a.Insert(3))
	fmt.Println(a.Insert(12))
	fmt.Println(a.Remove(0))
	fmt.Println(a.Get(0))
	fmt.Println(a.Get(1))
	fmt.Println(a.Get(2))
}

func stackTest() {
	// Stack testing
	fmt.Println("\n*******STACK TEST*******")
	s := stack.Construct()
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Push(4))
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	_, e := s.Pop()
	fmt.Println(e == false)
}

func ringbufferTest() {
	// RingBuffer testing
	fmt.Println("\n*******RINGBUFFER TEST*******")
	r := ringbuffer.Construct(4)
	fmt.Println(r.Push(1))
	fmt.Println(r.Push(2))
	fmt.Println(r.Push(3))
	fmt.Println(r.Push(4))
	fmt.Println(r.Push(5))
	fmt.Println(r.Push(6))
	fmt.Println(r.Peek())
	fmt.Println(r.Poll())
	fmt.Println(r.Poll())
}

func bloomfilterTest() {
	// BloomFilter testing
	fmt.Println("\n*******BLOOMFILTER TEST*******")
	bf := bloomfilter.Construct(500, 5, func(i, r int) uint {
		return uint((((i * i * i) % 64) + (r * r)) % 500)
	})
	fmt.Println(bf.Put(3))
	fmt.Println(bf.Put(100))
	fmt.Println(bf.Put(8))
	fmt.Println(bf.MightContain(3))
	fmt.Println(bf.MightContain(4) == false)
	fmt.Println(bf.MightContain(5) == false)
	fmt.Println(bf.MightContain(6) == false)
	fmt.Println(bf.MightContain(100))
	fmt.Println(bf.MightContain(8))
}

func binarytreeTest() {
	// BinaryTree testing
	fmt.Println("\n*******BINARYTREE TEST*******")
	bt := binarytree.Construct(func(o1, o2 int) int {
		if o1 == o2 {
			return 0
		} else if o1 < o2 {
			return -1
		}
		return 1
	})
	bt.Insert(13)
	bt.Insert(14)
	bt.Insert(15)
	bt.Insert(12)
	bt.Remove(12)
	bt.Remove(13)
	fmt.Println(bt.Max())
	fmt.Println(bt.Min())
}

func trieTest() {
	fmt.Println("\n*******TRIE TEST*******")
	// fmt.Println(strings.Contains("AB", "A"))
	containsFunction := func(n1, n2 uintptr) bool {
		s1 := trie.UintptrToString(n1)
		s2 := trie.UintptrToString(n2)
		result := strings.Contains(s1, s2)
		// fmt.Println("Comparing '"+s1+"' ; '"+s2+"' --> ", result)
		return result
	}
	equalsFunction := func(n1, n2 uintptr) bool {
		s1 := trie.UintptrToString(n1)
		s2 := trie.UintptrToString(n2)
		result := (s1 == s2)
		// fmt.Println("Equaling '"+s1+"' ; '"+s2+"' --> ", result)
		return result
	}
	emptyString := ""
	emptyStringPtr := unsafe.Pointer(&emptyString)
	mytrie := trie.Construct(uintptr(emptyStringPtr), containsFunction, equalsFunction)
	mytrie.Insert(trie.StringToUintptr("A"))
	mytrie.Insert(trie.StringToUintptr("B"))
	mytrie.Insert(trie.StringToUintptr("BC"))
	mytrie.Insert(trie.StringToUintptr("BCD"))
	mytrie.Insert(trie.StringToUintptr("BCDE"))
	mytrie.Insert(trie.StringToUintptr("AB"))
	mytrie.Insert(trie.StringToUintptr("ABC"))
	fmt.Println(mytrie.Depth() == 4)
	mytrie.Remove(trie.StringToUintptr("BCDE"))
	mytrie.Remove(trie.StringToUintptr("BCD"))
	mytrie.Remove(trie.StringToUintptr("BC"))
	fmt.Println(mytrie.Depth() == 3)
	fmt.Println(mytrie.Find(trie.StringToUintptr("A")) != nil)
	fmt.Println(mytrie.Find(trie.StringToUintptr("ABC")) != nil)
	fmt.Println(mytrie.Find(trie.StringToUintptr("ABCD")) == nil)

}
