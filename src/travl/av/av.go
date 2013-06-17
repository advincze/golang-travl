package av

import (
	"time"
)

type BitAv interface {
	Set(from, to time.Time, value bool)
	SetAt(at time.Time, value bool)
	GetAt(at time.Time) bool
	Get(from, to time.Time, res TimeResolution) *BitVector
}

type BitVector interface {
	Len() uint
	Set(i uint, value bool)
	Get(i uint) bool
}

func timeToUnit(t time.Time, res TimeResolution) int64 {
	return t.Unix() / int64(res)
}

func floorDate(t time.Time, res TimeResolution) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	return t
}
