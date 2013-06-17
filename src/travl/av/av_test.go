package av

import (
	"testing"
	"time"
)

func TestTimeToUnit(t *testing.T) {
	t0 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(1982, 2, 7, 0, 13, 0, 0, time.UTC)
	t2 := time.Date(1982, 2, 7, 0, 58, 0, 0, time.UTC)

	u0 := timeToUnit(t0, Hour)
	u1 := timeToUnit(t1, Hour)
	u2 := timeToUnit(t2, Hour)

	if !(u0 == u1 && u1 == u2) {
		t.Errorf("units should be equal, %v, %v, %v ", u0, u1, u2)
	}
}

func TestFloorDate(t *testing.T) {
	t0 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t1 := time.Date(1982, 2, 7, 0, 13, 0, 0, time.UTC)

	t2 := floorDate(t1, Hour)

	if t2 != t0 {
		t.Errorf("dates should be equal, %v, %v ", t0, t2)
	}
}
