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
