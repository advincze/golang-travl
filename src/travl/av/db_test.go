package av

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"
)

const testDBFileName string = "testfile.db"

func TestCreateDB(t *testing.T) {
	db := NewDB(testDBFileName)
	db.CloseDB()

	if _, err := os.Stat(testDBFileName); os.IsNotExist(err) {
		t.Errorf("no such file or directory: %s", testDBFileName)
	} else {
		os.Remove(testDBFileName)
	}
}

func TestOpenExistingDB(t *testing.T) {
	os.Create(testDBFileName)
	defer os.Remove(testDBFileName)

	db := NewDB(testDBFileName)
	db.CloseDB()
}

func TestWriteObjectIntoDB(t *testing.T) {
	defer os.Remove(testDBFileName)
	db := NewDB(testDBFileName)

	obj := &object{ID: "myID", TypeID: "myTypeID"}
	db.writeObject(obj)

	db.CloseDB()

	db = NewDB(testDBFileName)

	obj2 := db.getObject("myID")

	db.CloseDB()

	if obj.ID != obj2.ID {
		t.Errorf("the written and retrieved object should be equal, were %s, %s", obj.ID, obj2.ID)
	}

	if obj.TypeID != obj2.TypeID {
		t.Errorf("the written and retrieved object should be equal, were %s, %s", obj.TypeID, obj2.TypeID)
	}
}

func TestWriteObjectIntoInMemoryDB(t *testing.T) {

	db := NewInMemoryDB()

	obj := &object{ID: "myID", TypeID: "myTypeID"}
	db.writeObject(obj)

	obj2 := db.getObject("myID")

	db.CloseDB()

	if obj.ID != obj2.ID {
		t.Errorf("the written and retrieved object should be equal, were %s, %s", obj.ID, obj2.ID)
	}

	if obj.TypeID != obj2.TypeID {
		t.Errorf("the written and retrieved object should be equal, were %s, %s", obj.TypeID, obj2.TypeID)
	}
}

func TestGob(t *testing.T) {
	type Message struct {
		ID   string
		Text string
	}
	m := &Message{ID: "myID", Text: "myText"}
	gob.Register((*Message)(nil))
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	bb := buf.Bytes()

	bb2 := make([]byte, len(bb))
	copy(bb2, bb)

	buf2 := bytes.NewBuffer(bb2)
	dec := gob.NewDecoder(buf2)
	if err != nil {
		panic(err)
	}
	var m2 *Message
	err = dec.Decode(&m2)
	if err != nil {
		panic(err)
	}
	// println(m2.ID, m2.Text)
}
