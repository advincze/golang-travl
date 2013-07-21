package bitav

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type BitSegmentMgoPersistor struct {
	session *mgo.Session
}

func NewBitSegmentPersistor(kind string) *BitSegmentMgoPersistor {
	return new(BitSegmentMgoPersistor)
}

func (bsp *BitSegmentMgoPersistor) Save(s *BitSegment) {
	var err error
	if bsp.session == nil {
		bsp.session, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		bsp.session.SetMode(mgo.Monotonic, true)
	}

	c := bsp.session.DB("test").C("bitsegments")
	_, err = c.Upsert(
		bson.M{"id": s.ID, "start": s.start},
		bson.M{"id": s.ID, "start": s.start, "data": s.Bytes()},
	)
	if err != nil {
		panic(err)
	}
}

func (bsp *BitSegmentMgoPersistor) FindBitSegment(id string, start int) *BitSegment {
	return nil
}

func (bsp *BitSegmentMgoPersistor) FindAllSegments(id string) map[int]*BitSegment {
	return nil
}
