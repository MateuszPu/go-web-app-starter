package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		lastName := r.URL.Query().Get("lastName")
		age, _ := strconv.Atoi(r.URL.Query().Get("age"))
		json.NewEncoder(w).Encode(User{name, lastName, age})
	})
}