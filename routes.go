package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Contact - user message
type Contact struct {
	Message string `db:"message"`
	Email   string `db:"email"`
	Name    string `db:"name"`
}

// ContactsCreate - makes a contact
func ContactsCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("here i am")
	fmt.Print(r.Form["message"])
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nil)
}
