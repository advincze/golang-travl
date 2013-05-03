package byteav

import (
	"bytes"
	"testing"
	"time"
)

var now time.Time = time.Date(1998, time.February, 1, 0, 0, 0, 0, time.UTC)

func TestNewByteAv(t *testing.T) {
	byteAv := New(Minute5)
	if byteAv == nil {
		t.Errorf("Empty byteav should not be nil")
	}
}

func TestSetFromTo(t *testing.T) {
	byteAv := New(Minute5)

	err := byteAv.Set(now, now.Add(15*time.Minute), 1)
	if err != nil {
		t.Errorf("Set should not return an error")
	}
}

func TestGetFromEmptyAv(t *testing.T) {
	byteAv := New(Minute5)
	byteArr := byteAv.Get(now, now.Add(15*time.Minute))
	expected := []byte{0, 0, 0}
	if !bytes.Equal(byteArr, expected) {
		t.Errorf("empty av should return %v , was %v", expected, byteArr)
	}
}

func TestGetFromSingleSetUnit(t *testing.T) {
	byteAv := New(Minute5)
	byteAv.Set(now, now.Add(5*time.Minute), 1)
	byteArr := byteAv.Get(now, now.Add(15*time.Minute))
	expected := []byte{1, 0, 0}
	if !bytes.Equal(byteArr, expected) {
		t.SkipNow()
		t.Errorf(" av should return %v , was %v", expected, byteArr)
	}
}

func TestRoundDateWithNoRoundExpected(t *testing.T) {
	t0 := time.Date(1998, time.February, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(59 * time.Minute)
	expected := t1
	if act := roundDate(t1, Minute); act != t1 {
		t.Errorf(" %v rounded with %v should be %v, was %v", t1, Minute, expected, act)
	}
}

func TestRoundDateWithResMin15(t *testing.T) {
	expected := time.Date(1998, time.February, 1, 0, 0, 0, 0, time.UTC)
	t1 := expected.Add(11 * time.Minute)
	if act := roundDate(t1, Minute15); act != expected {
		t.Errorf(" %v rounded with %v should be %v, was %v", t1, Minute15, expected, act)
	}
}

func TestRoundDateWithResHour(t *testing.T) {
	expected := time.Date(1998, time.February, 1, 0, 0, 0, 0, time.UTC)
	t1 := expected.Add(59 * time.Minute)
	if act := roundDate(t1, Hour); act != expected {
		t.Errorf(" %v rounded with %v should be %v, was %v", t1, Minute15, expected, act)
	}
}
