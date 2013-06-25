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

type BitVector struct {
	Resolution TimeResolution
	Start      time.Time
	Data       []byte
}

func (bitVector *BitVector) All() bool {
	for _, b := range bitVector.Data {
		if b == 0 {
			return false
		}
	}
	return true
}

func (bitVector *BitVector) Any() bool {
	for _, b := range bitVector.Data {
		if b == 1 {
			return true
		}
	}
	return false
}

func (bitVector *BitVector) Count() int {
	count := 0
	for _, b := range bitVector.Data {
		if b == 1 {
			count++
		}
	}
	return count
}

func timeToUnit(t time.Time, res TimeResolution) int {
	return int(t.Unix() / int64(res))
}

func floorDate(t time.Time, res TimeResolution) time.Time {
	if tooMuch := t.Unix() % int64(res); tooMuch != 0 {
		return t.Add(time.Duration(-1*tooMuch) * time.Second)
	}
	return t
}
