package bloomfilter

type BitSet struct {
	set []int64
}

func construct(bits int) *BitSet {
	rem := bits % 64
	total := 0
	if rem == 0 {
		total = bits
	} else {
		total = bits + (64 - rem)
	}
	arr := make([]int64, (total / 64))
	return &BitSet{arr}
}

// Returns whether or not the 0-indexed ith bit is 1
func (bs *BitSet) bit(i uint) bool {
	n := bs.set[i/64]
	return ((n & (1 << (i % 64))) >> (i % 64)) == 1
}

func (bs *BitSet) setBit(i uint) bool {
	bs.set[i/64] = bs.set[i/64] | (1 << (i % 64))
	return true
}

func (bs *BitSet) unsetBit(i uint) bool {
	n := int64(0)
	for in := uint(0); in < 64; in++ {
		if in != (i % 64) {
			n = n | (1 << in)
		}
	}
	bs.set[i/64] = n
	return true
}

type BloomFilter struct {
	bitset       *BitSet
	hashFunction func(i, r int) uint
	rounds       int
}

func Construct(bits, rounds int, hashFunction func(i, r int) uint) *BloomFilter {
	b := BloomFilter{}
	b.bitset = construct(bits)
	b.hashFunction = hashFunction
	b.rounds = rounds
	return &b
}

func (bf *BloomFilter) Put(i int) bool {
	for r := 0; r < bf.rounds; r++ {
		bf.bitset.setBit(bf.hashFunction(i, r))
	}
	return true
}

func (bf *BloomFilter) MightContain(i int) bool {
	for r := 0; r < bf.rounds; r++ {
		index := bf.hashFunction(i, r)
		if bf.bitset.bit(index) == false {
			return false
		}
	}
	return true
}
