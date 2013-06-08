package av

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
	id := string(len(ot.Objects) + 1)
	ob := &object{Id: id}
	ot.Objects[id] = ob
	return ob
}

func (ot *objectType) GetObject(id string) *object {
	return ot.Objects[id]
}
