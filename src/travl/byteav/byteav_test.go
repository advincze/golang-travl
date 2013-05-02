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
