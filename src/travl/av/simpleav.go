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

func (ba *SimpleBitAv) Set(from, to time.Time, value byte) {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	boolValue := value == 1
	ba.setAvUnit(fromUnit, toUnit, boolValue)
}

func (ba *SimpleBitAv) setAvUnit(from, to int, value bool) {
	// println("setfromto", from, to, value)
	for i := from; i <= to; i++ {
		ba.bs.SetTo(uint(i), value)
	}
}

func (ba *SimpleBitAv) SetAt(at time.Time, value byte) {
	atUnit := timeToUnit(at, ba.internalRes)
	ba.setAvAtUnit(atUnit, value == 1)
}

func (ba *SimpleBitAv) setAvAtUnit(atUnit int, value bool) {
	// log.Println("SetAvAtUnit, ", atUnit)
	ba.bs.SetTo(uint(atUnit), value)
}

func (ba *SimpleBitAv) GetAt(at time.Time) byte {
	atUnit := timeToUnit(at, ba.internalRes)
	boolValue := ba.getAvAtUnit(atUnit)
	if boolValue {
		return 1
	}
	return 0
}

func (ba *SimpleBitAv) getAvAtUnit(atUnit int) bool {
	// log.Println("GetAvAtUnit, ", atUnit)
	return ba.bs.Test(uint(atUnit))
}

// 	Get(from, to time.Time, res TimeResolution) BitVector

func (ba *SimpleBitAv) Get(from, to time.Time, res TimeResolution) *BitVector {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	length := toUnit - fromUnit
	bs := ba.getAvUnit(fromUnit, toUnit)
	bv := &BitVector{Resolution: res, Data: make([]byte, length), Start: floorDate(from, res)}
	for i := 0; i < length; i++ {
		if bs.Test(uint(i)) {
			bv.Data[i] = 1
		} else {
			bv.Data[i] = 0
		}
	}

	return bv
}

func (ba *SimpleBitAv) getAvUnit(from, to int) *bitset.BitSet {
	// println("getfromto", from, to)
	length := uint(to - from)
	// data := NewBitVector(length)
	data := bitset.New(length)
	var value bool
	for i, k := from, uint(0); i < to; i, k = i+1, k+1 {
		value = ba.bs.Test(uint(i))
		data.SetTo(k, value)
	}
	return data
}
