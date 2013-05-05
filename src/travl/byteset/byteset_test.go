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

func TestSetSinglePoints(t *testing.T) {
	b := New(0)
	b.Set(10, 13)
	b.Set(18, 17)
	if b.Get(10) != 13 {
		t.Error("set byte should be 13")
	}
	if b.Get(18) != 17 {
		t.Error("set byte should be 17")
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

func TestSetFromToWithWrongFromAndTo(t *testing.T) {
	b := New(0)
	if _, err := b.SetFromTo(10, 9, 13); err == nil {
		t.Errorf("wrong ordered from and to should return an error ")
	}
}

func TestSetFromTo(t *testing.T) {
	b := New(0)
	b, err := b.SetFromTo(10, 20, 13)
	if err != nil {
		t.Error("err should be nil instead of %v", err)
	}
	if b.Get(9) != 0 {
		t.Error("set byte should be 0")
	}
	if b.Get(10) != 13 {
		t.Error("set byte should be 13")
	}
	if b.Get(19) != 13 {
		t.Error("set byte should be 13")
	}
	if b.Get(20) != 0 {
		t.Error("set byte should be 0")
	}
	if b.Get(21) != 0 {
		t.Error("set byte should be 0")
	}
}

func TestGetFromToSetPoint(t *testing.T) {
	b := New(0)
	b.Set(9, 13)
	if v, exp := b.GetFromTo(9, 10), []byte{0, 0}; bytes.Equal(v, exp) {
		t.Errorf("set bytes should be %v, was %v", exp, v)
	}
}

func TestMinEmpty(t *testing.T) {
	b := New(0)
	if act := b.Min(10, 20); act != 0 {
		t.Errorf("min of empty set should be 0, was %v", act)
	}
}

func TestMin(t *testing.T) {
	b := New(0)
	b.Set(15, 15)
	b.Set(14, 14)
	b.Set(13, 13)
	b.Set(12, 12)
	exp := byte(12)
	if act := b.Min(12, 15); act != exp {
		t.Errorf("min of the set should be %v, was %v", exp, act)
	}
}

func TestMaxEmpty(t *testing.T) {
	b := New(0)
	if act := b.Max(10, 20); act != 0 {
		t.Errorf("max of empty set should be 0, was %v", act)
	}
}

func TestMax(t *testing.T) {
	b := New(0)
	b.Set(15, 15)
	b.Set(14, 14)
	b.Set(13, 13)
	b.Set(12, 12)
	exp := byte(14)
	if act := b.Max(12, 15); act != exp {
		t.Errorf("max of the set should be %v, was %v", exp, act)
	}
}

func TestEqualEmpty(t *testing.T) {
	b := New(0)
	if act := New(0); !b.Equal(act) {
		t.Error("empty bytesets should be equal")
	}
}

func TestEqual(t *testing.T) {
	b := New(0)
	b.Set(15, 13)
	b.Set(105, 13)
	c := New(0)
	c.Set(15, 13)
	c.Set(105, 13)
	if !b.Equal(c) {
		t.Error("byetsets should be equal")
	}
}

func TestDifferentLegthNotEqual(t *testing.T) {
	b := New(0)
	if act := New(1); b.Equal(act) {
		t.Error("empty bytesets with different length should not be equal")
	}
}

func TestEmptyHistogram(t *testing.T) {
	bs := New(100)
	for i, b := range bs.Histogram() {
		switch i {
		case 0:
			if b != 100 {
				t.Errorf("histogram value should be 100, h[%v] was %v", i, b)
			}
		default:
			if b != 0 {
				t.Errorf("all histogram values should be 0, h[%v] was %v", i, b)
			}
		}
	}
}

func TestHistogram(t *testing.T) {
	bs := New(100)
	bs.Set(15, 15)
	bs.Set(16, 15)
	bs.Set(99, 15)
	bs.Set(88, 33)
	for i, b := range bs.Histogram() {
		switch i {
		case 0:
			if b != 96 {
				t.Errorf("histogram value should be 96, h[%v] was %v", i, b)
			}
		case 15:
			if b != 3 {
				t.Errorf("histogram value should be 3, h[%v] was %v", i, b)
			}
		case 33:
			if b != 1 {
				t.Errorf("histogram value should be 1, h[%v] was %v", i, b)
			}
		default:
			if b != 0 {
				t.Errorf("histogram values should be 0, h[%v] was %v", i, b)
			}
		}
	}
}

func TestLenEmpty(t *testing.T) {
	bs := New(0)
	if bs.Len() != 0 {
		t.Errorf("empty byteset should have len 0")
	}
}

func TestLen23(t *testing.T) {
	bs := New(23)
	if bs.Len() != 23 {
		t.Errorf("byteset should have len 23")
	}
}

func TestShiftBy(t *testing.T) {
	bs := New(10)
	bs.Set(0, 13)
	if bs.Get(0) != 13 {
		t.Error("should return set value")
	}
	if bs.Get(5) != 0 {
		t.Error("unset value should return zero")
	}

	bs.ShiftBy(5)

	if bs.Get(0) != 0 {
		t.Error("unset value should return zero after shift")
	}
	if bs.Get(5) != 13 {
		t.Error("should return set value after shift ")
	}

}
