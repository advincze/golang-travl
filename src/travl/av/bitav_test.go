package av

import (
	"testing"
	"time"
)

func newBitAv() *BitAv3 {
	return NewBitAv3() //NewSimpleBitAv()
}

func TestNewBitAvShouldNotBeNil(t *testing.T) {
	ba := newBitAv()

	if ba == nil {
		t.Errorf("BitAv should not be nil")
	}
}

func TestSetAvAtShouldNotPanic(t *testing.T) {
	ba := newBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)

	defer func() {
		if r := recover(); r != nil {
			t.Error("SetAv should not have caused a panic")
		}
	}()

	ba.SetAt(t1, 1)
}

func TestGetAvAtEmpty(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()

	//t
	if ba.GetAt(t1) == 1 {
		t.Errorf("the bit should not be set")
	}
}

func TestGetAvAt(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()
	ba.SetAt(t1, 1)

	//w
	res := ba.GetAt(t1)

	//t
	if res == 0 {
		t.Errorf("the bit should be set")
	}
}

func TestSetAvFromToShouldNotPanic(t *testing.T) {
	ba := newBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(24 * time.Hour)

	defer func() {
		if r := recover(); r != nil {
			t.Error("SetAv should not have caused a panic")
		}
	}()

	ba.Set(t1, t2, 1)
}

func TestGetAvNothingFromEmpty(t *testing.T) {
	// |000000...000000000000|
	//       || get
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()

	bitVector := ba.Get(t1, t1, Minute)

	if bitVector == nil {
		t.Errorf("the bitVector should not be nil")
	}
	if len(bitVector.Data) != 0 {
		t.Errorf("the bitVector should have length zero")
	}

}

func TestGetAvFromEmpty(t *testing.T) {
	// |000000...000000000000|
	//       |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(5 * time.Minute)
	ba := newBitAv()

	//w
	bitVector := ba.Get(t1, t2, Minute)

	//t
	if len(bitVector.Data) != 5 {
		t.Errorf("the bitVector should have length 5")
	}
	if bitVector.Any() {
		t.Errorf("none of the bits should be set")
	}
}

func TestGetAvFromBeforeSet(t *testing.T) {
	// |000000...000000000011001....01100000|
	//     |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(5 * time.Minute)
	t3 := t1.Add(25 * time.Minute)
	t4 := t1.Add(45 * time.Minute)
	ba := newBitAv()
	ba.Set(t3, t4, 1)

	//w
	bitVector := ba.Get(t1, t2, Minute)

	//t
	if len(bitVector.Data) != 5 {
		t.Errorf("the bitVector bitset should have length 5")
	}
	if bitVector.Any() {
		t.Errorf("none of the bits should be set")
	}
}

func TestGetAvFromAfterSet(t *testing.T) {
	// |0...000111000110111111100000....00000|
	//                          |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t1, t2, 1)

	//w
	bitVector := ba.Get(t3, t4, Minute)

	//t
	if len(bitVector.Data) != 10 {
		t.Errorf("the bitVector bitset should have length 10")
	}
	if bitVector.Any() {
		t.Errorf("none of the bits should be set")
	}
}

func TestGetAvFromInsideSet(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t1, t4, 1)

	//w
	bitVector := ba.Get(t2, t3, Minute)

	//t
	if len(bitVector.Data) != 30 {
		t.Errorf("the bitVector should have length 30")
	}
	if !bitVector.All() {
		t.Errorf("all of the bits should be set")
	}

}

func TestGetAvFromItersectSet(t *testing.T) {
	// |00..000000001111111111100000...00|
	//         |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t2, t4, 1)

	//w
	bitVector := ba.Get(t1, t3, Minute)

	//t
	if len(bitVector.Data) != 45 {
		t.Errorf("the bitVector should have length 40, was %v \n", len(bitVector.Data))
	}
	if bitVector.Count() != 30 {
		t.Errorf("30 of the bits should be set \n")
	}
}

func TestSetAvTwoYearsWorkingHours(t *testing.T) {

	ba := newBitAv()
	t1 := time.Date(1982, 2, 7, 9, 0, 0, 0, time.UTC)

	for i := 0; i < 20; i++ {
		ba.Set(t1, t1.Add(8*time.Hour), 1)
		ba.Set(t1.Add(8*time.Hour), t1.Add(12*time.Hour), 0)
		t1 = t1.Add(24 * time.Hour)
	}
	// println(len(ba.segments))
	println("size:", ba.size(), " bytes")
}

func BenchmarkSetAvTwoYearsWorkingHours(b *testing.B) {

	ba := newBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(1982, 2, 8, 0, 0, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		ba.Set(t1, t2, 1)
	}

}
