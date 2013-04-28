package byteset

import (
	"bytes"
	"testing"
)

func TestEmptyByteSet(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("A zero-length byteset should be fine")
		}
	}()
	b := New(0)
	if b == nil {
		t.Error("byteset should not be nil")
	}
}

func TestGetSingleUnsetPoint(t *testing.T) {
	b := New(0)
	if b.Get(10) != 0 {
		t.Error("unset byte should not be 0")
	}
}

func TestSetSinglePoint(t *testing.T) {
	b := New(0)
	b.Set(10, 13)
	if b.Get(10) != 13 {
		t.Error("set byte should be 13")
	}
}

func TestGetFromToUnsetPoint(t *testing.T) {
	b := New(0)
	if v, exp := b.GetFromTo(9, 10), []byte{0, 0}; bytes.Equal(v, exp) {
		t.Errorf("unset bytes should be %v, was %v", exp, v)
	}
}

func TestGetFromToUnsetPointWithWrongFromAndTo(t *testing.T) {
	b := New(0)
	if act := b.GetFromTo(10, 9); act != nil {
		t.Errorf("wrong ordered from and to should return nil instead of %v", act)
	}
}

func TestSetFromToUnsetPointWithWrongFromAndTo(t *testing.T) {
	b := New(0)
	if _, err := b.SetFromTo(10, 9, 13); err == nil {
		t.Errorf("wrong ordered from and to should return an error ")
	}
}

func TestGetFromToSetPoint(t *testing.T) {
	b := New(0)
	b.Set(9, 13)
	if v, exp := b.GetFromTo(9, 10), []byte{0, 0}; bytes.Equal(v, exp) {
		t.Errorf("set bytes should be %v, was %v", exp, v)
	}
}
