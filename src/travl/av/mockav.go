package av

import (
	"bytes"
	"github.com/willf/bitset"
	"time"
)

type MockBitAv struct {
	internalRes TimeResolution
	offset      int64
	bs          *bitset.BitSet
}

type bitAvResult struct {
	From time.Time
	To   time.Time
	Bs   *bitset.BitSet
}

func (ba *bitAvResult) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("BitAvResult{ ")
	buffer.WriteString("from:")
	buffer.WriteString(ba.From.String())
	buffer.WriteString(", to:")
	buffer.WriteString(ba.To.String())
	buffer.WriteString(", data:")
	bits := ba.Bs.DumpAsBits()
	buffer.WriteString(bits[:ba.Bs.Len()])
	buffer.WriteString(" }")
	return buffer.String()
}

func NewMockBitAv() *MockBitAv {
	return &MockBitAv{internalRes: Minute, bs: bitset.New(8000000)}
}

func (ba *MockBitAv) Set(from, to time.Time, value bool) {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	ba.setAvUnit(fromUnit, toUnit, value)
}

func (ba *MockBitAv) setAvUnit(from, to int64, value bool) {
	// println("setfromto", from, to, value)
	for i := from; i <= to; i++ {
		ba.bs.SetTo(uint(i), value)
	}
}

func (ba *MockBitAv) SetAt(at time.Time, value bool) {
	atUnit := timeToUnit(at, ba.internalRes)
	ba.setAvAtUnit(atUnit, value)
}

func (ba *MockBitAv) setAvAtUnit(atUnit int64, value bool) {
	// log.Println("SetAvAtUnit, ", atUnit)
	ba.bs.SetTo(uint(atUnit), value)
}

func (ba *MockBitAv) GetAt(at time.Time) bool {
	atUnit := timeToUnit(at, ba.internalRes)
	return ba.getAvAtUnit(atUnit)
}

func (ba *MockBitAv) getAvAtUnit(atUnit int64) bool {
	// log.Println("GetAvAtUnit, ", atUnit)
	return ba.bs.Test(uint(atUnit))
}

// 	Get(from, to time.Time, res TimeResolution) BitVector

func (ba *MockBitAv) Get(from, to time.Time, res TimeResolution) BitVector {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	return BitVectorImpl{data: ba.getAvUnit(fromUnit, toUnit)}
}

func (ba *MockBitAv) getAvUnit(from, to int64) bitset.BitSet {
	// println("getfromto", from, to)
	length := uint(to - from)
	data := bitset.New(length)
	var value bool
	for i, k := from, uint(0); i < to; i, k = i+1, k+1 {
		value = ba.bs.Test(uint(i))
		data.SetTo(k, value)
		// println(value, i, k)
	}
	return *data
}
