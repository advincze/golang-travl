package av

import (
	"strconv"
)

type objectType struct {
	Name       string
	Resolution TimeResolution
	Objects    map[string]*object
}

var objectTypes = make(map[string]*objectType)

func GetObjectType(name string) *objectType {
	if ot, ok := objectTypes[name]; ok {
		return ot
	} else {
		ot = &objectType{Name: name, Objects: make(map[string]*object)}
		objectTypes[name] = ot
		return ot
	}
}

type object struct {
	ID     string
	TypeID string
	Ba     BitAv
}

func (ot *objectType) newObject(id string, res TimeResolution) *object {
	return &object{ID: id, TypeID: ot.Name, Ba: NewBitSegmentAv(id, res)}
}

func (ot *objectType) CreateObject() *object {
	i := len(ot.Objects) + 1
	id := strconv.Itoa(i)
	for _, ok := ot.Objects[id]; ok; id = strconv.Itoa(i) {
		i += 1
	}
	ob := ot.newObject(id, Minute5)
	ot.Objects[id] = ob
	return ob
}

func (ot *objectType) GetObject(id string) *object {
	ob, ok := ot.Objects[id]
	if !ok {
		ob = ot.newObject(id, Minute5)
		ot.Objects[id] = ob
	}
	return ob
}

func GetObjectTypeAndObject(name, id string) (*objectType, *object) {
	ot := GetObjectType(name)
	ob := ot.GetObject(id)
	// println("GetObjectTypeAndObject", ot, len(ot.Objects), ot.Objects["8"], ob)
	return ot, ob
}
