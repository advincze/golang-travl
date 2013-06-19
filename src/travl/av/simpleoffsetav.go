package av

import (
	"github.com/willf/bitset"
	"time"
	"math"
)

type SimpleOffsetBitAv struct {
	internalRes TimeResolution
	offset      int64
	bs          *bitset.BitSet
}

func NewSimpleOffsetBitAv() *SimpleOffsetBitAv {
	return &SimpleOffsetBitAv{
		internalRes: Minute, 
		bs: bitset.New(0),
		offset: math.MaxInt64,
	}
}

func (ba *SimpleOffsetBitAv) Set(from, to time.Time, value bool) {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	ba.setAvUnit(fromUnit, toUnit, value)
}

func (ba *SimpleOffsetBitAv) shiftOffset(by int64)  

func (ba *SimpleOffsetBitAv) setAvUnit(from, to int64, value bool) {
	if(ba.offset> from){
		ba.offset = from;

	}

	for i := from; i <= to; i++ {
		ba.bs.SetTo(uint(i), value)
	}
}

func (ba *SimpleOffsetBitAv) SetAt(at time.Time, value bool) {
	atUnit := timeToUnit(at, ba.internalRes)
	ba.setAvAtUnit(atUnit, value)
}

func (ba *SimpleOffsetBitAv) setAvAtUnit(atUnit int64, value bool) {
	ba.bs.SetTo(uint(atUnit), value)
}

func (ba *SimpleOffsetBitAv) GetAt(at time.Time) bool {
	atUnit := timeToUnit(at, ba.internalRes)
	return ba.getAvAtUnit(atUnit)
}

func (ba *SimpleOffsetBitAv) getAvAtUnit(atUnit int64) bool {
	return ba.bs.Test(uint(atUnit))
}


func (ba *SimpleOffsetBitAv) Get(from, to time.Time, res TimeResolution) *bitset.BitSet {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	return ba.getAvUnit(fromUnit, toUnit)
}

func (ba *SimpleOffsetBitAv) getAvUnit(from, to int64) *bitset.BitSet {
	length := uint(to - from)
	data := bitset.New(length)
	var value bool
	for i, k := from, uint(0); i < to; i, k = i+1, k+1 {
		value = ba.bs.Test(uint(i))
		data.SetTo(k, value)
	}
	return data
}
