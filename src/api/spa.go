package api

import (
	"encoding/json"
	"fmt"
	"lib"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Spas struct {
	Spas []Spa `json:"spa"`
}

type Spa struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Url       string `json:"url"`
	Tel       string `json:"tel"`
	Effect    string `json:"effect"`
	Fee       string `json:"fee"`
	Image     string `json:"image"`
	Equipment string `json:"equipment"`
	Workday   string `json:"workday"`
}

func ShowSpaList(w http.ResponseWriter, r *http.Request) {
	db := lib.DbOpen()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM spa")
	spas := Spas{}

	for rows.Next() {
		spa := Spa{}
		rows.Scan(&spa.Id, &spa.Name, &spa.Address, &spa.Url, &spa.Tel, &spa.Effect, &spa.Fee, &spa.Image, &spa.Equipment, &spa.Workday)
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
		row.Scan(&spa.Id, &spa.Name, &spa.Address, &spa.Url, &spa.Tel, &spa.Effect, &spa.Fee, &spa.Image, &spa.Equipment, &spa.Workday)
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
