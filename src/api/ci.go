package api

import (
	"fmt"
	"net/http"
)

func Ci(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string("ok"))
}
