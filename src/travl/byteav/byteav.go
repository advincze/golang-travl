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

type ByteAv struct {
	internalRes TimeRes
	offset      int64
	byteset     *byteset.ByteSet
}

func New(res TimeRes) *ByteAv {
	return &ByteAv{internalRes: res}
}

func (b *ByteAv) Set(from, to time.Time, value byte) error {
	// fromInt := roundDate(from, b.internalRes)

	return nil
}

func (b *ByteAv) Get(from, to time.Time) []byte {
	length := int(to.Sub(from).Seconds()) / int(b.internalRes)
	return make([]byte, length)
}

func (b *ByteAv) setInternal(from, to int, value byte) error {
	return nil
}

func roundDate(t time.Time, res TimeRes) time.Time {
	println("rd", t.Unix(), int64(res), t.Unix()%int64(res))
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		println("too much", tooMuch)
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	println("no round")
	return t
}
