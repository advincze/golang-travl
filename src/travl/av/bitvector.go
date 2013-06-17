package av

import (
	"github.com/willf/bitset"
)

type BitVectorImpl struct {
	bs *bitset.BitSet
}

func NewBitVector(length uint) *BitVectorImpl {
	return &BitVectorImpl{bs: bitset.New(length)}
}

func (bitVector *BitVectorImpl) Len() uint {
	return bitVector.bs.Len()
}

func (bitVector *BitVectorImpl) All() bool {
	return bitVector.bs.All()
}

func (bitVector *BitVectorImpl) Any() bool {
	return bitVector.bs.Any()
}

func (bitVector *BitVectorImpl) Count() uint {
	return bitVector.bs.Count()
}

func (bitVector *BitVectorImpl) Set(i uint, value bool) {
	bitVector.bs.SetTo(i, value)
}

func (bitVector *BitVectorImpl) Get(i uint) bool {
	return bitVector.bs.Test(i)
}
