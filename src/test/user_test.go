package test

import (
	"api"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func TestGetUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "v1/user/1", nil)
	res := httptest.NewRecorder()
	api.GetUser(res, req)

	if res.Code != 200 {
		t.Fatalf("not 200, %s", res.Code)
	}
}

func TestGetUserJson(t *testing.T) {
	db.Query(
		"INSERT INTO users (id, name, email, address) VALUES(?, ?, ?, ?)",
		1, "test name", "test@sample.com", "test address")

	router := mux.NewRouter()
	router.HandleFunc("/v1/user/{id}", api.GetUser)

	testServer := httptest.NewServer(router)
	defer testServer.Close()

	url := testServer.URL + "/v1/user/1"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	user := new(api.User)
	json.Unmarshal(([]byte)(string(data)), user)

	if user.Name != "test name" {
		t.Fatalf("not test name, %s", user.Name)
	}
	if user.Email != "test@sample.com" {
		t.Fatalf("not test@sample.com, %s", user.Address)
	}
	if user.Address != "test address" {
		t.Fatalf("not test address, %s", user.Address)
	}
}

func TestAuthFail(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/kanai/pass", api.AuthUser)
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	url := testServer.URL + "/v1/auth/kanai/pass"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if string(data) != "fail" {
		t.Fatalf("not fail", string(data))
	}
}

func TestAuthSuccess(t *testing.T) {
	db.Query(
		"INSERT INTO users (name, pasword) VALUES(?, ?)",
		"kanai", "pass",
	)

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth/kanai/pass", api.AuthUser)
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	url := testServer.URL + "/v1/auth/kanai/pass"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if string(data) != "success" {
		t.Fatalf("not success", string(data))
	}
}
