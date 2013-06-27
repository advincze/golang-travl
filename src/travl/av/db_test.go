package av

import (
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
	db.CloseDB()

}
