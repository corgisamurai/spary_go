package api

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"lib"
	"net/http"
)

type Comment struct {
	Id      string `json:"id"`
	SpaId   string `json:"spa_id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &comment); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db := lib.DbOpen()
	query := "insert into comments (spa_id, user_id, comment) values (?, ?, ?)"
	db.Query(query, comment.SpaId, comment.UserId, comment.Comment)
}
