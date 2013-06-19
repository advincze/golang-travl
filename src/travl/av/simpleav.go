package av

import (
	"github.com/willf/bitset"
	"time"
)

type SimpleBitAv struct {
	internalRes TimeResolution
	bs          *bitset.BitSet
}

func NewSimpleBitAv() *SimpleBitAv {
	return &SimpleBitAv{internalRes: Minute, bs: bitset.New(8000000)}
}

func (ba *SimpleBitAv) Set(from, to time.Time, value bool) {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	ba.setAvUnit(fromUnit, toUnit, value)
}

func (ba *SimpleBitAv) setAvUnit(from, to int64, value bool) {
	// println("setfromto", from, to, value)
	for i := from; i <= to; i++ {
		ba.bs.SetTo(uint(i), value)
	}
}

func (ba *SimpleBitAv) SetAt(at time.Time, value bool) {
	atUnit := timeToUnit(at, ba.internalRes)
	ba.setAvAtUnit(atUnit, value)
}

func (ba *SimpleBitAv) setAvAtUnit(atUnit int64, value bool) {
	// log.Println("SetAvAtUnit, ", atUnit)
	ba.bs.SetTo(uint(atUnit), value)
}

func (ba *SimpleBitAv) GetAt(at time.Time) bool {
	atUnit := timeToUnit(at, ba.internalRes)
	return ba.getAvAtUnit(atUnit)
}

func (ba *SimpleBitAv) getAvAtUnit(atUnit int64) bool {
	// log.Println("GetAvAtUnit, ", atUnit)
	return ba.bs.Test(uint(atUnit))
}

// 	Get(from, to time.Time, res TimeResolution) BitVector

func (ba *SimpleBitAv) Get(from, to time.Time, res TimeResolution) *bitset.BitSet {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	return ba.getAvUnit(fromUnit, toUnit)
}

func (ba *SimpleBitAv) getAvUnit(from, to int64) *bitset.BitSet {
	// println("getfromto", from, to)
	length := uint(to - from)
	// data := NewBitVector(length)
	data := bitset.New(length)
	var value bool
	for i, k := from, uint(0); i < to; i, k = i+1, k+1 {
		value = ba.bs.Test(uint(i))
		data.SetTo(k, value)
		// println(value, i, k)
	}
	return data
}
