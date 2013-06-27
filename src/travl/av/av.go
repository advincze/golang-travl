package av

import (
	"encoding/json"
	"time"
)

type BitAv interface {
	Set(from, to time.Time, value byte)
	Get(from, to time.Time, resolution TimeResolution) *BitVector
	SetAt(at time.Time, value byte)
	GetAt(at time.Time) byte
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
