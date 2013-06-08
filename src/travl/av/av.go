package av

import (
	"bytes"
	"github.com/willf/bitset"
	"time"
)

type TimeResolution int

const (
	sec      TimeResolution = 1
	Minute   TimeResolution = sec * 60
	Minute5  TimeResolution = Minute * 5
	Minute15 TimeResolution = Minute * 15
	Hour     TimeResolution = Minute * 60
	Day      TimeResolution = Hour * 24
)

func (tr TimeResolution) String() string {
	switch tr {
	case sec:
		return "sec"
	case Minute:
		return "min"
	case Minute5:
		return "5 min"
	case Minute15:
		return "15 min"
	case Hour:
		return "hour"
	case Day:
		return "day"
	}
	panic("no other options")

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

type BitAv struct {
	internalRes TimeResolution
	offset      int64
	bs          *bitset.BitSet
}

type bitAvResult struct {
	From time.Time
	To   time.Time
	Bs   *bitset.BitSet
}

func (ba *bitAvResult) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("BitAvResult{ ")
	buffer.WriteString("from:")
	buffer.WriteString(ba.From.String())
	buffer.WriteString(", to:")
	buffer.WriteString(ba.To.String())
	buffer.WriteString(", data:")
	bits := ba.Bs.DumpAsBits()
	buffer.WriteString(bits[:ba.Bs.Len()])
	buffer.WriteString(" }")
	return buffer.String()
}

func NewBitAv() *BitAv {
	return &BitAv{internalRes: Minute, bs: bitset.New(16364801)}
}

func (ba *BitAv) SetAv(from, to time.Time, value bool) {

}

func (ba *BitAv) SetAvAt(at time.Time, value bool) {
	atUnit := timeToUnit(at, ba.internalRes)
	ba.SetAvAtUnit(atUnit, value)
}

func (ba *BitAv) SetAvAtUnit(atUnit int64, value bool) {
	// log.Println("SetAvAtUnit, ", atUnit)
	ba.bs.Set(uint(atUnit))
}

func (ba *BitAv) GetAvAt(at time.Time) bool {
	atUnit := timeToUnit(at, ba.internalRes)
	return ba.getAvAtUnit(atUnit)
}

func (ba *BitAv) GetAv(from, to time.Time, res TimeResolution) *bitAvResult {
	fromUnit := timeToUnit(from, ba.internalRes)
	toUnit := timeToUnit(to, ba.internalRes)
	bs := ba.getAvUnit(fromUnit, toUnit)

	return &bitAvResult{
		From: floorDate(from, ba.internalRes),
		To:   floorDate(to, ba.internalRes),
		Bs:   bs,
	}
}

func (ba *BitAv) getAvAtUnit(atUnit int64) bool {
	// log.Println("GetAvAtUnit, ", atUnit)
	return ba.bs.Test(uint(atUnit))
}

func (ba *BitAv) getAvUnit(from, to int64) *bitset.BitSet {
	length := uint(to - from)
	data := bitset.New(length)

	return data

	// return nil
}
