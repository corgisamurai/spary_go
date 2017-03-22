package test

import (
	"api"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	_ "net/url"
	_ "strings"
	"testing"
)

func TestShowSpa(t *testing.T) {
	db.Query("INSERT INTO spa (id, name, address) VALUES(?, ?, ?)", 1, "木下温泉", "北海道")
	db.Query("INSERT INTO spa (id, name, address) VALUES(?, ?, ?)", 2, "木下温泉2", "北海道2")

	r := mux.NewRouter()
	r.HandleFunc("/v1/spa/{id}", api.ShowSpa)

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	url := testServer.URL + "/v1/spa/1"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	spa := new(api.Spa)
	json.Unmarshal(([]byte)(string(data)), spa)

	if spa.Name != "木下温泉" {
		t.Fatalf("not 木下温泉, %s", spa.Name)
	}
	if spa.Address != "北海道" {
		t.Fatalf("not 北海道, %s", spa.Address)
	}
}

func TestShowAnotherSpa(t *testing.T) {
	db.Query("INSERT INTO spa (id, name, address) VALUES(?, ?, ?)", 1, "木下温泉", "北海道")
	db.Query("INSERT INTO spa (id, name, address) VALUES(?, ?, ?)", 2, "木下温泉2", "北海道2")

	r := mux.NewRouter()
	r.HandleFunc("/v1/spa/{id}", api.ShowSpa)

	testServer := httptest.NewServer(r)
	defer testServer.Close()

	url := testServer.URL + "/v1/spa/2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	spa := new(api.Spa)
	json.Unmarshal(([]byte)(string(data)), spa)

	if spa.Name != "木下温泉2" {
		t.Fatalf("not 木下温泉2, %s", spa.Name)
	}
	if spa.Address != "北海道2" {
		t.Fatalf("not 北海道2, %s", spa.Address)
	}
}

func execShowSpaList() *api.Spas {
	//ShowSpaList実行
	req, _ := http.NewRequest("GET", "v1/spas", nil)
	res := httptest.NewRecorder()
	api.ShowSpaList(res, req)

	//レスポンスを構造体に変換
	data, _ := ioutil.ReadAll(res.Body)
	spas := new(api.Spas)
	json.Unmarshal(([]byte)(string(data)), spas)
	return spas
}

func TestShowSpaList(t *testing.T) {
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉", "北海道")
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉2", "北海道2")

	result := execShowSpaList()

	assertEqual(t, result.Spas[0].Name, "木下温泉")
	assertEqual(t, result.Spas[0].Address, "北海道")
	assertEqual(t, result.Spas[1].Name, "木下温泉2")
	assertEqual(t, result.Spas[1].Address, "北海道2")
}
