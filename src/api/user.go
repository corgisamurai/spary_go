package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := lib.DbOpen()
	defer db.Close()

	query := "select * from users where id = ?"
	row, _ := db.Query(query, vars["id"])
	user := User{}
	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email, &user.Address)
	}

	result, _ := json.Marshal(user)
	fmt.Fprintf(w, string(result))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db := lib.DbOpen()
	query := "insert into users (name, email, address) values (?, ?, ?)"
	db.Query(query, user.Name, user.Email, user.Address)
}

func AuthUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := lib.DbOpen()
	defer db.Close()

	query := "select * from users where name = ? and password = ?"
	row, _ := db.Query(query, vars["name"], vars["pass"])

	count := 0
	for row.Next() {
		row.Scan(&count)
	}
	if count == 0 {
		fmt.Fprintf(w, "fail")
	} else {
		fmt.Fprintf(w, "success")
	}
}
