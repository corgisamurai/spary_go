package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"lib"
	"net/http"
)

type Spas struct {
	Spas []Spa `json:"spa"`
}

type Spa struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Url     string `json:"url"`
}

func ShowSpaList(w http.ResponseWriter, r *http.Request) {
	db := lib.DbOpen()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM spa")
	spas := Spas{}
	for rows.Next() {
		spa := Spa{}
		rows.Scan(&spa.Id, &spa.Name, &spa.Address, &spa.Url)
		spas.Spas = append(spas.Spas, spa)
	}
	result, _ := json.Marshal(spas)
	fmt.Fprintf(w, string(result))
}

func ShowSpa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := lib.DbOpen()
	defer db.Close()

	query := "select * from spa where id = ?"
	row, _ := db.Query(query, vars["id"])
	spa := Spa{}
	for row.Next() {
		row.Scan(&spa.Id, &spa.Name, &spa.Address, &spa.Url)
	}

	result, _ := json.Marshal(spa)
	fmt.Fprintf(w, string(result))
}

func CreateSpa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create spa")
	db := lib.DbOpen()
	defer db.Close()

	query := "INSERT INTO spa (name,address) VALUES(?, ?)"
	db.Query(query, "AAA", "Where")
}
