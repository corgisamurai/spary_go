package test

import (
  "api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func execShowSpaList() *api.Result {
  //ShowSpaList実行
	req, _ := http.NewRequest("GET", "v1/spas", nil)
	res := httptest.NewRecorder()
	api.ShowSpaList(res, req)

  //レスポンスを構造体に変換
	data, _ := ioutil.ReadAll(res.Body)
	obj := new(api.Result)
	json.Unmarshal(([]byte)(string(data)), obj)
  return obj
}

func TestShowSpaList(t *testing.T) {
  db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉", "北海道")
  db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉2", "北海道2")

  result := execShowSpaList()

  assetEqual(t, result.Spas[0].Name,    "木下温泉")
  assetEqual(t, result.Spas[0].Address, "北海道")
  assetEqual(t, result.Spas[1].Name,    "木下温泉2")
  assetEqual(t, result.Spas[1].Address, "北海道2")
}