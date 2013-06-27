package av

import (
	"encoding/gob"
	"github.com/steveyen/gkvlite"
	"os"
)

type DB struct {
	store   *gkvlite.Store
	objects *gkvlite.Collection
}

func NewDB(fname string) *DB {
	f, err := os.OpenFile(fname, os.O_SYNC, 0666)
	if err != nil {
		f, err = os.Create(fname)
		if err != nil {
			panic(err)
		}
	}
	store, _ := gkvlite.NewStore(f)
	objects := store.SetCollection("objects", nil)
	return &DB{
		store:   store,
		objects: objects,
	}
}

func (db *DB) CloseDB() {
	db.store.Close()
}

func (db *DB) getObject(id string) *object {
	typeID, err := db.objects.Get([]byte(id))
	if err != nil {
		return nil
	}
	return &object{ID: id, TypeID: string(typeID)}
}

func (db *DB) writeObject(obj *object) {
	gob.Register(obj)

	typeID, err := db.objects.Get([]byte(id))
	if err != nil {
		return nil
	}
	return &object{ID: id, TypeID: string(typeID)}
}

// fname := "myfile.db"
// 	f, _ := os.Create(fname)
// 	defer os.Remove(fname)
// 	s, _ := gkvlite.NewStore(f)
// 	c := s.SetCollection("cars", nil)

// 	c.Set([]byte("tesla"), []byte("$$$"))
// 	c.Set([]byte("mercedes"), []byte("$$"))
// 	c.Set([]byte("bmw"), []byte("$"))

// 	s.Flush()
// 	f.Sync()
