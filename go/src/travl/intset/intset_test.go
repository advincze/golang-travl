package intset

import "testing"

func TestShouldCreateAnEmptySet(t *testing.T) {
	if is := NewIntSet(); is == nil {
		t.Error("the created intset was nil")
	}
}

func TestShouldSetASinglePoint(t *testing.T) {
	var is *IntSet = NewIntSet()
	is.Set(100, 1)
	assertValue(t, is.Get(100), 1)
}

func TestSetSinglePointAndRetrieveOneOutOfIndex(t *testing.T) {
	var is *IntSet = NewIntSet()
	is.Set(100, 1)
	assertValue(t, is.Get(107), 0)
}

func TestUpdateSinglePoint(t *testing.T) {
	var is *IntSet = NewIntSet()
	is.Set(100, 1)
	is.Set(100, 0)
	assertValue(t, is.Get(100), 0)
}

func assertValue(t *testing.T, actual, expected uint32) {
	if actual != expected {
		t.Errorf("expected %v but was %v", expected, actual)
	}
}
