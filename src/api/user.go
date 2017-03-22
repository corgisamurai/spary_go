package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"lib"
	"net/http"
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
