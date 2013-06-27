package av

import (
	"encoding/json"
	"time"
)

type BitVector struct {
	Resolution TimeResolution `json:"resolution"`
	Start      time.Time      `json:"start"`
	Data       []byte         `json:"data"`
}

func (bv *BitVector) MarshalJSON() ([]byte, error) {

	intdata := make([]int, len(bv.Data))
	for k, v := range bv.Data {
		intdata[k] = int(v)
	}

	return json.Marshal(struct {
		Resolution string    `json:"resolution"`
		Start      time.Time `json:"start"`
		Data       []int     `json:"data"`
	}{
		Resolution: bv.Resolution.String(),
		Start:      bv.Start,
		Data:       intdata,
	})
}

func (bitVector *BitVector) All() bool {
	for _, b := range bitVector.Data {
		if b == 0 {
			return false
		}
	}
	return true
}

func (bitVector *BitVector) Any() bool {
	for _, b := range bitVector.Data {
		if b == 1 {
			return true
		}
	}
	return false
}

func (bitVector *BitVector) Count() int {
	count := 0
	for _, b := range bitVector.Data {
		if b == 1 {
			count++
		}
	}
	return count
}
