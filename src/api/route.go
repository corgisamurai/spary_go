package api

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/ci",        Ci         ).Methods("GET")
	r.HandleFunc("/v1/spas",      ShowSpaList).Methods("GET")
	r.HandleFunc("/v1/spas/{id}", ShowSpa    ).Methods("GET")
	r.HandleFunc("/v1/spas",      CreateSpa  ).Methods("POST")
	http.Handle("/", r)

	fmt.Printf("Server is running... localhost:8080")
	http.ListenAndServe(":8080", nil)
}
