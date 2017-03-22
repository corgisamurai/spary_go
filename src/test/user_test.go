package test

import (
	"api"
	"net/http"
	"net/http/httptest"
	_ "net/url"
	_ "strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "v1/user/1", nil)
	res := httptest.NewRecorder()
	api.GetUser(res, req)

	if res.Code != 200 {
		t.Fatalf("not 200, %s", res.Code)
	}
}
