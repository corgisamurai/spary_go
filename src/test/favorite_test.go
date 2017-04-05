package test

import (
	"api"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Spa struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	SpaId  string `json:"spa_id"`
}

func TestAddFavorite(t *testing.T) {
	// r := mux.NewRouter()
	// r.HandleFunc("/v1/favorite/", api.AddFavorite)

	// testServer := httptest.NewServer(r)
	// defer testServer.Close()

	// url := testServer.URL + "/v1/favorite"
	// resp, err := http.Get(url)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}
