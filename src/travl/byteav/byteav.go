package byteav

import (
	"time"
	"travl/byteset"
)

type TimeRes int

const second TimeRes = 1
const Minute TimeRes = 60 * second
const Minute5 TimeRes = 5 * Minute
const Minute15 TimeRes = 15 * Minute
const Hour TimeRes = 60 * Minute
const Day TimeRes = 24 * Hour

func (tr TimeRes) String() string {
	switch tr {
	case second:
		return "sec"
	case Minute:
		return "min"
	case Minute5:
		return "5 min"
	case Minute15:
		return "15 min"
	case Hour:
		return "hour"
	case Day:
		return "day"
	}
	panic("no other options")
}

type ByteAv struct {
	internalRes TimeRes
	offset      int64
	byteset     *byteset.ByteSet
}

func New(res TimeRes) *ByteAv {
	return &ByteAv{internalRes: res}
}

func (b *ByteAv) SetAt(at time.Time, value byte) *ByteAv {
	atUnit := timeToUnit(at, b.internalRes)
	b.setUnit(atUnit, atUnit+1, value)
	return b
}

func (b *ByteAv) Set(from, to time.Time, value byte) *ByteAv {
	fromUnit := timeToUnit(from, b.internalRes)
	toUnit := timeToUnit(to, b.internalRes)
	b.setUnit(fromUnit, toUnit, value)
	return b
}

func (b *ByteAv) Get(from, to time.Time) []byte {
	fromUnit := timeToUnit(from, b.internalRes)
	toUnit := timeToUnit(to, b.internalRes)
	return b.getUnit(fromUnit, toUnit)
}

func (b *ByteAv) GetAt(at time.Time) byte {
	atUnit := timeToUnit(at, b.internalRes)
	return b.getUnit(atUnit, atUnit+1)[0]
}

func (b *ByteAv) getUnit(from, to int64) []byte {
	length := to - from
	data := make([]byte, length)
	if b.byteset == nil {
		return data
	}
	k := int64(0)
	i := uint(0)
	if from < b.offset {
		k = b.offset - from
	} else {
		i = uint(from - b.offset)
	}
	for ; k < length; i, k = i+1, k+1 {
		data[k] = b.byteset.Get(i)
	}
	return data
}

func (b *ByteAv) setUnit(from, to int64, value byte) {
	if b.byteset == nil {
		b.offset = from
		b.byteset = byteset.New(0)

	} else if from < b.offset {
		b.shiftOffset(from)
	}
	b.setInternal(uint(from-b.offset), uint(to-b.offset), value)
}

func (b *ByteAv) setInternal(from, to uint, value byte) {
	println("setInternal", from, to, value)
	b.byteset.SetFromTo(from, to, value)
}

func timeToUnit(t time.Time, res TimeRes) int64 {
	return t.Unix() / int64(res)
}

func roundDate(t time.Time, res TimeRes) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	return t
}

func (b *ByteAv) shiftOffset(maxOffset int64) {
	prepend := uint(b.offset - maxOffset)
	println("shift off", b.offset, maxOffset, prepend)
	b.byteset.ShiftBy(prepend)
	b.offset = maxOffset
}
