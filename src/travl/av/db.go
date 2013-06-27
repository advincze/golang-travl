package av

import (
	"bytes"
	"encoding/gob"
	"github.com/steveyen/gkvlite"
	"os"
)

type DB struct {
	file    *os.File
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
		file:    f,
		store:   store,
		objects: objects,
	}
}

func (db *DB) CloseDB() {
	db.store.Close()
}

func (db *DB) getObject(id string) *object {
	gob.Register((*object)(nil))
	gobData, err := db.objects.Get([]byte(id))
	buf := bytes.NewBuffer(gobData)
	dec := gob.NewDecoder(buf)
	if err != nil {
		panic(err)
	}
	var obj *object
	err = dec.Decode(&obj)
	if err != nil {
		panic(err)
	}
	return obj
}

func (db *DB) writeObject(obj *object) {
	gob.Register(obj)
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)
	enc.Encode(obj)

	db.objects.Set([]byte(obj.ID), m.Bytes())
	db.store.Flush()
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
