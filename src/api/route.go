package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/ci", Ci).Methods("GET")
	r.HandleFunc("/v1/spas", ShowSpaList)
	r.HandleFunc("/v1/spa/{id}", ShowSpa)
	r.HandleFunc("/v1/spa", CreateSpa).Methods("POST")
	r.HandleFunc("/v1/user/{id}", GetUser)
	r.HandleFunc("/v1/user", CreateUser).Methods("POST")
	r.HandleFunc("/v1/comment", CreateUser).Methods("POST")
	http.Handle("/", r)

	fmt.Printf("Server is running... localhost:8080")
	http.ListenAndServe(":8080", nil)
}
