package bitav

import (
	"bytes"
	. "launchpad.net/gocheck"
	"testing"
	"time"
	"travl/av"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type BitAvSuite struct {
	ba BitAv
	t1 time.Time
}

var _ = Suite(&BitAvSuite{})

func (s *BitAvSuite) SetUpTest(c *C) {
	s.ba = NewSegmentBitAv("testID", av.Minute5, "mem")
	s.t1 = time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
}

func (s *BitAvSuite) TearDownTest(c *C) {
	s.ba = nil
}

func (s *BitAvSuite) TestNewBitAvShouldNotBeNil(c *C) {
	c.Assert(s.ba, NotNil)
}

func (s *BitAvSuite) TestGetAvAtEmpty(c *C) {
	c.Assert(s.ba.GetAt(s.t1), Equals, byte(0))
}

func (s *BitAvSuite) TestGetAvAtSet(c *C) {
	s.ba.SetAt(s.t1, 1)
	c.Assert(s.ba.GetAt(s.t1), Equals, byte(1))
}

func (s *BitAvSuite) TestGetAvAtUnset(c *C) {
	s.ba.SetAt(s.t1, 0)
	c.Assert(s.ba.GetAt(s.t1), Equals, byte(0))
}

func (s *BitAvSuite) TestGetAvNothingFromEmpty(c *C) {
	bitVector := s.ba.Get(s.t1, s.t1, av.Minute5)
	c.Check(bitVector, NotNil)
	c.Check(len(bitVector.Data), Equals, 0)
}

func (s *BitAvSuite) TestGetAvFromEmpty(c *C) {
	bitVector := s.ba.Get(s.t1, s.t1.Add(25*time.Minute), av.Minute5)
	c.Check(len(bitVector.Data), Equals, 5)
	c.Assert(bitVector.Any(), Equals, false)
}

func (s *BitAvSuite) TestGetAvFromBeforeSet(c *C) {
	s.ba.Set(s.t1.Add(45*time.Minute), s.t1.Add(75*time.Minute), 1)

	bitVector := s.ba.Get(s.t1, s.t1.Add(25*time.Minute), av.Minute5)

	c.Check(len(bitVector.Data), Equals, 5)
	c.Assert(bitVector.Any(), Equals, false)
}

func (s *BitAvSuite) TestGetAvFromAfterSet(c *C) {
	s.ba.Set(s.t1, s.t1.Add(25*time.Minute), 1)

	bitVector := s.ba.Get(s.t1.Add(45*time.Minute), s.t1.Add(75*time.Minute), av.Minute5)

	c.Check(len(bitVector.Data), Equals, 6)
	c.Assert(bitVector.Any(), Equals, false)
}

func (s *BitAvSuite) TestGetAvFromInsideSet(c *C) {
	s.ba.Set(s.t1, s.t1.Add(55*time.Minute), 1)

	bitVector := s.ba.Get(s.t1.Add(25*time.Minute), s.t1.Add(45*time.Minute), av.Minute5)

	c.Check(len(bitVector.Data), Equals, 4)
	c.Assert(bitVector.All(), Equals, true)
}

func (s *BitAvSuite) TestGetAvFromItersectBeforeSet(c *C) {
	s.ba.Set(s.t1.Add(15*time.Minute), s.t1.Add(55*time.Minute), 1)

	bitVector := s.ba.Get(s.t1, s.t1.Add(45*time.Minute), av.Minute5)

	c.Assert(len(bitVector.Data), Equals, 9)
	c.Assert(bitVector.Count(), Equals, 6)
}

func (s *BitAvSuite) TestGetAvFromItersectAfterSet(c *C) {
	s.ba.Set(s.t1, s.t1.Add(45*time.Minute), 1)

	bitVector := s.ba.Get(s.t1.Add(15*time.Minute), s.t1.Add(55*time.Minute), av.Minute5)

	c.Assert(len(bitVector.Data), Equals, 8)
	c.Assert(bitVector.Count(), Equals, 6)
}

func (s *BitAvSuite) TestGetAvWithLowerResolution(c *C) {
	s.ba.Set(s.t1.Add(15*time.Minute), s.t1.Add(55*time.Minute), 1)

	bitVector := s.ba.Get(s.t1, s.t1.Add(45*time.Minute), av.Minute15)

	c.Assert(len(bitVector.Data), Equals, 3)
	c.Assert(bitVector.Count(), Equals, 2)
}

func (s *BitAvSuite) TestGetAvWithHigherResolution(c *C) {
	s.ba.Set(s.t1.Add(15*time.Minute), s.t1.Add(55*time.Minute), 1)

	bitVector := s.ba.Get(s.t1, s.t1.Add(45*time.Minute), av.Minute)

	c.Assert(len(bitVector.Data), Equals, 45)
	c.Assert(bitVector.Count(), Equals, 30)
}

func (s *BitAvSuite) TestSetAvTwoYearsWorkingHours(c *C) {
	t1 := time.Date(1982, 2, 7, 9, 0, 0, 0, time.UTC)
	t2 := time.Date(1983, 4, 5, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 2*365; i++ {
		s.ba.Set(t1, t1.Add(8*time.Hour), 1)
		s.ba.Set(t1.Add(8*time.Hour), t1.Add(12*time.Hour), 0)
		t1 = t1.Add(24 * time.Hour)
	}

	bitVector := s.ba.Get(t2, t2.Add(24*time.Hour), av.Hour)
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}
	c.Assert(bytes.Equal(expected, bitVector.Data), Equals, true)

}

func (s *BitAvSuite) TestSetAvTwoYearsWorkingHoursBackwards(c *C) {
	t1 := time.Date(1982, 2, 7, 9, 0, 0, 0, time.UTC)
	t2 := time.Date(1981, 4, 5, 0, 0, 0, 0, time.UTC)

	for i := 0; i < 2*365; i++ {
		s.ba.Set(t1, t1.Add(8*time.Hour), 1)
		s.ba.Set(t1.Add(8*time.Hour), t1.Add(12*time.Hour), 0)
		t1 = t1.Add(-24 * time.Hour)
	}

	bitVector := s.ba.Get(t2, t2.Add(24*time.Hour), av.Hour)
	expected := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0}
	c.Assert(bytes.Equal(expected, bitVector.Data), Equals, true)

}

func (s *BitAvSuite) BenchmarkSetAvOneDay2(c *C) {
	t1 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)

	for i := 0; i < c.N; i++ {
		s.ba.Set(t1, t1.Add(24*time.Hour), 1)
		t1 = t1.Add(48 * time.Hour)
	}
}

func (s *BitAvSuite) BenchmarkGetAvOneDay2(c *C) {
	t0 := time.Date(1982, 2, 7, 0, 0, 0, 0, time.UTC)
	t1 := t0

	for i := 0; i < 2*365; i++ {
		s.ba.Set(t1, t1.Add(8*time.Hour), 1)
		s.ba.Set(t1.Add(8*time.Hour), t1.Add(12*time.Hour), 0)
		t1 = t1.Add(24 * time.Hour)
	}

	t1 = t0.Add(4 * time.Hour)
	for i := 0; i < c.N; i++ {
		s.ba.Get(t1, t1.Add(72*time.Hour), av.Hour)

		t1 = t1.Add(23 * time.Hour)
		if i%365 == 0 {
			t1 = t0
		}
	}
}
