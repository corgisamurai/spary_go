package test

import (
	"api"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type Spas struct {
	Spas []Spa `json:"spa"`
}
type Spa struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func TestShowSpa(t *testing.T) {
	newSpaRecord(1, "木下温泉", "北海道", "http://kika.pc/onsen")
	newSpaRecord(2, "木下温泉2", "北海道2", "http://kika.pc/onsen2")

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

	// assertEqual(t, spa.Name, "木下温泉")
	// assertEqual(t, spa.Address, "北海道")
	// assertEqual(t, spa.Url, "http://kika.pc/onsen")
}
func TestShowSpaListRequest(t *testing.T) {
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉", "北海道")
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉2", "北海道2")

	r := mux.NewRouter()
	r.HandleFunc("/v1/spas", api.ShowSpa)

	ts := httptest.NewServer(r)
	defer ts.Close()

	url := ts.URL + "/v1/spas"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) == "404 page not found" {
		t.Fatal("404 error")
	}
	//spas := new(Spas)
	//result := json.Unmarshal(([]byte)(string(body)), spas)
	//assertEqual(t, result.Spas[0].Name, "木下温泉")
	//assertEqual(t, result.Spas[0].Address, "北海道")
	//assertEqual(t, result.Spas[1].Name, "木下温泉2")
	//assertEqual(t, result.Spas[1].Address, "北海道2")
}

func TestShowAnotherSpa(t *testing.T) {
	newSpaRecord(1, "木下温泉", "北海道", "http://kika.pc/onsen")
	newSpaRecord(2, "木下温泉2", "北海道2", "http://kika.pc/onsen2")

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

	// assertEqual(t, spa.Name, "木下温泉2")
	// assertEqual(t, spa.Address, "北海道2")
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

func newSpaRecord(id int, name string, address string, url string) {
	db.Query("INSERT INTO spa (id, name, address, url) VALUES(?, ?, ?, ?)", id, name, address, url)
}

func TestShowSpaList(t *testing.T) {
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉", "北海道")
	db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉2", "北海道2")

	// result := execShowSpaList()

	// assertEqual(t, result.Spas[0].Name, "木下温泉")
	// assertEqual(t, result.Spas[0].Address, "北海道")
	// assertEqual(t, result.Spas[1].Name, "木下温泉2")
	// assertEqual(t, result.Spas[1].Address, "北海道2")
}
