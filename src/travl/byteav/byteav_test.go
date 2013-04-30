package byteav

import "testing"

func TestShouldReturnAnEmptyNewByteAv(t *testing.T) {
	b := New()
	if b == nil {
		t.Errorf("Empty set should not be nil")
	}
}

func TestEmptyNewByteAvShouldHaveLenZero(t *testing.T) {
	b := New()
	if b.Len() != 0 {
		t.Errorf("Empty set should not have length 0")
	}
}
