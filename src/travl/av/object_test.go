package av

import (
	"testing"
)

func TestShouldCreateAnObjectType(t *testing.T) {
	//given
	name := "myType"

	//when
	ot := GetObjectType(name)

	//then
	if ot == nil {
		t.Errorf("the created objectType must not be nil")
	}
}

func TestShouldRetrieveAnObjectType(t *testing.T) {
	//given
	name := "myType"
	ot := GetObjectType(name)

	//when
	ot2 := GetObjectType(name)

	//then
	if ot != ot2 {
		t.Errorf("the returned objectTypes should be equal: %v %v\n ", ot, ot2)
	}
}

func TestShouldCreateANewObject(t *testing.T) {
	//g
	ot := GetObjectType("myType")

	//w
	ob := ot.NewObject()

	//t
	if ob == nil {
		t.Errorf("the created object must not be nil")
	}
	if ob.Id == "" {
		t.Errorf("the created object must have an id")
	}
}

func TestShouldGetObject(t *testing.T) {
	//g
	ot := GetObjectType("myType")
	ob := ot.NewObject()

	//w
	ob2 := ot.GetObject(ob.Id)

	//t
	if ob2 != ob {
		t.Errorf("the retrieved object must be the same as the created")
	}
	if ob2.Id != ob.Id {
		t.Errorf("the retrieved object must have the same Id as the created one")
	}
}
