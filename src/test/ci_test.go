package test

import (
  "testing"
  "net/http/httptest"
  "net/http"
  "io/ioutil"
  "api"
)

func execCi() string {
	req, _ := http.NewRequest("GET", "v1/ci", nil)
	res := httptest.NewRecorder()
	api.Ci(res, req)
  data, _ := ioutil.ReadAll(res.Body)
  return string(data)
}

func TestCi(t *testing.T) {
  result := execCi();
  assertEqual(t, "0", result)
}
