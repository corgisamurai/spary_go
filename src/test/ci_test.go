package test

import (
  "testing"
  "net/http/httptest"
  "net/http"
  "io/ioutil"
  "api"
  // "strings"
)

func execCi() string {
	req, _ := http.NewRequest("GET", "v1/ci", nil)
	res := httptest.NewRecorder()
	api.Ci(res, req)
  data, _ := ioutil.ReadAll(res.Body)
  return string(data)
}

func TestCi(t *testing.T) {
  // TODO: slackのresponseをどう制御するか
  // result := execCi();
  // if strings.Contains(string(result), "Success") == false {
  //   t.Fatalf("%v != %v", string(result), "Success")
  // }
}
