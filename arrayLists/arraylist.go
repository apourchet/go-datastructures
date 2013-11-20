package arraylist

type ArrayList struct {
	array  []int
	load   int
	size   int
	factor float64
}

func Construct(factor float64) *ArrayList {
	if factor == 0 {
		factor = 0.75
	}
	al := ArrayList{}
	al.array = make([]int, 4)
	al.load = 0
	al.size = 4
	al.factor = factor
	return &al
}

func (al *ArrayList) resize() {
	newArray := make([]int, al.size*2)
	copy(newArray, al.array)
	al.array = newArray
	al.size *= 2
}

func (al *ArrayList) Get(index int) int {
	if index >= al.size {
		return 0
	}
	return al.array[index]
}

func (al *ArrayList) Insert(objectPointer int) bool {
	if float64(al.load+1) >= float64(al.size)*al.factor {
		al.resize()
	}
	al.array[al.load] = objectPointer
	al.load++
	return true
}

func (al *ArrayList) Remove(index int) bool {
	if index >= al.size {
		return false
	}
	al.array[index] = 0
	return true
}
