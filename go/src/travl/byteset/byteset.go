package byteset

import (
	"errors"
	"log"
)

type ByteSet struct {
	length uint
	set    []byte
}

func New(length uint) *ByteSet {
	return &ByteSet{0, make([]byte, length)}
}

func (bs *ByteSet) Set(i uint, value byte) *ByteSet {
	bs.extendIfNeeded(i)
	bs.set[i] = value
	return bs
}

func (bs *ByteSet) SetFromTo(from uint, to uint, value byte) *ByteSet {
	if checkFromTo(from, to) != nil {
		panic("oh no an error")
	}
	bs.extendIfNeeded(to)
	for i := from; i < to; i++ {
		bs.set[i] = value
	}
	return bs
}

func (bs *ByteSet) Get(i uint) byte {
	bs.extendIfNeeded(i)
	return bs.set[i]
}

func (bs *ByteSet) GetFromTo(from uint, to uint) []byte {
	bs.extendIfNeeded(to)
	return bs.set[from:to]
}

func checkFromTo(from, to uint) error {
	if from > to {
		return errors.New("from must be <= than to")
	}
	return nil
}

func (bs *ByteSet) extendIfNeeded(i uint) {

	if bs.length < i {
		log.Println(bs, i)
		newlength := i + 1
		bs.set = make([]byte, newlength)
		bs.length = newlength
	}
}
