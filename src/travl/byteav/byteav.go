package byteav

import (
	"travl/byteset"
)

type ByteAv struct {
	bs     *byteset.ByteSet
	offset uint
}

func New(length uint) *ByteAv {
	return &ByteAv{}
}

func (ba *ByteAv) Set(i uint, value byte) *ByteAv {
	if ba.bs == nil {
		ba.bs = byteset.New(0)
	}
	ba.bs.Set(i, value)
	return ba
}

func (ba *ByteAv) SetFromTo(from, to uint, value byte) (*ByteAv, error) {
	if ba.bs == nil {
		ba.bs = byteset.New(0)
	}
	_, err := ba.bs.SetFromTo(from, to, value)
	return ba, err
}

func (ba *ByteAv) Get(i uint) byte {
	if ba.bs == nil {
		return 0
	}
	return ba.bs.Get(i)
}

func (ba *ByteAv) GetFromTo(from, to uint) []byte {
	if ba.bs == nil {
		ba.bs = byteset.New(0)
	}
	return ba.bs.GetFromTo(from, to)
}

func (ba *ByteAv) Max(from, to uint) byte {
	if ba.bs == nil {
		return 0
	}
	return ba.bs.Max(from, to)
}

func (ba *ByteAv) Min(from, to uint) byte {
	if ba.bs == nil {
		return 0
	}
	return ba.bs.Min(from, to)
}

func (ba *ByteAv) Equal(c *ByteAv) bool {
	if c == nil {
		return false
	}
	if ba.offset != c.offset {
		return false
	}
	if ba.bs == nil {
		return c.bs == nil
	}
	return ba.bs.Equal(c.bs)
}

func (ba *ByteAv) Histogram() (h [256]uint) {
	return ba.bs.Histogram()
}
