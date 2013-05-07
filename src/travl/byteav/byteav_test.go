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
	defer func() {
		if r := recover(); r != nil {
			t.Error("set should not panic")
		}
	}()

	byteAv := New(Minute5)
	byteAv.Set(now, now.Add(15*time.Minute), 1)
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
		// t.SkipNow()
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

func TestSetWhereShiftIsNeeded(t *testing.T) {
	b := New(Minute5)
	b.Set(now, now.Add(5*time.Minute), 17)
	b.Set(now.Add(-10*time.Minute), now.Add(-5*time.Minute), 13)

	expected := []byte{0, 13, 0, 17, 0}

	if act := b.Get(now.Add(-15*time.Minute), now.Add(10*time.Minute)); !bytes.Equal(act, expected) {
		t.Errorf("av should return %v , was %v", expected, act)
	}
}

func TestSetAt(t *testing.T) {
	b := New(Minute5)
	b.SetAt(now, 22)
	if bytes := b.Get(now.Add(-5*time.Minute), now); len(bytes) != 1 || bytes[0] != 0 {
		t.Errorf("unset byte before should be zero, was %v", bytes)
	}
	if bytes := b.Get(now, now.Add(5*time.Minute)); len(bytes) != 1 || bytes[0] != 22 {
		t.Errorf("set byte should be 22, was %v", bytes)
	}
	if bytes1 := b.Get(now.Add(5*time.Minute), now.Add(10*time.Minute)); len(bytes1) != 1 || bytes1[0] != 0 {
		t.Errorf("unset byte after should be zero, was %v", bytes1)
	}
}

func TestGetAt(t *testing.T) {
	b := New(Minute5)
	b.SetAt(now, 23)

	if byt := b.GetAt(now.Add(-5 * time.Minute)); byt != 0 {
		t.Errorf("unset byte before should be zero, was %v", byt)
	}

	if byt := b.GetAt(now); byt != 23 {
		t.Errorf("set byte should be 23, was %v", byt)
	}
	if byt := b.GetAt(now.Add(5 * time.Minute)); byt != 0 {
		t.Errorf("unset byte after should be zero, was %v", byt)
	}

}
