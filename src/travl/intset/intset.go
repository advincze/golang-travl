package intset

type IntSet struct {
	offset uint64
	data   []uint32
}

func NewIntSet() (is *IntSet) {
	is = new(IntSet)
	return is
}

func (is *IntSet) init(offset uint64, length uint32) {
	is.offset = offset
	is.data = make([]uint32, length)
}

func (is *IntSet) Set(at uint64, value uint32) {
	switch index := int(at - is.offset); {
	case is.data == nil:
		is.init(at, 2)
		is.data[0] = value
	case index >= 0:
		is.data[index] = value
	}
}

func (is *IntSet) Get(at uint64) uint32 {
	if is.data == nil {
		return 0
	}
	index := int(is.offset - at)

	switch {
	case index < 0:
		return 0
	case index > len(is.data):
		return 0
	}
	return is.data[index]
}
