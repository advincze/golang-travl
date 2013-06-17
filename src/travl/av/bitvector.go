package av

import (
	"github.com/willf/bitset"
)

type BitVectorImpl struct {
	data bitset.BitSet
}

func NewBitVectorImpl() *BitVectorImpl {
	return &BitVectorImpl{data: *bitset.New()}
}

func (bitVector *BitVectorImpl) Len() uint {
	return bitVector.data.Len()
}

func (bitVector *BitVectorImpl) Set(i uint, value bool) {
	bitVector.data.SetTo(i, value)
}

func (bitVector *BitVectorImpl) Get(i uint) bool {
	return bitVector.data.Test(i)
}
