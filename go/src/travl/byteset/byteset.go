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

func (bs *ByteSet) SetFromTo(from, to uint, value byte) (*ByteSet, error) {
	if err := checkFromTo(from, to); err != nil {
		return bs, err
	}
	bs.extendIfNeeded(to)
	for i := from; i <= to; i++ {
		bs.set[i] = value
	}
	return bs, nil
}

func (bs *ByteSet) Get(i uint) byte {
	log.Printf("Get(%v)", i)
	bs.extendIfNeeded(i)
	return bs.set[i]
}

func (bs *ByteSet) GetFromTo(from, to uint) []byte {
	if err := checkFromTo(from, to); err != nil {
		return nil
	}
	bs.extendIfNeeded(to)
	return bs.set[from:to]
}

func checkFromTo(from, to uint) error {
	log.Println("check", from, to)
	if from > to {
		return errors.New("from must be <= than to")
	}
	return nil
}

func (bs *ByteSet) extendIfNeeded(i uint) {
	if bs.length <= i {
		log.Println("extend", bs, i)
		newlength := i + 1
		oldset := bs.set //make([]byte, newlength)
		bs.set = make([]byte, newlength)
		copy(bs.set, oldset)
		bs.length = newlength
	}
}
