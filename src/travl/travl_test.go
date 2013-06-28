package main

import (
	"bytes"
	"encoding/json"
	. "launchpad.net/gocheck"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type TravlSuite struct {
	testServer *httptest.Server
}

var _ = Suite(&TravlSuite{})

func (s *TravlSuite) SetUpTest(c *C) {
	s.testServer = httptest.NewServer(createRouter())
}

func (s *TravlSuite) TearDownTest(c *C) {
	s.testServer.Close()
}

func (s *TravlSuite) TestDefineAv(c *C) {
	msg, _ := json.Marshal(struct {
		From  time.Time `json:"from"`
		To    time.Time `json:"to"`
		Value byte      `json:"value"`
	}{time.Now(), time.Now().Add(25 * time.Hour), 1})

	req, _ := http.NewRequest("PUT", s.testServer.URL+"/5/_av", bytes.NewBuffer(msg))
	res, _ := http.DefaultClient.Do(req)

	c.Check(res.StatusCode, Equals, http.StatusOK)

}

func (s *TravlSuite) TestRetrieveAv(c *C) {

	res, _ := http.Get(s.testServer.URL + "/5/_av?from=2013-06-28&to=2013-06-30&resolution=day")

	c.Check(res.StatusCode, Equals, http.StatusOK)

}
