package av

import (
	"testing"
	"time"
)

func newBitAv() BitAv {
	return NewMockBitAv()
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

	ba.SetAt(t1, true)
}

func TestGetAvAtEmpty(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()

	//w
	res := ba.GetAt(t1)

	//t
	if res {
		t.Errorf("the bit should not be set")
	}
}

func TestGetAvAt(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()
	ba.SetAt(t1, true)

	//w
	res := ba.GetAt(t1)

	//t
	if !res {
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

	ba.Set(t1, t2, true)
}

func TestGetAvNothingFromEmpty(t *testing.T) {
	// |000000...000000000000|
	//       || get
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := newBitAv()

	BitVector := ba.Get(t1, t1, Minute)

	if BitVector == nil {
		t.Errorf("the BitVector should not be nil")
	}
	if BitVector.Len() != 0 {
		t.Errorf("the BitVector should have length zero")
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
	if bitVector.Len() != 5 {
		t.Errorf("the baRes bitset should have length 5")
	}
	//TODO
	// if bitVector.Bs.Any() {
	// 	t.Errorf("none of the bits should be set")
	// }
}

func TestGetAvFromBeforeSet(t *testing.T) {
	// |000000...000000000011001....01100000|
	//     |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(5 * time.Minute)
	t3 := t1.Add(25 * time.Minute)
	t4 := t1.Add(45 * time.Minute)
	ba := newBitAv()
	ba.Set(t3, t4, true)

	//w
	bitVector := ba.Get(t1, t2, Minute)

	//t
	if bitVector.Len() != 5 {
		t.Errorf("the bitVector bitset should have length 5")
	}
	//TODO
	// if bitVector.Any() {
	// 	t.Errorf("none of the bits should be set")
	// }
}

func TestGetAvFromAfterSet(t *testing.T) {
	// |0...000111000110111111100000....00000|
	//                          |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t1, t2, true)

	//w
	bitVector := ba.Get(t3, t4, Minute)

	//t
	if bitVector.Len() != 10 {
		t.Errorf("the bitVector bitset should have length 10")
	}
	//TODO
	// if baRes.Bs.Any() {
	// 	t.Errorf("none of the bits should be set")
	// }
}

func TestGetAvFromInsideSet(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t1, t4, true)

	//w
	bitVector := ba.Get(t2, t3, Minute)

	//t
	if bitVector.Len() != 30 {
		t.Errorf("the baRes bitset should have length 30")
	}
	//TODO
	// if !baRes.Bs.All() {
	// 	t.Errorf("all of the bits should be set")
	// }

}

func TestGetAvFromItersectSet(t *testing.T) {
	// |00..000000001111111111100000...00|
	//         |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(15 * time.Minute)
	t3 := t1.Add(45 * time.Minute)
	t4 := t1.Add(55 * time.Minute)
	ba := newBitAv()
	ba.Set(t2, t4, true)

	//w
	bitVector := ba.Get(t1, t3, Minute)

	//t
	if bitVector.Len() != 45 {
		t.Errorf("the baRes bitset should have length 40, was %v \n", bitVector.Len())
	}
	if bitVector.Count() != 30 {
		t.Errorf("30 of the bits should be set \n")
	}
}

func BenchmarkSetAvTwoYearsWorkingHours(b *testing.B) {

	ba := newBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(1982, 2, 8, 0, 0, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		ba.Set(t1, t2, true)
	}

}
