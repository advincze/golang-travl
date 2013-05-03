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

func (b *ByteAv) Set(from, to time.Time, value byte) error {
	fromFrame := timeToFrame(from, b.internalRes)
	toFrame := timeToFrame(to, b.internalRes)
	return b.setFrame(fromFrame, toFrame, value)
}

func (b *ByteAv) Get(from, to time.Time) []byte {
	length := int(to.Sub(from).Seconds()) / int(b.internalRes)
	return make([]byte, length)
}

func (b *ByteAv) setFrame(from, to int64, value byte) error {
	if b.byteset == nil {
		b.offset = from
		b.byteset = byteset.New(0)

	} else if from < b.offset {
		b.shiftOffset(from)
	}
	return b.setInternal(from-b.offset, to-b.offset, value)
}

func (b *ByteAv) setInternal(from, to int64, value byte) error {
	return nil
}

func timeToFrame(t time.Time, res TimeRes) int64 {
	return t.Unix() / int64(res)
}

func roundDate(t time.Time, res TimeRes) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	return t
}

func (b *ByteAv) shiftOffset(maxOffset int64) {
	//prepend := b.offset - maxOffset

}
