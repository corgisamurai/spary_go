package api

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show spa")
}
