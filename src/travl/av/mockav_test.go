package av

import (
	"testing"
	"time"
)

func TestNewBitAvShouldNotBeNil(t *testing.T) {
	ba := NewMockBitAv()

	if ba == nil {
		t.Errorf("BitAv should not be nil")
	}
}

func TestSetAvAtShouldNotPanic(t *testing.T) {
	ba := NewBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)

	defer func() {
		if r := recover(); r != nil {
			t.Error("SetAv should not have caused a panic")
		}
	}()

	ba.SetAvAt(t1, true)
}

func TestGetAvAtEmpty(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := NewBitAv()

	//w
	res := ba.GetAvAt(t1)

	//t
	if res {
		t.Errorf("the bit should not be set")
	}
}

func TestGetAvAt(t *testing.T) {
	// |0...0001111111111111111100000...000|
	//          |----get-----|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := NewBitAv()
	ba.SetAvAt(t1, true)

	//w
	res := ba.GetAvAt(t1)

	//t
	if !res {
		t.Errorf("the bit should be set")
	}
}

func TestSetAvFromToShouldNotPanic(t *testing.T) {
	ba := NewBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(24 * time.Hour)

	defer func() {
		if r := recover(); r != nil {
			t.Error("SetAv should not have caused a panic")
		}
	}()

	ba.SetAv(t1, t2, true)
}

func TestGetAvNothingFromEmpty(t *testing.T) {
	// |000000...000000000000|
	//       || get
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	ba := NewBitAv()

	baRes := ba.GetAv(t1, t1, Minute)

	if baRes == nil {
		t.Errorf("the baRes should not be nil")
	}
	if baRes.Bs == nil {
		t.Errorf("the baRes bitset should not be nil")
	}
	if baRes.Bs.Len() != 0 {
		t.Errorf("the baRes bitset should have length zero")
	}

}

func TestGetAvFromEmpty(t *testing.T) {
	// |000000...000000000000|
	//       |---get---|
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := t1.Add(5 * time.Minute)
	ba := NewBitAv()

	//w
	baRes := ba.GetAv(t1, t2, Minute)

	//t
	if baRes.Bs.Len() != 5 {
		t.Errorf("the baRes bitset should have length 5")
	}
	if baRes.Bs.Any() {
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
	ba := NewBitAv()
	ba.SetAv(t3, t4, true)

	//w
	baRes := ba.GetAv(t1, t2, Minute)

	//t
	if baRes.Bs.Len() != 5 {
		t.Errorf("the baRes bitset should have length 5")
	}
	if baRes.Bs.Any() {
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
	ba := NewBitAv()
	ba.SetAv(t1, t2, true)

	//w
	baRes := ba.GetAv(t3, t4, Minute)

	//t
	if baRes.Bs.Len() != 10 {
		t.Errorf("the baRes bitset should have length 10")
	}
	if baRes.Bs.Any() {
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
	ba := NewBitAv()
	ba.SetAv(t1, t4, true)

	//w
	baRes := ba.GetAv(t2, t3, Minute)

	//t
	if baRes.Bs.Len() != 30 {
		t.Errorf("the baRes bitset should have length 30")
	}
	if !baRes.Bs.All() {
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
	ba := NewBitAv()
	ba.SetAv(t2, t4, true)

	//w
	baRes := ba.GetAv(t1, t3, Minute)

	//t
	if baRes.Bs.Len() != 45 {
		t.Errorf("the baRes bitset should have length 40, was %v \n", baRes.Bs.Len())
	}
	if baRes.Bs.Count() != 30 {
		t.Errorf("30 of the bits should be set \n")
	}
}

func BenchmarkSetAvTwoYearsWorkingHours(b *testing.B) {

	ba := NewBitAv()
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(1982, 2, 8, 0, 0, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		ba.SetAv(t1, t2, true)
	}

}
