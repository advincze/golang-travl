package bitav

import (
	"time"
	"travl/av"
)

type BitAv interface {
	Set(from, to time.Time, value byte)
	Get(from, to time.Time, resolution av.TimeResolution) *BitVector
	SetAt(at time.Time, value byte)
	GetAt(at time.Time) byte
}

var bitavs map[string]BitAv = make(map[string]BitAv)

func NewBitAv(id string) BitAv {
	bitav := NewBitSegmentAv(id, av.Minute5, "mem")
	bitavs[id] = bitav
	return bitav
}

func FindBitAv(id string) BitAv {
	return bitavs[id]
}

func FindOrNewBitAv(id string) BitAv {
	bitav := FindBitAv(id)
	if bitav == nil {
		bitav = NewBitAv(id)
	}
	return bitav
}

func DeleteBitAv(id string) {
	delete(bitavs, id)
}
