package av

import (
	"time"
)

type BitAv interface {
	Set(from, to time.Time, value byte)
	Get(from, to time.Time, resolution TimeResolution) *BitVector
	SetAt(at time.Time, value byte)
	GetAt(at time.Time) byte
}

func timeToUnitFloor(t time.Time, res TimeResolution) int {
	return int(t.Unix() / int64(res))
}

func floorDate(t time.Time, res TimeResolution) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	return t
}

func ceilDate(t time.Time, res TimeResolution) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second).Add(time.Duration(int(res)) * time.Second)
	}
	return t
}

var bitavs map[string]BitAv = make(map[string]BitAv)

func NewBitAv(id string) BitAv {
	bitav := NewBitSegmentAv(id, Minute5)
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
