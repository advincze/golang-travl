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
	Id string
}

func (ot *objectType) NewObject() *object {
	i := len(ot.Objects) + 1
	id := strconv.Itoa(i)
	for _, ok := ot.Objects[id]; ok; id = strconv.Itoa(i) {
		i += 1
	}
	ob := &object{Id: id}
	ot.Objects[id] = ob
	return ob
}

func (ot *objectType) GetObject(id string) *object {
	ob, ok := ot.Objects[id]

	if !ok {
		ob = &object{Id: id}

		ot.Objects[id] = ob
	}
	return ob
}
