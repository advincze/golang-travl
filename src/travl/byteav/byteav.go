package byteav

import (
	"time"
)

type TimeRes int

const Minute TimeRes = 1
const Minute5 TimeRes = 5 * Minute
const Minute15 TimeRes = 15 * Minute
const Hour TimeRes = 60 * Minute
const Day TimeRes = 24 * Hour

type ByteAv struct {
	Resolution TimeRes
	offset     int64
	arr        []byte
}

func New(res TimeRes) *ByteAv {
	return &ByteAv{Resolution: res}
}

func (b *ByteAv) Set(from, to time.Time, value byte) error {
	return nil
}

func (b *ByteAv) Get(from, to time.Time) []byte {
	println("res", int(b.Resolution))
	println(int(to.Sub(from).Minutes()))
	length := int(to.Sub(from).Minutes()) / int(b.Resolution)
	println("length", length)
	return make([]byte, length)
}
